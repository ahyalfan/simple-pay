package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/util"
	"context"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type transactionService struct {
	accountRepository     domain.AccountRepository
	transactionRepository domain.TransactionRepository
	cacheRepository       domain.CacheRepository
	db                    *gorm.DB
}

func NewTransaction(accountRepository domain.AccountRepository, transactionRepository domain.TransactionRepository, cacheRepository domain.CacheRepository, db *gorm.DB) domain.TransactionService {
	return &transactionService{
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
		cacheRepository:       cacheRepository,
		db:                    db,
	}
}

// TransferExecute implements domain.TransactionService.
func (t *transactionService) TransferExecute(ctx context.Context, req dto.TransferExecuteReq) error {
	tx := t.db.Begin() // agar jika terjadi error maka akan di rollback semua
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

	tx.Commit()
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
