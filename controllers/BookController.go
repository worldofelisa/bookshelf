package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tattooedtrees/customerrors"
	"tattooedtrees/model"
)

type FormBook struct {
	ISBN string `json:"isbn"`
}

//BookHandler allows you to display the page from the template which has a form to add a book
func BookHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "addbook", BasicPageData{Title: "Tattoo Trees"})
}

func PostBookHandler(w http.ResponseWriter, r *http.Request) {
	result, err := ioutil.ReadAll(r.Body)
	customerrors.PrintErrorHandler(err)
	var formData FormBook
	if err = json.Unmarshal(result, &formData); err != nil {
		customerrors.PrintErrorHandler(err)
	}
	barcoding(formData.ISBN)

}

func barcoding(barcode string) {
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
	//fmt.Println(bookData.Title, bookData.ISBN10, bookData.NumberOfPages, bookData.Covers, bookData.Series)
	book := model.Book{Key: bookData.Key, Title: bookData.Title, ISBN: barcode, PageNumber: bookData.NumberOfPages, Authors: authors}
	model.Create(conn, &book)
}

//create the book page (ensure that it adds mutliple db worth of data, i.e. book info + reviews + tags + genres info)

//view the book once created

//update the book with review and tag

//deleting a book off of your shelf
