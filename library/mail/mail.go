package mail

func Send(mail Mail, platform Platform) error {
	switch platform {
	case PlatformSmtp:
		return sendBySmtp(mail)
	default:
		return sendBySmtp(mail)
	}
}
