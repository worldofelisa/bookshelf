package main

import (
	"fmt"
	"gorm.io/gorm"
	_ "image/jpeg"
	model "tattooedtrees/models"
)

type HomePageData struct {
	BasicPageData BasicPageData
	ShelfTitles   []string
}

type BasicPageData struct {
	Title string
}

var conn *gorm.DB

func barcoding() {
	barcode := "9781408890042"
	bookInfo := model.GetBookInfo(barcode)
	bookData := model.ParseBookJson(bookInfo)

	authInfo := bookData.Authors
	authors := []model.Author{}
	for _, author := range authInfo {
		//adding author information to a slice of new authors
		returnedAuthors := model.ParseAuthInfo(model.GetAuthorInfo(author))
		authors = append(authors, model.Author{
			Key:  returnedAuthors.Key,
			Name: returnedAuthors.Name,
		})
	}

	model.CoverPicURL(barcode)
	fmt.Println(bookData.Title, bookData.ISBN10, bookData.NumberOfPages, bookData.Covers, bookData.Series)
	book := model.Book{Key: bookData.Key, Title: bookData.Title, ISBN: barcode, PageNumber: bookData.NumberOfPages, Authors: authors}
	model.Create(conn, &book)
}

func initialise() {
	model.MigrateDB(conn)
	model.SeedGenres(conn)
}

func main() {
	conn = model.ConnectToDB()
	initialise()
	barcoding()
	//log.Fatal(http.ListenAndServe(":8080", Handler()))
}
