package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
  	http.Handle("/", fs)

    testHandler := ReqHandler{ path: "/entry" }
    testHandler.handle(getEntry, postEntry)

    port := ":8080"
  	fmt.Printf("Listening on %s...\n", port)
	http.ListenAndServe(port, nil)
}

func getEntry() {
    fmt.Println("Get entry")
}

func postEntry() {
    fmt.Println("Post entry")
}

type ReqHandler struct {
    path string
}

func (r *ReqHandler) handle(get func(), post func()) {
    http.HandleFunc(r.path, func(res http.ResponseWriter, req *http.Request) {
        m := req.Method
        fmt.Printf("Method %q\n", m)
        switch m {
            case "GET": get()
            case "POST": post()
        }
    })
}