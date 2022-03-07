package main

import "gorm.io/gorm"

type Shelf struct {
	gorm.Model
	UserID uint
	Name   string
}
