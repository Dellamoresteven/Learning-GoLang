package server

import "fmt"
import "strings"
import http "net/http"

func FooHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Start FooHandler")
    defer fmt.Println("End FooHandler")
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}
