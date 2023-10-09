package payment

import (
	"fmt"
	"github.com/albenik-go/yookassa"
	"github.com/google/uuid"
	"shop-smart-api/internal/entity"
	"shop-smart-api/pkg"
)

type Payment struct {
	client *yookassa.Client

	cfg *pkg.AppConfig
}

func New(cfg *pkg.AppConfig) Payment {
	client := yookassa.New(cfg.Yookassa.ID, cfg.Yookassa.ApiKey)

	return Payment{
		client: client,

		cfg: cfg,
	}
}

func (p Payment) CreatePayment(paymentRequest *entity.CreatePaymentRequest) (interface{}, error) {
	payment := &yookassa.PaymentRequest{
		Amount: yookassa.Amount{
			Value:    paymentRequest.Amount,
			Currency: "RUB",
		},
		Description: fmt.Sprintf("Create payment for %s - %s", paymentRequest.PaymentType, paymentRequest.TypeID),
		Capture:     true,
		Confirmation: entity.Confirmation{
			ConfirmationType: "redirect",
			RedirectUrl:      paymentRequest.RedirectUrl,
		},
	}

	resp, err := p.client.CreatePayment(uuid.NewString(), payment)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
