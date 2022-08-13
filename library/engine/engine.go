package engine

import (
	"github.com/soxft/waline-async-mail/config"
)

func ParseMail() error {
	return nil
}

func ParseText() error {
	return nil
}

// GetSendPlatform get platform
// 通过 mail.Typ 获取发送平台
func GetSendPlatform(mail Mail) Platform {
	var sendTyp string
	if mail.Typ == "toGuest" {
		sendTyp = config.SendBy.Guest
	} else {
		sendTyp = config.SendBy.Owner
	}

	if sendTyp == "smtp" {
		return PlatformSmtp
	} else if sendTyp == "aliyun" {
		return PlatformAliyun
	} else if sendTyp == "bark" {
		return PlatformBark
	} else if sendTyp == "telegram" {
		return PlatformTelegram
	} else {
		return PlatformSmtp
	}
}
