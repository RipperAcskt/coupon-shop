package transaction

import (
	"shop-smart-api/internal/entity"
)

type PaymentInterface interface {
	CreatePayment(payment entity.Payment) (*entity.Payment, error)
	GetPayments(userId string) ([]entity.Payment, error)
	UpdatePayment(id string) (*entity.Payment, error)
}

type PaymentService struct {
	repository PaymentInterface
}

func CreatePaymentService(r PaymentInterface) PaymentService {
	return PaymentService{r}
}

func (s PaymentService) CreatePayment(payment entity.Payment) (*entity.Payment, error) {
	return s.repository.CreatePayment(payment)
}

func (s PaymentService) GetPayments(userId string) ([]entity.Payment, error) {
	return s.repository.GetPayments(userId)
}

func (s PaymentService) UpdatePayment(id string) (*entity.Payment, error) {
	return s.repository.UpdatePayment(id)
}
