package main

import (
	"fmt"
	_ "image/jpeg"
)

func main() {
	barcode := scanImage("./IMG_4779.jpg")
	bookInfo := getBookInfo(barcode)
	bookData := parseBookJson(bookInfo)
	authInfo := bookData.Authors
	returnedAuthors := []APIAuthor{}
	for _, author := range authInfo {
		//adding author information to a slice of new authors
		returnedAuthors = append(returnedAuthors, parseAuthInfo(getAuthorInfo(author)))
	}

	authName := []string{}

	for _, name := range returnedAuthors {
		//takes the response of the returned name and only selects the name from it
		//output is then placed into an array/slice but is readable as the name
		authName = append(authName, name.Name)

	}
	coverPicURL(barcode)
	fmt.Println(authName, bookData.Title, bookData.ISBN10, bookData.NumberOfPages, bookData.Covers, bookData.Series)

	//conn = connectDB
	conn := connectToDB()

	//migrates the DB tables!
	migrateDB(conn)

	//add book to DB table
	addBook(conn)
}
