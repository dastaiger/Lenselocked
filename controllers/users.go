package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

type Users struct {
	NewView *view.View
}

//NewUsers can panic -> only use this function during setup
func NewUsers() *Users {
	return &Users{
		NewView: view.NewView("bootstrap", "views/users/new.gohtml"),
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
	r.ParseForm()

	fmt.Fprintf(w, "this is a Temp Mesage: %v, %v", r.PostFormValue("email"), r.PostFormValue("Password"))
}
