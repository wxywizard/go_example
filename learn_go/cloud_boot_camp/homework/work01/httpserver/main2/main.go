package main

import (
	"fmt"
	"net/http"
)

type HelloHandlerStruct struct {
	content string
}

func (handler *HelloHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		//fmt.Printf("%s:%v\n",k,v)
		w.Header().Add(k, arrayToString(v))
	}
	fmt.Fprintf(w, handler.content)
}

func main() {
	http.Handle("/healthz", &HelloHandlerStruct{content: "Hello World"})
	http.ListenAndServe(":8000", nil)
}

func arrayToString(arr []string) string {
	var result string
	for _, i := range arr {
		result += i
	}
	return result
}
