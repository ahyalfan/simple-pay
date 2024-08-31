package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/component"
	"ahyalfan/golang_e_money/internal/util"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type transactionService struct {
	accountRepository     domain.AccountRepository
	transactionRepository domain.TransactionRepository
	cacheRepository       domain.CacheRepository
	db                    *gorm.DB
	// cara lama
	// notificationRepository domain.NotificationRepository
	// hub                    *dto.Hub
	notificationService domain.NotificationService
}

func NewTransaction(accountRepository domain.AccountRepository,
	transactionRepository domain.TransactionRepository,
	cacheRepository domain.CacheRepository,
	db *gorm.DB,
	notificationService domain.NotificationService) domain.TransactionService {
	return &transactionService{
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
		cacheRepository:       cacheRepository,
		db:                    db,
		notificationService:   notificationService,
	}
}

// TransferExecute implements domain.TransactionService.
func (t *transactionService) TransferExecute(ctx context.Context, req dto.TransferExecuteReq) error {
	tx := t.db.Begin() // agar jika terjadi error maka akan di rollback semua
	component.Log.Info("starting execute transfer transaction")
	defer tx.Rollback()
	// atau sebenarnya kita bisa bikin coloumn baru credit ulang, yg mana ini bisa di refound ketika ada massalah pengirimanya
	// yang nanti akan ke record sehingga kita bisa memantaunya

	val, err := t.cacheRepository.Get(req.InquiryKey)
	if err != nil {
		return domain.ErrInquiryNotFound
	}

	var reqInquiry dto.TransferInQuiryReq
	_ = json.Unmarshal(val, &reqInquiry)
	if reqInquiry == (dto.TransferInQuiryReq{}) {
		return domain.ErrInquiryNotFound
	}
	user := ctx.Value("x-user").(dto.UserData)
	myAccount, err := t.accountRepository.FindByUserId(ctx, user.ID)
	if err != nil {
		return err
	}

	dofAccount, err := t.accountRepository.FindByAccountNumber(ctx, reqInquiry.AccountNumber)
	if err != nil {
		return err
	}

	component.Log.Debugf("%s to %s", myAccount.AccountNumber, dofAccount.AccountNumber)

	debitTransaction := domain.Transaction{
		AccountId:       myAccount.ID,
		SofNumber:       myAccount.AccountNumber,
		DofNumber:       dofAccount.AccountNumber,
		Amount:          reqInquiry.Ammount,
		TransactionType: "D", // di ini debit
	}
	err = t.transactionRepository.Insert(ctx, &debitTransaction)
	if err != nil {
		return err
	}
	creditTransaction := domain.Transaction{
		AccountId:       dofAccount.ID,
		DofNumber:       dofAccount.AccountNumber,
		SofNumber:       myAccount.AccountNumber,
		Amount:          reqInquiry.Ammount,
		TransactionType: "C", // di ini credit
	}
	err = t.transactionRepository.Insert(ctx, &creditTransaction)
	if err != nil {
		return err
	}
	// ini tidak diperlukan karena kita pakai db transacation sudah cukup
	// var mutext sync.Mutex
	// defer mutext.Unlock()
	// mutext.Lock()
	myAccount.Balance -= reqInquiry.Ammount
	err = t.accountRepository.Update(ctx, &myAccount)
	if err != nil {

		return err
	}

	dofAccount.Balance += reqInquiry.Ammount
	err = t.accountRepository.Update(ctx, &dofAccount)
	if err != nil {
		return err
	}
	// mengahapus cache di redis , karena sudah diexceute
	err = t.cacheRepository.Delete(req.InquiryKey)
	if err != nil {
		return err
	}

	tx.Commit()
	go t.notificationAfterTransfer(myAccount, dofAccount, reqInquiry.Ammount)
	return nil
}

