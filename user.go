package main

import (
	"gorm.io/gorm"
)

//addNewUser adds a new user to the database
func addNewUser(conn *gorm.DB, name, password, email string) {
	userResult := conn.Create(&User{Name: name, Password: password, Email: email})
	rowsAddedResponse(userResult.RowsAffected)
	printErrorHandler(userResult.Error)
}

// RETRIEVE

//UPDATE
func updateUserInfo() {

}

//DELETE
