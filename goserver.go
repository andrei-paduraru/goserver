package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
  	http.Handle("/", fs)

    entryHandler := ReqHandler{ path: "/entry" }
    entryHandler.handle(getEntry, postEntry, nil, nil)

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

func (r *ReqHandler) handle(get func(), post func(), put func(), delete func()) {
    http.HandleFunc(r.path, func(res http.ResponseWriter, req *http.Request) {
        m := req.Method
        var status int
        status = http.StatusOK
        fmt.Printf("%s %s ", m, r.path)
        var call func()
        switch m {
            case "GET": if get != nil {call = get}else{status = http.StatusNotFound}
            case "POST": if post != nil {call = post}else{status = http.StatusNotFound}
            case "PUT": if put != nil {call = put}else{status = http.StatusNotFound}
            case "DELETE": if delete != nil {call = delete}else{status = http.StatusNotFound}
        }
        fmt.Printf("(%d)\n", status)
        if status == http.StatusNotFound {
            http.Error(res, "Not found", http.StatusNotFound)
        }else{
            call()
        }
    })
}