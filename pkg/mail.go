package pkg

import (
	"github.com/alexeyco/unisender"
)

func CreateMailDialer(cfg Mailer) *unisender.UniSender {
	return unisender.New(cfg.ApiMailer).SetLanguageEnglish()
}
