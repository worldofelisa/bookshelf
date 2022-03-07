package main

import (
	"gorm.io/gorm"
)

type BookTag struct {
	gorm.Model
	TagID  uint
	BookID uint
}

type ShelfBook struct {
	gorm.Model
	ShelfID uint
	BookID  uint
}
