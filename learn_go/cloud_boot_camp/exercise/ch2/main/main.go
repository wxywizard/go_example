package main

import "fmt"

type ParameterStruct struct {
	Name string
}

func changeParameter(para ParameterStruct, value string) {
	para.Name = value
	fmt.Println(para)
}

func main() {
	para := ParameterStruct{Name: "123"}
	fmt.Println(para)
	changeParameter(para, "bbb")
	fmt.Println(para)
}
