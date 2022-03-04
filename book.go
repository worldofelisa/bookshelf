package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
)

// BookInfo a type which helps to pick out and translate the info returned from the API
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

//Creating a book
//Adds a book using information and inputs the info into the relative tables
//checks for errors and prints if nothing was added.
func addBook(conn *gorm.DB) {
	//takes information and assigns it to variable
	bookInfo := Book{
		Title:      "Mexican Gothic",
		PageNumber: 301,
		ISBN:       "9781529402681"}
	bookGenre := Genre{Name: "Gothic Fiction"}
	bookTags := []Tag{
		{Name: "Gothic"},
		{Name: "Fiction"},
		{Name: "Historical"},
		{Name: "Horror"},
		{Name: "Mystery"},
		{Name: "Fantasy"},
		{Name: "Thriller"},
		{Name: "Adult"}}
	bookAuthors := []Author{{Name: "Silvia Moreno-Garcia", Key: "/authors/OL7481396A"}}

	//adds the variable information to the data using pass by reference
	//checks to make sure things are added, if nothing prints nothing is added or a error response
	bookResult := conn.Create(&bookInfo)
	rowsAddedResponse(bookResult.RowsAffected)
	printErrorHandler(bookResult.Error)

	genreResult := conn.Create(&bookGenre)
	rowsAddedResponse(genreResult.RowsAffected)
	printErrorHandler(genreResult.Error)

	//loops over the array of authors and adds them to the table using pass by reference
	//checks to make sure things are added, if not same as above
	for _, author := range bookAuthors {
		authorResult := conn.Create(&author)
		rowsAddedResponse(authorResult.RowsAffected)
		printErrorHandler(authorResult.Error)
	}

	for _, author := range bookAuthors {
		bookAuthorResult := conn.Create(&BookAuthor{AuthorID: author.ID, BookID: bookInfo.ID})
		rowsAddedResponse(bookAuthorResult.RowsAffected)
		printErrorHandler(bookAuthorResult.Error)
	}

	for _, bookTag := range bookTags {
		tagsResult := conn.Create(&bookTag)
		rowsAddedResponse(tagsResult.RowsAffected)
		printErrorHandler(tagsResult.Error)
	}

	for _, bookTag := range bookTags {
		bookTagResult := conn.Create(&BookTag{BookID: bookInfo.ID, TagID: bookTag.ID})
		rowsAddedResponse(bookTagResult.RowsAffected)
		printErrorHandler(bookTagResult.Error)
	}
}

//Viewing/Reading a book

//Updating a book

//Deleting a book
