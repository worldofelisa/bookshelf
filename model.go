package main

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN       string `gorm:"unique"`
	Key        string
	Title      string `gorm:"index;<-:create"`
	Series     string
	GenreID    uint
	PageNumber int
}

type Author struct {
	gorm.Model
	Key  string
	Name string `gorm:"index;<-:create"`
}

type BookAuthor struct {
	gorm.Model
	BookID   uint
	AuthorID uint
}

type Genre struct {
	gorm.Model
	Name string `gorm:"index;<-:create"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"index;<-:create"`
}

type BookTag struct {
	gorm.Model
	TagID  uint
	BookID uint
}

type Shelf struct {
	gorm.Model
	UserID uint
	Name   string
}

type ShelfBook struct {
	gorm.Model
	ShelfID uint
	BookID  uint
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;index"`
	Password string
}

type Review struct {
	gorm.Model
	BookID    uint
	UserID    uint
	StarValue float64 `gorm:"index"`
}

type ReadStatus struct {
	gorm.Model
	Status           bool
	CurrentlyReading bool
	DNF              bool
	BookID           uint
	UserID           uint
}

type PageTracker struct {
	gorm.Model
	BookID      uint
	UserID      uint
	CurrentPage uint
}
