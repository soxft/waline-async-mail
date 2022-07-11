package mail

import (
	"github.com/soxft/waline-async-mail/app/config"
	"testing"
)

func TestSmtp(t *testing.T) {
	mail := Mail{
		ToAddress: config.BlogInfo.AuthorEmail,
		Subject:   "测试邮件",
		Content:   "测试邮件内容",
		Typ:       "test",
	}
	if err := Send(mail, PlatformSmtp); err != nil {
		panic(err)
	}
}
