package main

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"unique;index;<-:create"`
}

// Create adds  tags to Books
func (t *Tag) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&t)
}

// Retrieve see the tags on a book
func (t *Tag) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&t).Find(&t)
}

func (t *Tag) Update(conn *gorm.DB) *gorm.DB {
	return conn.Save(&t)
}

func (t *Tag) Delete(conn *gorm.DB) *gorm.DB {
	return conn.Delete(&t)
}
