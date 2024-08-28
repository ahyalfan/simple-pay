package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/internal/config"
	"context"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type midtransService struct {
	// ini kita pakai bawaan dari midtrans, jika ingin customize bisa sebenarnya tapi
	midtransCnf config.Midtrans
	envi        midtrans.EnvironmentType
}

func NewMidtrans(cnf *config.Config) domain.MidtransService {

	envi := midtrans.Sandbox
	if cnf.Midtrans.IsProd {
		envi = midtrans.Production
	}
	return &midtransService{midtransCnf: cnf.Midtrans, envi: envi}
}

// GenerateSnapURL implements domain.MidtransService.
func (ms *midtransService) GenerateSnapURL(ctx context.Context, t *domain.Topup) error {
	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  t.ID,
			GrossAmt: int64(t.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	var client snap.Client
	client.New(ms.midtransCnf.Key, ms.envi)
	// 3. Request create Snap transaction to Midtrans
	snapResp, err := client.CreateTransaction(req)
	if err != nil {
		return err
	}
	t.SnapLink = snapResp.RedirectURL
	return nil
}

// VerifyPayment implements domain.MidtransService.
func (m *midtransService) VerifyPayment(ctx context.Context, orderId string) (bool, error) {
	var client coreapi.Client

	client.New(m.midtransCnf.Key, m.envi)

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(orderId)
	if e != nil {
		return false, e
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			// ini kayak response nya sukses akan kamu apain
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					return true, nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				return true, nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
	return false, nil
}
