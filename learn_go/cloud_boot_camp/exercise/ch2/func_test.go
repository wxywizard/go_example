package ch2

import (
	"fmt"
	"os"
	"testing"
)

func TestFunc(t *testing.T) {
	err, res := DuplicateString("aaa")
	if err == nil {
		t.Log(res)
	} else {
		t.Log(err)
	}

	args := os.Args
	if len(args) != 0 {
		t.Log("Do not accept any argument")
		//os.Exit(1)
	}

	DoOperation(2, decrease)
}

func DuplicateString(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input + input
}

func increase(a, b int) int {
	return a + b
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) {
	println("decrease result is :", a-b)
	//闭包，匿名函数
	defer func() {
		if r := recover(); r != nil {
			println("recovered in FuncX")
		}
	}()
}
