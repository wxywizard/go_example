package ch1

import (
	"fmt"
	"testing"
)

func TestMySlice(t *testing.T) {
	var mySlice []int
	intSlice := append(mySlice, 2, 3, 4, 5)
	resSlice := intSlice[1:2]
	t.Logf("%+v\n", resSlice)
}

func TestGeneratorSlice(t *testing.T) {
	mySlice1 := new([]int)
	mySlice2 := make([]int, 0)
	mySlice3 := make([]int, 10)
	mySlice4 := make([]int, 10, 20) //length初始化给多少，最长给多少

	t.Logf("mySlice1==%+v\n", mySlice1)
	t.Logf("mySlice2==%+v\n", mySlice2)
	t.Logf("mySlice3==%+v\n", mySlice3)
	t.Logf("mySlice4==%+v\n", mySlice4)
	t.Log(mySlice3)
}

//go 值传递
func TestValuePass(t *testing.T) {
	mySlice := []int{10, 20, 30, 40, 50}
	for _, value := range mySlice {
		value *= 2
	}
	t.Logf("%+v\n", mySlice)
	for index, _ := range mySlice {
		mySlice[index] *= 2
	}
	t.Logf("%+v\n", mySlice)
	fmt.Printf("%v\n", mySlice)
	fmt.Printf("%+v\n", mySlice)
}
