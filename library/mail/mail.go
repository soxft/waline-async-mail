package mail

import "github.com/soxft/waline-async-mail/config"

func Send(mail Mail, platform Platform) error {
	switch platform {
	case PlatformSmtp:
		return sendBySmtp(mail)
	case PlatformAliyun:
		return sendByAliyun(mail)
	default:
		return sendBySmtp(mail)
	}
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
	} else {
		return PlatformSmtp
	}
}
