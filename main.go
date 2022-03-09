package main

import (
	"fmt"
	_ "image/jpeg"
)

func main() {
	barcode := scanImage("./IMG_4779.jpg")
	bookInfo := getBookInfo(barcode)
	//fmt.Println(string(bookInfo))
	bookData := parseBookJson(bookInfo)
	authInfo := bookData.Authors
	returnedAuthors := []APIAuthor{}
	for _, author := range authInfo {
		//adding author information to a slice of new authors
		returnedAuthors = append(returnedAuthors, parseAuthInfo(getAuthorInfo(author)))
	}

	coverPicURL(barcode)
	fmt.Println(bookData.Title, bookData.ISBN10, bookData.NumberOfPages, bookData.Covers, bookData.Series)

	conn := connectToDB()

	//migrates the DB tables!
	migrateDB(conn)
	//seedGenres(conn)
	//
	//addABook(conn, bookData, returnedAuthors, barcode, 9)

	//user := User{}
	//user.Name = "Elisa"
	//user.Password = "123456"
	//user.Email = "elisa@elisa.com"
	//user.Create(conn)
	//user.Retrieve(conn)
	//fmt.Println(user.Name, user.Password, user.Email)
	//user.Name = "Princess"
	//user.Update(conn)
	////newUserInfo := user.Retrieve(conn)
	////rowsAddedResponse(newUserInfo.RowsAffected)
	////printErrorHandler(newUserInfo.Error)
	////fmt.Println(user.Name, user.Password, user.Email)
	//user.Delete(conn)

	pages := PageTracker{}
	//TODO ask Drew to show/explain how to link the BookID/UserID then update table
	pages.CurrentPage = 57
	//pages.Create(conn)
	//pages.Retrieve(conn)
	//fmt.Println(pages.CurrentPage)
	//pages.Update(conn)

}
