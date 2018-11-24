package main

import (
	"html/template"
	"os"
)

type user struct {
	Name   string
	Dog    string
	Alter  int
	Doof   bool
	Frauen map[string]int
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := user{
		Name:   "WurstMann",
		Dog:    "Bello",
		Alter:  25,
		Doof:   true,
		Frauen: map[string]int{"Helga": 20, "lola": 30},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
