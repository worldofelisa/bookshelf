package controllers

import "net/http"

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "index", BasicPageData{Title: "Tattoo Trees"})
}
