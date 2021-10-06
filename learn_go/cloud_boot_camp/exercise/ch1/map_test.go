package ch1

import (
	"testing"
)

func TestGeneratorMap(t *testing.T) {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myFuncMap := map[string]func() int{
		"funcA": func() int {
			return 1
		},
	}
	t.Log(myMap)
	f := myFuncMap["funcA"]
	t.Log(f())
	value, exists := myMap["a"]
	if exists {
		println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}
}
