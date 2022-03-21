package controllers

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "login", BasicPageData{Title: "Tattoo Trees"})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "register", BasicPageData{Title: "Tattoo Trees"})
}
