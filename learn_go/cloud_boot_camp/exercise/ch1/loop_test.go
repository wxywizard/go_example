package ch1

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	fullString := "hello world"
	for index, value := range fullString {
		fmt.Println(index, string(value))
	}
}
