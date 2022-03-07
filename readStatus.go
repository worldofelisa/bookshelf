package main

import "gorm.io/gorm"

//CREATE
//addReadStatus adds whether the book is currently being read to the DB table.
//TODO information to come in via code on user end via tick boxes
func addReadStatus(conn *gorm.DB, read bool, currentlyReading bool, dnf bool) {
	readStatus := ReadStatus{Status: read, CurrentlyReading: currentlyReading, DNF: dnf}
	readResult := conn.Create(&readStatus)
	rowsAddedResponse(readResult.RowsAffected)
	printErrorHandler(readResult.Error)
}

//RETRIEVE
func readReadStatus(conn *gorm.DB, ID uint) {
	conn.Take(&ReadStatus{BookID: ID})
}

//UPDATE
func updateReadStatus(conn *gorm.DB, read bool, currentlyReading bool, dnf bool) {
	//Take a specific book
	conn.Take(ReadStatus{})
	//update the read information for that book
	conn.Save(&ReadStatus{Status: read, CurrentlyReading: currentlyReading, DNF: dnf})
}
