package controllers

import (
	"net/http"
)

//ParseForm parses forms and returns the error
func ParseForm(r *http.Request, dst interface{}) error {

	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
