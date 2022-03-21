package controllers

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"tattooedtrees/customerrors"
)

var templates = template.Must(template.ParseGlob("templates/*.gohtml"))

type HomePageData struct {
	BasicPageData
	ShelfTitles []string
}

type BasicPageData struct {
	Title string
}

//Handler registers the route mapping to the URL
func Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", LandingHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/home", HomeHandler)
	return r
}

func DisplayPage(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".gohtml", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		customerrors.PrintErrorHandler(err)
		return
	}
}
