package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	UserID uint
	Name   string
	Price  uint
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique_index;not null"`
	Color string
	Order []Product
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked"
)

func main() {

	connStr := fmt.Sprintf("host=%s user=%s port=%v dbname=%s password=%s sslmode=disable", host, user, port, dbname, password)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	u := []User{}
	// db.Where("Name LIKE ?", "%Bort%").Find(&u)

	// p := Product{
	// 	UserID: u[0].ID,
	// 	Name:   "Haus",
	// 	Price:  120215412,
	// }
	// db.Where("Name LIKE ?", "%Borg%").Find(&u)
	// p1 := Product{
	// 	UserID: u[0].ID,
	// 	Name:   "Tier",
	// 	Price:  125214,
	// }
	// db.Create(&p)
	// db.Create(&p1)
	db.AutoMigrate(&User{}, &Product{})
	//db.Set("gorm:auto_preload", true).Find(&User{})

	if err := db.Preload("Order").Find(&u).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("Record not Found!")

		}
	}
	var results []User
	db.Raw("SELECT * FROM users").Scan(&results)
	fmt.Printf("%v", results)
}
