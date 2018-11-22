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

	if r.URL.Path == "/contact" {
		fmt.Fprint(w, "contact me @ <a href=\"mailto:bla@spam.de\"> mailaddr </a>.")
	}
	fmt.Fprint(w, "<h1> hiho </h1>")
}
