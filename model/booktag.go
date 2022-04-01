package model

import "gorm.io/gorm"

type UserBookTag struct {
	gorm.Model
	TagID  uint
	BookID uint
	UserID uint
}

func (ubt *UserBookTag) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&ubt)
}

func (ubt *UserBookTag) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&ubt).Find(&ubt)
}

func (ubt *UserBookTag) Update(conn *gorm.DB) *gorm.DB {
	return conn.Save(&ubt)
}

func (ubt *UserBookTag) Delete(conn *gorm.DB) *gorm.DB {
	return conn.Delete(&ubt)
}
