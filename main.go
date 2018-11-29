package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

var (
	homeTemplate    *view.View
	contactTemplate *view.View
)

func main() {
	//var err error
	homeTemplate = view.NewView("bootstrap",
		"views/home.gohtml",
	)

	contactTemplate = view.NewView("bootstrap",
		"views/contact.gohtml",
	)
	r := mux.NewRouter()
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r) //nil = use buildin MUX

}

func contact(w http.ResponseWriter, r *http.Request) {
	contactTemplate.Render(w, nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Render(w, nil)
}
