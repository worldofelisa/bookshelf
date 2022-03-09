package main

import "gorm.io/gorm"

type PageTracker struct {
	gorm.Model
	BookID      uint
	UserID      uint
	CurrentPage uint
}

// Create adds a page tracker to a book and user
func (pt *PageTracker) Create(conn *gorm.DB, bookID, userID uint) {
	bookErr := conn.Model(&pt).Association("BookID").Append(&PageTracker{BookID: bookID})
	exitErrorHandler(bookErr)
	userErr := conn.Model(&pt).Association("UserID").Append(&PageTracker{UserID: userID})
	exitErrorHandler(userErr)
	ptResult := conn.Create(&pt)
	rowsAddedResponse(ptResult.RowsAffected)
	printErrorHandler(ptResult.Error)
}

// Retrieve gets info from db and allows us to view it
func (pt *PageTracker) Retrieve(conn *gorm.DB) {
	viewPages := conn.Where(&pt).Find(&pt)
	rowsAddedResponse(viewPages.RowsAffected)
	printErrorHandler(viewPages.Error)
}

// Update saves the information that has changed.
func (pt *PageTracker) Update(conn *gorm.DB) {
	pageStatus := conn.Save(&pt)
	rowsAddedResponse(pageStatus.RowsAffected)
	printErrorHandler(pageStatus.Error)
}
