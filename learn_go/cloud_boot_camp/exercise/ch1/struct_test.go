package ch1

import (
	"reflect"
	"testing"
)

type MyType struct {
	Name string `json:"name"`
}

func TestStruct(t *testing.T) {
	mt := MyType{Name: "test"}
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")
	t.Logf("%s\n", tag)
}
