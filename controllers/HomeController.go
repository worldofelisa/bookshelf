package controllers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "home", HomePageData{
		BasicPageData{Title: "Tattoo Trees"},
		[]string{
			"Read",
			"Currently Reading",
			"To Read",
			"5 Stars",
			"4+ Stars",
			"3+ Stars",
			"DNF",
		}})
}
