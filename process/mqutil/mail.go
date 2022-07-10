package mqutil

import (
	"encoding/json"
	"github.com/soxft/waline-async-mail/library/mail"
	"log"
)

// Mail
// @description: 邮件发送相关
func Mail(msg string) {
	var mailMsg mail.Mail
	if err := json.Unmarshal([]byte(msg), &mailMsg); err != nil {
		log.Panic(err)
	}
	if mailMsg.ToAddress == "" {
		log.Printf("[ERROR] Mail(%s) 空收件人", mailMsg.Typ)
		return
	}
	log.Printf("[INFO] Mail(%s) %s", mailMsg.Typ, mailMsg.ToAddress)

	// get mail platform
	var platform mail.Platform
	switch mailMsg.Typ {
	case "register":
		platform = mail.PlatformSmtp
	default:
		platform = mail.PlatformSmtp
	}
	// send mail
	if err := mail.Send(mailMsg, platform); err != nil {
		log.Panic(err)
	}
}
