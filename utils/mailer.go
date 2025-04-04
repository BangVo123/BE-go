package utils

import (
	mail "gopkg.in/gomail.v2"
)

func SendMail(dialer *mail.Dialer, from, to, code string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Verify code")
	msg.SetBody("text/plain", "This is my verify code: "+code+". It will be expired after 5 minutes")

	return dialer.DialAndSend(msg)
}
