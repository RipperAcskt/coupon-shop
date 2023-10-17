package payment

import (
	"fmt"
	"github.com/albenik-go/yookassa"
	"github.com/google/uuid"
	"math/rand"
	"shop-smart-api/internal/entity"
	"shop-smart-api/pkg"
	"strconv"
	"time"
)

var (
	ErrNoSuchCoupon = fmt.Errorf("no such coupon")
)

type PaymentInterface interface {
	CreatePayment(payment entity.Payment) (*entity.Payment, error)
	GetPayments(userId string) ([]entity.Payment, error)
	UpdatePayment(id string) (*entity.Payment, error)
	UpdateSubscription(id, userId string) error
}

type CodesInterface interface {
	CreateCode(code entity.UsersCodes) (*entity.UsersCodes, error)
	GetCodeByCoupon(coupon string) (*entity.UsersCodes, error)
	GetCodeByUserID(userID string) ([]entity.UsersCodes, error)
	DeleteCode(coupon string) (*entity.UsersCodes, error)
}

type SmsInterface interface {
	SendCoupon(phone, code string)
}

type MailInterface interface {
	SendCoupon(email, code string)
}

type UserRepository interface {
	Get(id string) (*entity.User, error)
}

type Payment struct {
	client *yookassa.Client

	repository PaymentInterface
	userRepo   UserRepository
	smsClient  SmsInterface
	mailClient MailInterface
	codeRepo   CodesInterface

	cfg *pkg.AppConfig
}

func New(r PaymentInterface, cfg *pkg.AppConfig, sms SmsInterface, mail MailInterface, userRepo UserRepository, code CodesInterface) Payment {
	client := yookassa.New(cfg.Yookassa.ID, cfg.Yookassa.ApiKey)

	return Payment{
		client: client,

		repository: r,
		userRepo:   userRepo,
		smsClient:  sms,
		mailClient: mail,
		codeRepo:   code,

		cfg: cfg,
	}
}

func (p Payment) CreatePayment(paymentRequest *entity.CreatePaymentRequest, userId string) (interface{}, error) {
	id := uuid.NewString()
	redirectUrl := p.cfg.Server.ServiceHost + "/api/payment/confirm/" + id + "/" + userId

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
		Id:          id,
		Value:       am,
		Status:      false,
		UserID:      userId,
		PaymentType: paymentRequest.PaymentType,
		TypeID:      paymentRequest.TypeID,
	}
	paymentResp, err := p.repository.CreatePayment(paymentTransfer)
	if err != nil {
		return nil, fmt.Errorf("create payment failed: %w", err)
	}

	paymentResp.RedirectURL = redirectUrl
	paymentResp.ConfirmationURL = resp.Confirmation
	return paymentResp, nil
}

func (p Payment) GetPayments(userId string) ([]entity.Payment, error) {
	return p.repository.GetPayments(userId)
}

func (p Payment) ConfirmPayment(id, userId string) error {
	payment, err := p.repository.UpdatePayment(id)
	if err != nil {
		return fmt.Errorf("update payment failed: %w", err)
	}

	var code string
	for {
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		code = fmt.Sprint(r.Intn(999999-100000) + 100000)
		_, err := p.codeRepo.GetCodeByCoupon(code)
		if err != nil {
			break
		}
	}

	codeUser := entity.UsersCodes{
		Id:     uuid.NewString(),
		UserId: payment.UserID,
		Coupon: code,
	}

	_, err = p.codeRepo.CreateCode(codeUser)
	if err != nil {
		return fmt.Errorf("create coupon failed: %w", err)
	}

	switch payment.PaymentType {
	case entity.Coupon:
		user, err := p.userRepo.Get(userId)
		if err != nil {
			return fmt.Errorf("get user failed: %w", err)
		}

		if user.Phone != "" {
			p.smsClient.SendCoupon(user.Phone, code)
		}
		if user.Email != "" {
			p.mailClient.SendCoupon(user.Email, code)
		}

	case entity.Subscription:
		err := p.repository.UpdateSubscription(payment.TypeID, userId)
		if err != nil {
			return fmt.Errorf("update subscription failed: %w", err)
		}
	}
	return nil
}

func (p Payment) ActivatePayment(coupon string) error {
	_, err := p.codeRepo.GetCodeByCoupon(coupon)
	if err != nil {
		return ErrNoSuchCoupon
	}

	_, err = p.codeRepo.DeleteCode(coupon)
	if err != nil {
		return fmt.Errorf("delete coupon failed: %w", err)
	}

	return nil
}
