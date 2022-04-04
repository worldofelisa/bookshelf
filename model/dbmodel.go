package model

import "gorm.io/gorm"

type DBModel interface {
	Create(conn *gorm.DB) *gorm.DB
	Retrieve(conn *gorm.DB) *gorm.DB
	Update(conn *gorm.DB) *gorm.DB
	Delete(conn *gorm.DB) *gorm.DB
}

func Create(conn *gorm.DB, model DBModel) *gorm.DB {
	return model.Create(conn)
}

func Retrieve(conn *gorm.DB, model DBModel) *gorm.DB {
	return model.Retrieve(conn)
}

func Update(conn *gorm.DB, model DBModel) *gorm.DB {
	return model.Update(conn)
}

func Delete(conn *gorm.DB, model DBModel) *gorm.DB {
	return model.Delete(conn)
}

func SeedDB(conn *gorm.DB, model DBModel) {
	result := conn.Where(&model).Find(&model)
	if result.RowsAffected == 0 {
		conn.Create(&model)
	}
}
