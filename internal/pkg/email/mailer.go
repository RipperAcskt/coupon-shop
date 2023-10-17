package email

import (
	"fmt"
	"github.com/alexeyco/unisender"
	"log"
)

type Mailer interface {
	Send(email, code string)
	SendCoupon(email, code string)
}

type mailer struct {
	d *unisender.UniSender
}

func CreateMailer(d *unisender.UniSender) Mailer {
	return &mailer{d}
}

func (m *mailer) Send(email, code string) {
	_, err := m.d.SendEmail(email).
		SenderName("coupon-shop").
		SenderEmail("hjadsfbnajv@gmail.com").
		Subject("ОТП код").
		Body(fmt.Sprintf("ОТП код: %v", code)).
		LangDE().
		ListID(1).
		WrapTypeSkip().Execute()

	if err != nil {
		log.Printf("send email otp failed: %v", err)
	}
}

func (m *mailer) SendCoupon(email, code string) {
	_, err := m.d.SendEmail(email).
		SenderName("coupon-shop").
		SenderEmail("hjadsfbnajv@gmail.com").
		Subject("ОТП код").
		Body(fmt.Sprintf("Ваш купон: %s\nНазовите его администратору магазина для активации!", code)).
		LangRU().
		ListID(1).
		WrapTypeSkip().Execute()
	if err != nil {
		log.Printf("send email otp failed: %v", err)
	}
}
