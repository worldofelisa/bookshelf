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
func (u *User) Create(conn *gorm.DB) {
	userResult := conn.Create(&u)
	rowsAddedResponse(userResult.RowsAffected)
	printErrorHandler(userResult.Error)
}

// Retrieve gets information from the db table
func (u *User) Retrieve(conn *gorm.DB) {
	userInfo := conn.Where(&u).Find(&u)
	rowsAddedResponse(userInfo.RowsAffected)
	printErrorHandler(userInfo.Error)
}

// Update retrieves information on the user and then replaces the name / password as needed
func (u *User) Update(conn *gorm.DB) {
	userUpdate := conn.Save(&u)
	rowsAddedResponse(userUpdate.RowsAffected)
	printErrorHandler(userUpdate.Error)
}

//Delete creates a soft delete - user is still in db table but can not be located with normal queries
//to query deleted users use Unscoped.
func (u *User) Delete(conn *gorm.DB) {
	userStatus := conn.Delete(&u)
	rowsAddedResponse(userStatus.RowsAffected)
	printErrorHandler(userStatus.Error)
}
