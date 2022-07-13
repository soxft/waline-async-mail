package mqutil

import (
	"encoding/json"
	"github.com/soxft/waline-async-mail/library/mail"
	"log"
	"os"
)

// Mail
// @description: 邮件发送相关
func Mail(msg string) {
	log.SetOutput(os.Stdout)

	var mailMsg mail.Mail
	if err := json.Unmarshal([]byte(msg), &mailMsg); err != nil {
		log.Panic(err)
	}
	if mailMsg.ToAddress == "" {
		log.Printf("[ERROR] Mail(%s) 空收件人", mailMsg.Typ)
		return
	}
	log.Printf("[INFO] Mail send [%s]: %s", mailMsg.Typ, mailMsg.ToAddress)

	// send mail
	if err := mail.Send(mailMsg, mail.PlatformSmtp); err != nil {
		log.Panic(err)
	}
}
