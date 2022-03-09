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
func (u *User) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&u)

}

// Retrieve gets information from the db table
func (u *User) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&u).Find(&u)
}

// Update retrieves information on the user and then replaces the name / password as needed
func (u *User) Update(conn *gorm.DB) *gorm.DB {
	return conn.Save(&u)
}

//Delete creates a soft delete - user is still in db table but can not be located with normal queries
//to query deleted users use Unscoped.
func (u *User) Delete(conn *gorm.DB) *gorm.DB {
	return conn.Delete(&u)

}
