package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked"
)

func main() {

	connStr := fmt.Sprintf("host=%s user=%s port=%v dbname=%s password=%s sslmode=disable", host, user, port, dbname, password)

	user := models.User{
		Name:  "Hans Dampf",
		Email: "hans@dampfmail.com",
	}
	us, _ := models.NewUserService(connStr)
	us.Destrcution()
	us.Create(&user)
	var ReadUser *models.User
	ReadUser, _ = us.ByEmail("hans@dampfmail.com")
	fmt.Println(ReadUser.Name)
	ReadUser, _ = us.ByName("Hans Dampf")
	println(ReadUser.Name)
	ReadUser, _ = us.ByID(1)
	println(ReadUser.Name)

}
