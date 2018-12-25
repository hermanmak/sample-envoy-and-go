package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/service", handle)
	http.ListenAndServe(":7000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World SERVICE!")
}
