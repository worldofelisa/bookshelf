package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Author struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	PersonalName string `json:"personal_name"`
}

func getAuthorInfo(authInfo Author) []byte {
	//make the url depending on the author code
	url := []string{"https://openlibrary.org/", authInfo.Key, ".json"}

	//get this url and output it to a response or an error
	//if error, print error text and exit
	response, err := http.Get(strings.Join(url, ""))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//read the response we get from api, if can't read, run fatalError
	//if can read, return responseData
	//returns the information in a byte array
	responseData, err := ioutil.ReadAll(response.Body)
	fatalErrorHandler(err)
	return responseData
}

func parseAuthInfo(returnedAuthors []byte) Author {
	//declare the variable of data and when it is unmarshalled it goes into this variable
	var data Author

	//converts the byte array into json
	if err := json.Unmarshal(returnedAuthors, &data); err != nil {
		panic(err)
	}
	return data
}