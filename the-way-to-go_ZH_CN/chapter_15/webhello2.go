package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shouthello/"):]
	fmt.Fprintf(w, "hello %s!", strings.ToUpper(remPartOfURL))
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	http.ListenAndServe("localhost:9999", nil)
}
