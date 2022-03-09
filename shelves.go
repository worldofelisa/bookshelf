package main

import "gorm.io/gorm"

type Shelf struct {
	gorm.Model
	UserID uint
	Name   string
}

// Create makes a shelf
func (s *Shelf) Create(conn gorm.DB) {
	s.UserID = User{ID}
	shelfResult := conn.Create(conn)
	rowsAddedResponse(shelfResult.RowsAffected)
	printErrorHandler(shelfResult.Error)
}

// Retrieve the shelf information
func (s *Shelf) Retrieve(conn gorm.DB) *gorm.DB {
	return conn.Where(&s).Find(&s)
}

// Delete soft deletes shelf so it is no longer visible but still in db table
//deleted at on table is updated
//TODO add a 'content disables, re-enable feature for adding a shelf that had been deleted
func (s *Shelf) Delete(conn gorm.DB) {
	shelfStatus := conn.Delete(&s)
	rowsAddedResponse(shelfStatus.RowsAffected)
	printErrorHandler(shelfStatus.Error)
}
