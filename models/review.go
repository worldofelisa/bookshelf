package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	BookID    uint
	UserID    uint
	StarValue float64 `gorm:"index"`
}

// Create the review for a book
func (r *Review) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&r)
}

// Retrieve views the review information
func (r *Review) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&r).Find(&r)
}

// Update review information
func (r *Review) Update(conn gorm.DB) *gorm.DB {
	return conn.Save(&r)
}

func (r *Review) Delete(conn *gorm.DB) *gorm.DB {
	return conn.Delete(&r)
}
