package engine

import (
	"github.com/soxft/waline-async-mail/config"
)

func Parse(platform Platform) error {

	return nil
}

func parseMail() error {
	return nil
}

func parseText() error {
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

	return strToPlatform[sendTyp]
}

func GetParseBy(platform Platform) ParseBy {
	if platform == PlatformSmtp || platform == PlatformAliyun {
		return ParseByMail
	} else {
		return ParseByText
	}
}
