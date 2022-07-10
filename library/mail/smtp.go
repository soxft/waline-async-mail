package mail

import (
	"crypto/tls"
	"github.com/soxft/waline-async-mail/config"
	"gopkg.in/gomail.v2"
)

func sendBySmtp(mail Mail) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Smtp.SenderEmail) // 发件人
	m.SetHeader("To", mail.ToAddress)            // 收件人
	m.SetHeader("Subject", mail.Subject)         // 邮件主题

	m.SetBody("text/html; charset=UTF-8", mail.Content)

	if config.Smtp.Secure {
		return sendMailUsingTLS(m)
	} else {
		return sendSmtp(m)
	}
}

func sendSmtp(m *gomail.Message) error {
	d := gomail.NewDialer(
		config.Smtp.Host,
		config.Smtp.Port,
		config.Smtp.User,
		config.Smtp.Pwd,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func sendMailUsingTLS(m *gomail.Message) error {
	return sendSmtp(m)
}
