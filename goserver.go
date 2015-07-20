package main

import (
	"fmt"
	"net/http"
	"github.com/andrei-paduraru/httputils"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
  	http.Handle("/", fs)

    entryHandler := httputils.ReqHandler{ Path: "/entry" }
    entryHandler.Handle(getEntry, postEntry, nil, nil)

    port := ":8080"
  	fmt.Printf("Listening on %s...\n", port)
	http.ListenAndServe(port, nil)
}

func getEntry(res http.ResponseWriter, req *http.Request) {
    fmt.Println("Get entry")
    fmt.Println(req.URL)
}

func postEntry(res http.ResponseWriter, req *http.Request) {
    fmt.Println("Post entry")
    fmt.Println(req.URL)
}
