package mail

import "github.com/soxft/waline-async-mail/library/engine"

// Send notification by mail
func Send(mail engine.Mail, platform engine.Platform) error {
	switch platform {
	case engine.PlatformSmtp:
		return sendBySmtp(mail)
	case engine.PlatformAliyun:
		return sendByAliyun(mail)
	default:
		return sendBySmtp(mail)
	}
}
