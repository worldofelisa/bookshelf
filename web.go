package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"tattooedtrees/models"
)

//Handler registers the route mapping to the URL
func Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", LandingHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/home", HomeHandler)
	return r
}

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "index", BasicPageData{Title: "Tattoo Trees"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "login", BasicPageData{Title: "Tattoo Trees"})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "register", BasicPageData{Title: "Tattoo Trees"})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "home", HomePageData{
		BasicPageData: BasicPageData{Title: "Tattoo Trees"}, ShelfTitles: []string{
			"Read",
			"Currently Reading",
			"To Read",
			"5 Stars",
			"4+ Stars",
			"3+ Stars",
			"DNF",
		}})
	book := model.Book{}

}

var templates = template.Must(template.ParseGlob("templates/*.gohtml"))

func DisplayPage(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".gohtml", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
