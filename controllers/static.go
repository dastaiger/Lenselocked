package controllers

import (
	"lenslocked.com/views"
)

//NewStatic returns a *Static with all our static pages
func NewStatic() *Static {
	return &Static{
		Home:    view.NewView("bootstrap", "static/home"),
		Contact: view.NewView("bootstrap", "static/contact"),
		Faq:     view.NewView("bootstrap", "static/faq"),
	}
}

//Static holds the definition of all our static pages
type Static struct {
	Home    *view.View
	Contact *view.View
	Faq     *view.View
}
