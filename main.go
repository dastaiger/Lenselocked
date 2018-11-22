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
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Host", "Wurst")
	fmt.Fprint(w, "<h1> Welcome to great my Site! </h1>")
}
