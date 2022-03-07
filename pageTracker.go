package main

import "gorm.io/gorm"

type PageTracker struct {
	gorm.Model
	BookID      uint
	UserID      uint
	CurrentPage uint
}

//CREATE

//VIEW

//UPDATE
