package mail

import (
	"github.com/soxft/waline-async-mail/config"
	"log"
	"reflect"
)

func ParseOwner(args OwnerArgs) (string, error) {
	_ = config.OwnerTemplate
	r := reflect.TypeOf(args)
	v := reflect.ValueOf(args)
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		log.Println(v.Field(i), field.Tag.Get("args"))
	}

	return "", nil
}

func ParseGuest(args GuestArgs) (string, error) {
	_ = config.GuestTemplate

	return "", nil
}
