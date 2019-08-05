package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	current_time := time.Now().Local()
	fmt.Fprintf(w, "Hello, You're running zarvis-example-go! %s\n\n", current_time.Format("2006-01-02 15:04:05"))
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":12222", nil)
}
