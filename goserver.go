package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
  	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
  	fmt.Printf("Listening...")
}
