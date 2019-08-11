package main

import server "./mylib"
import "fmt"
import http "net/http"
import "html"

func main() {
	fmt.Println("Hello World")
	serve()
}

func serve() {
	http.HandleFunc("/foo", server.FooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	_ = http.ListenAndServe(":8080", nil)
}
