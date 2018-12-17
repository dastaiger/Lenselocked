package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/models"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
)

//Remove in Prod!
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked"
)

func main() {
	connStr := fmt.Sprintf("host=%s user=%s port=%v dbname=%s password=%s sslmode=disable", host, user, port, dbname, password)
	us, err := models.NewUserService(connStr)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	usersC := controllers.NewUsers(us)
	staticC := controllers.NewStatic()

	us.Destrcution()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r) //nil = use buildin MUX

}
