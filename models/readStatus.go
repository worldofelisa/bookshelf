package main

import "gorm.io/gorm"

type ReadStatus struct {
	gorm.Model
	Status           bool
	CurrentlyReading bool
	DNF              bool
	BookID           uint
	UserID           uint
}

// Create
//addReadStatus adds whether the book is currently being read to the DB table.
//TODO information to come in via code on user end via tick boxes
func (rs *ReadStatus) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&rs)
}

// Retrieve
func (rs *ReadStatus) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&rs).Find(&rs)
}

// Update
func (rs *ReadStatus) Update(conn *gorm.DB) *gorm.DB {
	return conn.Save(&rs)
}

func (rs *ReadStatus) Delete(conn *gorm.DB) *gorm.DB {
	return conn.Delete(&rs)
}
