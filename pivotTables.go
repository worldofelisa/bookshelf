package main

import "gorm.io/gorm"

type ShelfBook struct {
	gorm.Model
	ShelfID uint
	BookID  uint
}

//TODO - add CRUD for this or will this be automatically done through gorm/other code?