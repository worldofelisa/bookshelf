package customerrors

import (
	"fmt"
	"log"
	"os"
)

func FatalErrorHandler(err error) {
	//if there is an error with opening/finding the image, it will output a fatal error
	if err != nil {
		log.Fatal(err)
	}
}

func ExitErrorHandler(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func PrintErrorHandler(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
