package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"tattooedtrees/customerrors"
	"tattooedtrees/model"
)

type FormBook struct {
	ISBN string `json:"isbn"`
}

type ConfirmBook struct {
	BookResponse model.Book
	Genres       []string
}

type SubmitBook struct {
	Title  string
	Author string
	Pages  string
	Genre  string
	Tags   string
	Review string
	ISBN   string
	Key    string
}

//BookHandler allows you to display the page from the template which has a form to add a book
func BookHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "addbook", BasicPageData{Title: "Tattoo Trees"})
}

func PostBookHandler(w http.ResponseWriter, r *http.Request) {
	result, err := ioutil.ReadAll(r.Body)
	customerrors.PrintErrorHandler(err)
	var formData FormBook
	err = json.Unmarshal(result, &formData)
	customerrors.PrintErrorHandler(err)
	book := barcoding(formData.ISBN)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ConfirmBook{
		book,
		model.Genres,
	})
	customerrors.PrintErrorHandler(err)
}

func barcoding(barcode string) model.Book {
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
	book := model.Book{Key: bookData.Key, Title: bookData.Title, ISBN: barcode, PageNumber: bookData.NumberOfPages, Authors: authors}
	return book
}

//create the book page (ensure that it adds mutliple db worth of data, i.e. book info + reviews + tags + genres info)
func PostSubmitBookHandler(w http.ResponseWriter, r *http.Request) {
	result, err := ioutil.ReadAll(r.Body)
	customerrors.PrintErrorHandler(err)
	fmt.Println(result)
	var submitData SubmitBook
	err = json.Unmarshal(result, &submitData)
	customerrors.PrintErrorHandler(err)
	page, err := strconv.Atoi(submitData.Pages)
	customerrors.PrintErrorHandler(err)
	genre, err := strconv.ParseUint(submitData.Genre, 10, 64)
	customerrors.PrintErrorHandler(err)
	book := model.Book{
		ISBN:       submitData.ISBN,
		Key:        submitData.Key,
		Title:      submitData.Title,
		Authors:    nil,
		Series:     "",
		GenreID:    uint(genre),
		PageNumber: page,
	}
	//TODO add tags into db too
	//model.Create(conn, &book)
	fmt.Println(submitData)
}

//view the book once created

//update the book with review and tag

//deleting a book off of your shelf
