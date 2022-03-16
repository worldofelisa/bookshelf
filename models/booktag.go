package main

import "gorm.io/gorm"

type BookTag struct {
	gorm.Model
	TagID  uint
	BookID uint
}

func (bt *BookTag) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&bt)
}

func (bt *BookTag) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&bt).Find(&bt)
}

func (bt *BookTag) Update(conn *gorm.DB) *gorm.DB {
	return conn.Save(&bt)
}

func (bt *BookTag) Delete(conn *gorm.DB) *gorm.DB {
	return conn.Delete(&bt)
}
