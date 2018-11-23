package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r) //nil = use buildin MUX

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "contact me @ <a href=\"mailto:bla@spam.de\"> mailaddr </a>.")

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> hiho, this is home :)</h1>")
}
