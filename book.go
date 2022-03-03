package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func parseBookJson(bookInfo []byte) BookInfo {
	//declare the variable of data and when it is unmarshalled it goes into this variable
	var data BookInfo

	//converts the byte array into json
	if err := json.Unmarshal(bookInfo, &data); err != nil {
		panic(err)
	}
	return data
}

func coverPicURL(barcode string) string {
	//creates a url
	url := []string{"https://covers.openlibrary.org/b/isbn/", barcode, "-M.jpg"}

	//creates a joined url that can be called upon later
	//TODO have this send to the front-end for the cover image
	return strings.Join(url, "")
}
