package mqutil

import (
	"encoding/json"
	"github.com/soxft/waline-async-mail/library/engine"
	"github.com/soxft/waline-async-mail/library/mail"
	"log"
	"os"
)

// Mail
// @description: 邮件发送相关
func Mail(msg string) {
	log.SetOutput(os.Stdout)

	var _mail engine.Mail
	if err := json.Unmarshal([]byte(msg), &_mail); err != nil {
		log.Panic(err)
	}
	if _mail.ToAddress == "" {
		log.Printf("[ERROR] Mail(%s) 空收件人", _mail.Typ)
		return
	}
	log.Printf("[INFO] Mail send [%s]: %s", _mail.Typ, _mail.ToAddress)

	// send mail
	if err := mail.Send(_mail, engine.GetSendPlatform(_mail)); err != nil {
		log.Panic(err)
	}
}
