package main

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	BookID    uint
	UserID    uint
	StarValue float64 `gorm:"index"`
}
