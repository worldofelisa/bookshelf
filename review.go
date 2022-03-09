package main

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	BookID    uint
	UserID    uint
	StarValue float64 `gorm:"index"`
}

// Create the review for a book
func (r *Review) Create(conn *gorm.DB, bookID, userID uint) {
	bookErr := conn.Model(&r).Association("BookID").Append(&Review{BookID: bookID})
	exitErrorHandler(bookErr)
	userErr := conn.Model(&r).Association("UserID").Append(&Review{UserID: userID})
	exitErrorHandler(userErr)
	rResult := conn.Create(&r)
	rowsAddedResponse(rResult.RowsAffected)
	printErrorHandler(rResult.Error)
}

// Retrieve views the review information
func (r *Review) Retrieve(conn *gorm.DB) {
	viewReview := conn.Where(&r).Find(&r)
	rowsAddedResponse(viewReview.RowsAffected)
	printErrorHandler(viewReview.Error)
}

// Update review information
func (r *Review) Update(conn gorm.DB) {
	reviewValue := conn.Save(&r)
	rowsAddedResponse(reviewValue.RowsAffected)
	printErrorHandler(reviewValue.Error)
}
