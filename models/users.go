package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ErrNotFound = errors.New("Models: Resource not fround!")

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique_index;not null"`
}

type UserService struct {
	db *gorm.DB
}

//ByID will look for a User with the ID provided
func (us *UserService) ByID(ID uint) (*User, error) {
	var user User
	_, err := First(us.db.Where("ID = ?", ID), &user)
	return &user, err

}

//ByEmail will return a *User with the provided Email
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	_, err := First(us.db.Where("Email = ?", email), &user)
	return &user, err

}

//ByName will return a *User with the provided Name
func (us *UserService) ByName(name string) (*User, error) {
	var user User
	_, err := First(us.db.Where("Name = ?", name), &user)
	return &user, err
}

//First is the db.grom.First method with error handling attachted to it.
// It returns a User and nil in the best case or Nil, Error or Nil, ErrRecordNotFound in case of an error.
func First(db *gorm.DB, user *User) (*User, error) {
	err := db.First(&user).Error
	switch err {
	case nil:
		return user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//NewUserService will create a new UserService struct with an active DB session inside
func NewUserService(connectionString string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionString)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	return &UserService{db: db}, nil
}

//Destrcution will drop the table if it exists and then AutoMigrate the clear DB
func (us *UserService) Destrcution() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

//Create will create an user
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

//Update will Update the Information of an user
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

//Delete will delete the an user
func (us *UserService) Delete(user *User) error {
	return us.db.Where("Email = ?", user.Email).Delete(&user).Error
}
