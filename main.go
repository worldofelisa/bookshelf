package main

import (
	_ "image/jpeg"
)

type BasicPageData struct {
	Title string
}

//func barcoding() {
//	barcode := scanImage("./IMG_4779.jpg")
//	bookInfo := getBookInfo(barcode)
//	bookData := parseBookJson(bookInfo)
//	authInfo := bookData.Authors
//	returnedAuthors := []APIAuthor{}
//	for _, author := range authInfo {
//		//adding author information to a slice of new authors
//		returnedAuthors = append(returnedAuthors, parseAuthInfo(getAuthorInfo(author)))
//	}
//
//	coverPicURL(barcode)
//	fmt.Println(bookData.Title, bookData.ISBN10, bookData.NumberOfPages, bookData.Covers, bookData.Series)
//
//	conn := connectToDB()
//
//	//migrates the DB tables!
//	migrateDB(conn)
//	seedGenres(conn)
//}

func main() {

}
