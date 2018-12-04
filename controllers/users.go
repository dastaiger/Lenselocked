package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"lenslocked.com/views"
)

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.
var decoder = schema.NewDecoder()

type Users struct {
	NewView *view.View
}

type SignUp struct {
	Email    string `schema:"email,required"`
	Password string `schema:"password,required"`
}

//NewUsers can panic -> only use this function during setup
func NewUsers() *Users {
	return &Users{
		NewView: view.NewView("bootstrap", "users/new"),
	}
}

// New shows the User Signup Form
//
// /Signup GET
func (u *Users) New(w http.ResponseWriter, r *http.Request) {

	u.NewView.Render(w, nil)
}

// Create handles the creation of a new User
//
// /Signup POST
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {

	var user SignUp

	if err := ParseForm(r, &user); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "this is a Temp Mesage: %v, %v \n %v", r.PostFormValue("email"), r.PostFormValue("Password"), user)
}
