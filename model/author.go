package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
	"tattooedtrees/customerrors"
)

type APIAuthor struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	PersonalName string `json:"personal_name"`
}

type Author struct {
	gorm.Model
	Key   string
	Name  string  `gorm:"index;<-:create"`
	Books []*Book `gorm:"many2many:book_author;"`
}

func GetAuthorInfo(authInfo APIAuthor) []byte {
	//make the url depending on the author code
	url := []string{"https://openlibrary.org/", authInfo.Key, ".json"}

	//get this url and output it to a response or an error
	//if error, print error text and exit
	response, err := http.Get(strings.Join(url, ""))
	customerrors.ExitErrorHandler(err)

	//read the response we get from api, if can't read, run fatalError
	//if can read, return responseData
	//returns the information in a byte array
	responseData, err := ioutil.ReadAll(response.Body)
	customerrors.FatalErrorHandler(err)
	return responseData
}

func ParseAuthInfo(returnedAuthors []byte) APIAuthor {
	//declare the variable of data and when it is unmarshalled it goes into this variable
	var data APIAuthor

	//converts the byte array into json
	if err := json.Unmarshal(returnedAuthors, &data); err != nil {
		panic(err)
	}
	return data
}

// Create an author
//sends the data through to gorm to create the row within the db table
func (a *Author) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&a)
}

// Retrieve checks book is in db table and gets it
func (a *Author) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&a).Find(&a)
}

func (a *Author) Update(conn *gorm.DB) *gorm.DB {
	return conn.Save(&a)
}

func (a *Author) Delete(conn *gorm.DB) *gorm.DB {
	return conn.Delete(&a)
}
