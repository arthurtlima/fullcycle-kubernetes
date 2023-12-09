package main

import (
	"net/http"
	"fmt"
	"os"
)

func main () {
	http.HandleFunc("/", Hello) 
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I`m %s and I`m %s years old", name, age)
}