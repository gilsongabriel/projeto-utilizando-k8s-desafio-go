package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	_ = http.ListenAndServe("0.0.0.0:8000", nil)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write(greeting("Code.education Rocks!"));
}

func greeting(msg string) []byte {
	return []byte(fmt.Sprintf("<b>%s</b>", msg))
}