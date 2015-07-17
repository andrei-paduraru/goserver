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
        fmt.Printf("Method %q\n", m)
        var test bool
        test = put != nil
        fmt.Printf("%v\n", test)
        switch m {
            case "GET": get()
            case "POST": post()
            case "PUT": if put != nil {put()}else{handleError(res, 404, 1)}
            case "DELETE": delete()
        }
    })
}

func handleError(res http.ResponseWriter, status int, code int) {
    codes := []string{
        "Server error",
        "Not found"}
    var statuses = map[int]int{
        404: http.StatusNotFound}
    //if code == nil {code = 0}
    http.Error(res, codes[code], statuses[status])
}