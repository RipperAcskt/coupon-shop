package email

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"shop-smart-api/pkg"
)

type Mailer interface {
	Send(email, code string)
	SendCoupon(email, code string)
}

type mailer struct {
	cfg *pkg.AppConfig
}

func CreateMailer(cfg *pkg.AppConfig) Mailer {
	return &mailer{cfg}
}

func (m *mailer) Send(email, code string) {
	baseURL := "https://api.unisender.com"
	resource := "/ru/api/sendEmail"
	params := url.Values{}
	params.Add("format", "json")
	params.Add("api_key", m.cfg.ApiMailer)
	params.Add("email", email)
	params.Add("sender_name", "Coupon shop")
	params.Add("sender_email", "hjadsfbnajv@gmail.com")
	params.Add("subject", "ОТП код")
	params.Add("body", code)
	params.Add("list_id", "1")
	params.Add("lang", "ru")

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		log.Printf("parse request uri failed: %v", err)
	}
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resp, err := http.Get(urlStr)
	if err != nil {
		log.Printf("get failed: %v", err)
	}

	var b []byte
	resp.Body.Read(b)
	log.Println(string(b))
}

func (m *mailer) SendCoupon(email, code string) {
	baseURL := "https://api.unisender.com"
	resource := "/ru/api/sendEmail"
	params := url.Values{}
	params.Add("format", "json")
	params.Add("api_key", m.cfg.ApiMailer)
	params.Add("email", email)
	params.Add("sender_name", "Coupon shop")
	params.Add("sender_email", "hjadsfbnajv@gmail.com")
	params.Add("subject", "Купон")
	params.Add("body", fmt.Sprintf("Ваш купон: %s\nНазовите его администратору магазина для активации!", code))
	params.Add("list_id", "1")
	params.Add("lang", "ru")

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		log.Printf("parse request uri failed: %v", err)
	}
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	_, err = http.Get(urlStr)
	if err != nil {
		log.Printf("get failed: %v", err)
	}
}
