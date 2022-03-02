package main

import (
	"fmt"
	"github.com/sbinet/isbn"
	"image"
	"os"
	"strings"
)

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
