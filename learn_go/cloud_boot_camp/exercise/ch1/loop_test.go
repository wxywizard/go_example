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

func TestExec(t *testing.T) {
	myArray := [5]string{"I", "am", "stupid", "and", "weak"}
	for index, _ := range myArray {
		if index == 2 {
			myArray[2] = "smart"
		}
		in := len(myArray) - 1
		if index == in {
			myArray[in] = "strong"
		}
	}
	t.Log(myArray)
}
