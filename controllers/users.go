package controllers

import (
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

func (u *Users) New(w http.ResponseWriter, r *http.Request) {

	u.NewView.Render(w, nil)
}
