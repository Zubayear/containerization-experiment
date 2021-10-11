package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Hello, Friend!\n"))
	})
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
