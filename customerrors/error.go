package customerrors

import (
	"fmt"
	"log"
	"os"
)

func FatalErrorHandler(err error) {
	//if there is an error with opening/finding the image, it will output a fatal error
	if err != nil {
		log.Fatalln(err)
	}
}

func ExitErrorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func PrintErrorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
