package entity

import (
	"github.com/albenik-go/yookassa"
	"time"
)

type CreatePaymentRequest struct {
	Amount      string      `json:"amount"`
	RedirectUrl string      `json:"redirect_url"`
	PaymentType paymentType `json:"payment_type"`
	TypeID      string      `json:"type_id"`
}

type paymentType string

const (
	Coupon       paymentType = "coupon"
	Subscription paymentType = "subscription"
)

type Confirmation struct {
	ConfirmationType string `json:"type"`
	RedirectUrl      string `json:"return_url"`
}

type Payment struct {
	Id              string
	Value           float64
	Status          bool
	UserID          string
	UpdatedAt       time.Time
	PaymentType     paymentType
	TypeID          string
	RedirectURL     string
	ConfirmationURL yookassa.ConfirmationInfo
}
