package templates

import (
	_ "embed"
	"io/ioutil"
)

var (
	//go:embed owner.html
	Owner string
	//go:embed guest.html
	Guest string
)

func init() {
	// get file content
	if _ownerTemplate, err := ioutil.ReadFile("./templates/owner.html"); err == nil {
		Owner = string(_ownerTemplate)
	}
	if _guestTemplate, err := ioutil.ReadFile("./templates/guest.html"); err == nil {
		Guest = string(_guestTemplate)
	}
}
