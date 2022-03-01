package main

import (
	"encoding/json"
	"fmt"
	"github.com/sbinet/isbn"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type BookInfo struct {
	Publishers    []string `json:"publishers"`
	NumberOfPages int      `json:"number_of_pages"`
	ISBN10        []string `json:"isbn_10"`
	ISBN13        []string `json:"isbn_13"`
	Covers        []int    `json:"covers"`
	Key           string   `json:"key"`
	Authors       []Author `json:"authors"`
	Contributions []string `json:"contributions"`
	Title         string   `json:"title"`
	PublishDate   string   `json:"publishDate"`
}

type Author struct {
	Key string `json:"key"`
}

func fatalErrorHandler(err error) {
	//if there is an error with opening/finding the image, it will output a fatal error
	if err != nil {
		log.Fatal(err)
	}
}

func scanImage(img string) string {
	//takes two variables and says open this image (specific via relative path)
	reader, err := os.Open(img)
	fatalErrorHandler(err)

	//defer says at the end of this function, run this command
	//closes the reader
	defer reader.Close()

	// decode returns an image, the format of the image as a string, and an error
	// we don't need the string so we throw it away/don't use it.
	// m = decoded image
	m, _, err := image.Decode(reader)
	fatalErrorHandler(err)

	//takes our image and scans it from the barcode reader from github --see import
	//Scan returns a barcode and error, barcode is just an int array
	barcode, err := isbn.Scan(m)
	fatalErrorHandler(err)

	//takes the int array of barcode and converts it into a string without spaces between each number
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(barcode), " "), ""), "[]")
}

func getBookInfo(barcode string) []byte {
	//generating the url depending on the barcode
	url := []string{"https://openlibrary.org/isbn/", barcode, ".json"}

	//get this url and output it to either a response or an error
	//if it is an error, print an error text and exit
	response, err := http.Get(strings.Join(url, ""))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

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

	//num := data["num"].(float64)
	//fmt.Println(num)
}

func main() {
	barcode := scanImage("./IMG_4779.jpg")
	bookInfo := getBookInfo(barcode)
	bookData := parseBookJson(bookInfo)
	fmt.Println(bookData.Authors, bookData.Title, bookData.ISBN10, bookData.NumberOfPages)
}
