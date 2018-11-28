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
	homeTemplate = view.NewView(
		"views/home.gohtml",
	)

	contactTemplate = view.NewView(
		"views/contact.gohtml",
	)
	r := mux.NewRouter()
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r) //nil = use buildin MUX

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Template.ExecuteTemplate(w, "yield", nil); err != nil {
		panic(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Template.ExecuteTemplate(w, "yield", nil); err != nil {

		panic(err)
	}
}
