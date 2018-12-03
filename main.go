package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"lenslocked.com/views"
)

var (
	homeTemplate    *view.View
	contactTemplate *view.View
	faqTemplate     *view.View
)

func main() {
	//var err error
	homeTemplate = view.NewView("cover", "views/home.gohtml")
	contactTemplate = view.NewView("bootstrap", "views/contact.gohtml")
	faqTemplate = view.NewView("bootstrap", "views/faq.gohtml")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r) //nil = use buildin MUX

}

func contact(w http.ResponseWriter, r *http.Request) {
	contactTemplate.Render(w, nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Render(w, nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
	faqTemplate.Render(w, nil)
}
