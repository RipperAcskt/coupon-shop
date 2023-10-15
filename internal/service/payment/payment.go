package payment

import (
	"fmt"
	"github.com/albenik-go/yookassa"
	"github.com/google/uuid"
	"shop-smart-api/internal/entity"
	"shop-smart-api/pkg"
	"strconv"
)

type PaymentInterface interface {
	CreatePayment(payment entity.Payment) (*entity.Payment, error)
	GetPayments(userId string) (*entity.Payment, error)
	UpdatePayment(id int64) (*entity.Payment, error)
}

type Payment struct {
	client *yookassa.Client

	repository PaymentInterface

	cfg *pkg.AppConfig
}

func New(r PaymentInterface, cfg *pkg.AppConfig) Payment {
	client := yookassa.New(cfg.Yookassa.ID, cfg.Yookassa.ApiKey)

	return Payment{
		client: client,

		repository: r,

		cfg: cfg,
	}
}

func (p Payment) CreatePayment(paymentRequest *entity.CreatePaymentRequest, userId string) (interface{}, error) {
	id := uuid.NewString()
	redirectUrl := "http://localhost:8080/api/payment/confrim/" + id

	payment := &yookassa.PaymentRequest{
		Amount: yookassa.Amount{
			Value:    paymentRequest.Amount,
			Currency: "RUB",
		},
		Description: fmt.Sprintf("Create payment for %s - %s", paymentRequest.PaymentType, paymentRequest.TypeID),
		Capture:     true,
		Confirmation: entity.Confirmation{
			ConfirmationType: "redirect",
			RedirectUrl:      redirectUrl,
		},
	}

	resp, err := p.client.CreatePayment(id, payment)
	if err != nil {
		return nil, err
	}

	am, err := strconv.ParseFloat(paymentRequest.Amount, 64)
	if err != nil {
		return nil, fmt.Errorf("parse float failed: %w", err)
	}

	paymentTransfer := entity.Payment{
		Id:     id,
		Value:  am,
		Status: false,
		UserID: userId,
	}
	paymentResp, err := p.repository.CreatePayment(paymentTransfer)
	if err != nil {
		return nil, fmt.Errorf("create payment failed: %w", err)
	}

	paymentResp.RedirectURL = redirectUrl
	paymentResp.ConfirmationURL = resp.Confirmation
	return paymentResp, nil
}

func (p Payment) GetPayments(userId string) (*entity.Payment, error) {
	return p.repository.GetPayments(userId)
}

func (p Payment) UpdatePayment(id int64) (*entity.Payment, error) {
	return p.repository.UpdatePayment(id)
}
