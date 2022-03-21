package main

import (
	"gorm.io/gorm"
	_ "image/jpeg"
	"log"
	"net/http"
	"tattooedtrees/controllers"
	"tattooedtrees/database"
	model "tattooedtrees/models"
)

var conn *gorm.DB

var user model.User

//func barcoding() {
//	barcode := "9781408890042"
//	bookInfo := model.GetBookInfo(barcode)
//	bookData := model.ParseBookJson(bookInfo)
//
//	authInfo := bookData.Authors
//	authors := []model.Author{}
//	for _, author := range authInfo {
//		//adding author information to a slice of new authors
//		returnedAuthors := model.ParseAuthInfo(model.GetAuthorInfo(author))
//		authors = append(authors, model.Author{
//			Key:  returnedAuthors.Key,
//			Name: returnedAuthors.Name,
//		})
//	}
//
//	model.CoverPicURL(barcode)
//	//fmt.Println(bookData.Title, bookData.ISBN10, bookData.NumberOfPages, bookData.Covers, bookData.Series)
//	book := model.Book{Key: bookData.Key, Title: bookData.Title, ISBN: barcode, PageNumber: bookData.NumberOfPages, Authors: authors}
//	model.Create(conn, &book)
//}

func initialising() {
	conn = database.ConnectToDB()
	database.MigrateDB(conn)
	model.SeedGenres(conn)
	user = model.User{Email: "elisa@elisa.com"}
	database.Retrieve(conn, &user)
}

func main() {
	initialising()
	//barcoding()
	log.Fatal(http.ListenAndServe(":8080", controllers.Handler()))
}
