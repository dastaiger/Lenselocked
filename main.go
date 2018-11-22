package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil) //nil = use buildin MUX

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone was on the page!")
	fmt.Fprint(w, "<h1> Welcome to great my Site! </h1>")
}
