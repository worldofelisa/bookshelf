package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

/*things I have tried:
-alt+enter [surprise, did nothing]
-adding a listen and serve (as it is saying serve http in the error
-adding a function to the homehandler passing arguments
-making a makehandler interface and using that to pass home handler in
where I have looked:
-github gorilla mux
-go docs (HTTP PACKAGE)
-https://www.calhoun.io/why-cant-i-pass-this-function-as-an-http-handler/
-https://medium.com/geekculture/demystifying-http-handlers-in-golang-a363e4222756
*/

//Handler registers the route mapping to the URL
func Handler() {
	r := mux.NewRouter()
	r.Handle("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	DisplayPage(w, "index", BasicPageData{Title: "Tattoo Trees"})
	log.Fatal(http.ListenAndServe(":8080", nil))

	//vars := mux.Vars(r)
	//w.WriteHeader(http.StatusOK)
}

var templates = template.Must(template.ParseGlob("templates/*.gohtml"))

func DisplayPage(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".gohtml", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
