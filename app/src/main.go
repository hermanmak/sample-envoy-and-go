package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/service", handle)
        http.HandleFunc("/ping", ping)
	http.ListenAndServe(":80", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World SERVICE!")
}


func ping(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Server","A Go Web Server")
        w.WriteHeader(200)
}
