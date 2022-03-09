package main

import "gorm.io/gorm"

type PageTracker struct {
	gorm.Model
	BookID      uint
	UserID      uint
	CurrentPage uint
}

// Create adds a page tracker to a book and user
func (pt *PageTracker) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&pt)

}

// Retrieve gets info from db and allows us to view it
func (pt *PageTracker) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&pt).Find(&pt)

}

// Update saves the information that has changed.
func (pt *PageTracker) Update(conn *gorm.DB) *gorm.DB {
	return conn.Save(&pt)

}
