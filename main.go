package main

import (
	_ "image/jpeg"
	"log"
	"net/http"
	"tattooedtrees/controllers"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", controllers.Handler()))
}
