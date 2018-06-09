package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are in view")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
