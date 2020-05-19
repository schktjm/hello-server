package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Listening localhost:8080")
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/a", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Go.")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := strings.TrimLeft(r.URL.Path, "/")
	fmt.Fprintf(w, "Hello %v\n", name)
	fmt.Printf("%v\n", formattingRequest(r))
}

func formattingRequest(r *http.Request) string {
	var request []string
	request = append(request, fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto))
	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	return strings.Join(request, "\n")
}
