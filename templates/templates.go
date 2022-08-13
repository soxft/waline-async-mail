package templates

import (
	_ "embed"
	"io/ioutil"
)

var (
	//go:embed owner.html
	MailOwner string
	//go:embed guest.html
	MailGuest string

	//go:embed owner.tpl
	TextOwner string
	//go:embed guest.tpl
	TextGuest string
)

func init() {
	// get file content
	if _ownerTpl, err := ioutil.ReadFile("./templates/owner.html"); err == nil {
		MailOwner = string(_ownerTpl)
	}
	if _guestTpl, err := ioutil.ReadFile("./templates/guest.html"); err == nil {
		MailGuest = string(_guestTpl)
	}

	if _ownerTplText, err := ioutil.ReadFile("./templates/owner.tpl"); err == nil {
		TextOwner = string(_ownerTplText)
	}
	if _guestTplText, err := ioutil.ReadFile("./templates/guest.tpl"); err == nil {
		TextGuest = string(_guestTplText)
	}
}