// TransferInquiry implements domain.TransactionService.
func (t *transactionService) TransferInquiry(ctx context.Context, req dto.TransferInQuiryReq) (dto.TransferInQuiryRes, error) {
	// kita ambil data account dari middleware yg sudah disimpan ke locals menggunakan ctx.local
	// dengan begitu kita bisa ambil menggunkan ctx.value, yg mmana kita perlu conver datanya tapi ke dto yg sudah disesuaikan
	user := ctx.Value("x-user").(dto.UserData) // tapi ini kita wajib kasih middleware yg sudah dibuat, agar tidak erro
	myAccount, err := t.accountRepository.FindByUserId(ctx, user.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.TransferInQuiryRes{}, domain.ErrAccountNotFound
	}
	if err != nil {
		return dto.TransferInQuiryRes{}, err
	}

	// kita cek mau dikirim ke mana, apakah ada accountnya
	_, err = t.accountRepository.FindByAccountNumber(ctx, req.AccountNumber)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.TransferInQuiryRes{}, domain.ErrAccountNotFound
	}
	if err != nil {
		return dto.TransferInQuiryRes{}, err
	}

	// cek balancenya harus lebih besar dari yg dikiirm
	if myAccount.Balance < req.Ammount {
		return dto.TransferInQuiryRes{}, domain.ErrInsufficientBalance
	}

	inquiryKey := util.GeneratorRandomString(32)

	jsonData, _ := json.Marshal(req)

	_ = t.cacheRepository.Set(inquiryKey, jsonData)

	return dto.TransferInQuiryRes{InquiryKey: inquiryKey}, nil
}

// cara lama
// func (t *transactionService) notificationAfterTransfer(sofAccount domain.Account, dofAccount domain.Account, amount float64) {
// 	notificationSender := domain.Notification{
// 		UserID: sofAccount.UserID,
// 		Title:  "Transfer Berhasil",
// 		Body:   fmt.Sprintf("Transfer berhasil senilai %.2f berhasil", amount),
// 		IsRead: 0,
// 		Status: 1,
// 	}

// 	notificationReceiver := domain.Notification{
// 		UserID: dofAccount.UserID,
// 		Title:  "Dana Diterima",
// 		Body:   fmt.Sprintf("Dana diterima sebesar %.2f", amount),
// 		IsRead: 0,
// 		Status: 1,
// 	}

// 	_ = t.notificationRepository.Insert(context.Background(), &notificationSender)
// 	// artinya jika di hub notifikasi channel memiliki key user id yg dicantumkan, maka lakukan perintah ifnya
// 	if channel, ok := t.hub.NotificationChannel[sofAccount.UserID]; ok {
// 		channel <- dto.NotificationData{
// 			// jika ada kita kirimkan datanya ke channel dengan key yg sudah ditentukan
// 			ID:        notificationSender.ID,
// 			Title:     notificationSender.Title,
// 			Body:      notificationSender.Body,
// 			IsRead:    notificationSender.IsRead,
// 			Status:    notificationSender.Status,
// 			CreatedAt: notificationSender.CreatedAt,
// 		}
// 	}

// 	_ = t.notificationRepository.Insert(context.Background(), &notificationReceiver)
// 	if channel, ok := t.hub.NotificationChannel[dofAccount.UserID]; ok {
// 		channel <- dto.NotificationData{
// 			ID:        notificationReceiver.ID,
// 			Title:     notificationReceiver.Title,
// 			Body:      notificationReceiver.Body,
// 			IsRead:    notificationReceiver.IsRead,
// 			Status:    notificationReceiver.Status,
// 			CreatedAt: notificationReceiver.CreatedAt,
// 		}
// 	}
// }

func (t *transactionService) notificationAfterTransfer(sofAccount domain.Account, dofAccount domain.Account, amount float64) {
	data := map[string]string{
		"amount": fmt.Sprintf("%.2f", amount),
	}

	_ = t.notificationService.Insert(context.Background(), sofAccount.UserID, "TRANSFER", data)
	_ = t.notificationService.Insert(context.Background(), dofAccount.UserID, "TRANSFER_DEST", data)
}
