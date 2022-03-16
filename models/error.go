package main

import (
	"fmt"
	"log"
	"os"
)

func fatalErrorHandler(err error) {
	//if there is an error with opening/finding the image, it will output a fatal error
	if err != nil {
		log.Fatal(err)
	}
}

func exitErrorHandler(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func printErrorHandler(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
