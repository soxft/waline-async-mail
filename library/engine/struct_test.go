package engine

import (
	"log"
	"reflect"
	"testing"
)

func TestStruct(t *testing.T) {
	r := reflect.TypeOf(GuestArgs{})
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		log.Println(field.Name, field.Tag.Get("args"))
	}
}
