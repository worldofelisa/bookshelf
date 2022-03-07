package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;index"`
	Password string
}

// Create adds a new user to the database
func (u *User) Create(conn *gorm.DB, name, password, email string) {
	userResult := conn.Create(&User{Name: name, Password: password, Email: email})
	rowsAddedResponse(userResult.RowsAffected)
	printErrorHandler(userResult.Error)
}

// Retrieve gets information from the db table
func (u *User) Retrieve(conn *gorm.DB, email string) *gorm.DB {
	user := User{}
	return conn.Where(&User{Email: email}).Find(&user)
}

//UPDATE
func updateUserInfo() {

}

//DELETE
