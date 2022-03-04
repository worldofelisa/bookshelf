package main

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN       string
	Key        string
	Title      string `gorm:"index"`
	Series     string
	GenreID    uint
	PageNumber int
}

type Author struct {
	gorm.Model
	Key  string
	Name string `gorm:"index"`
}

type BookAuthor struct {
	gorm.Model
	BookID   uint
	AuthorID uint
}

type Genre struct {
	gorm.Model
	Name string `gorm:"index"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"index"`
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
	Email    string
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
