package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
)

// BookInfo a type which helps to pick out and translate the info returned from the API
//struct means that this is a type with multiple types within it (ints/strings/arrays/etc.)
//types can be anything - but if you say the type is an int, only numbers will go into them.
type BookInfo struct {
	Publishers    []string    `json:"publishers"`
	NumberOfPages int         `json:"number_of_pages"`
	ISBN10        []string    `json:"isbn_10"`
	ISBN13        []string    `json:"isbn_13"`
	Series        []string    `json:"series"`
	Covers        []int       `json:"covers"`
	Key           string      `json:"key"`
	Authors       []APIAuthor `json:"authors"`
	Contributions []string    `json:"contributions"`
	Title         string      `json:"title"`
	PublishDate   string      `json:"publishDate"`
}

type Book struct {
	gorm.Model
	ISBN       string `gorm:"unique"`
	Key        string
	Title      string   `gorm:"index;<-:create"`
	Authors    []Author `gorm:"many2many:book_author;"`
	Series     string
	GenreID    uint
	PageNumber int
}

//gets the book information from the API
func getBookInfo(barcode string) []byte {
	//generating the url depending on the barcode
	url := []string{"https://openlibrary.org/isbn/", barcode, ".json"}

	//get this url and output it to either a response or an error
	//if it is an error, print an error text and exit
	response, err := http.Get(strings.Join(url, ""))
	exitErrorHandler(err)

	//read the response that we get from the api, if can't read run fatalError
	//if can read, return the responseData
	//returns the information in a byte array
	responseData, err := ioutil.ReadAll(response.Body)
	fatalErrorHandler(err)
	return responseData
}

//helps to make the book information more readable
func parseBookJson(bookInfo []byte) BookInfo {
	//declare the variable of data and when it is unmarshalled it goes into this variable
	var data BookInfo

	//converts the byte array into json
	if err := json.Unmarshal(bookInfo, &data); err != nil {
		panic(err)
	}
	return data
}

//gets the cover pic from the API and returns it as a working URL
func coverPicURL(barcode string) string {
	//creates a url
	url := []string{"https://covers.openlibrary.org/b/isbn/", barcode, "-M.jpg"}

	//creates a joined url that can be called upon later
	//TODO have this send to the front-end for the cover image
	return strings.Join(url, "")
}

//addABook adds a book if it is not in the Book db table, if it does exist, it returns
func addABook(conn *gorm.DB, bookData BookInfo, returnedAuthors []APIAuthor, barcode string, genre uint) {
	book := Book{}
	book.Title = bookData.Title
	bookTableData := book.Retrieve(conn)
	if bookTableData.RowsAffected != 0 {
		return
	} else {
		book.ISBN = barcode
		book.PageNumber = bookData.NumberOfPages
		book.GenreID = genre
		//for every author within the array authors, add to the empty author array in bookAuthors
		for _, author := range returnedAuthors {
			book.Authors = append(book.Authors, Author{Name: author.Name, Key: author.Key})
		}
		book.Create(conn)
	}

}

// Create a book
//sends the data through to gorm to create the row within the db table
func (b *Book) Create(conn *gorm.DB) *gorm.DB {
	//adds the variable information to the data using pass by reference
	//checks to make sure things are added, if nothing prints nothing is added or a error response
	return conn.Create(&b)

}

// Retrieve checks book is in db table and gets it
func (b *Book) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&b).Find(&b)
}
