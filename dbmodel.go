package main

import "gorm.io/gorm"

type Model interface {
	Create(conn *gorm.DB) *gorm.DB
	Retrieve(conn *gorm.DB) *gorm.DB
	Update(conn *gorm.DB) *gorm.DB
	Delete(conn *gorm.DB) *gorm.DB
}
