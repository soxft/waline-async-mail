package mail

import (
	"crypto/tls"
	"fmt"
	"github.com/soxft/waline-async-mail/config"
	"gopkg.in/gomail.v2"
	"mime"
)

func sendBySmtp(mail Mail) error {
	m := gomail.NewMessage()

	senderNameUtf8 := mime.QEncoding.Encode("utf-8", config.BlogInfo.Title)
	m.SetHeader("From", fmt.Sprintf("\"%s\" <%s>", senderNameUtf8, config.Smtp.User)) // 发件人
	m.SetHeader("To", mail.ToAddress)                                                 // 收件人
	m.SetHeader("Subject", mail.Subject)                                              // 邮件主题

	m.SetBody("text/html; charset=UTF-8", mail.Content)

	d := gomail.NewDialer(
		config.Smtp.Host,
		config.Smtp.Port,
		config.Smtp.User,
		config.Smtp.Pwd,
	)
	if !config.Smtp.Secure {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
