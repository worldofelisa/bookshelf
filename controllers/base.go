package controllers

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"tattooedtrees/customerrors"
	"tattooedtrees/database"
	"tattooedtrees/model"
)

var conn *gorm.DB

var user model.User

var templates = template.Must(template.ParseGlob("templates/*/*.gohtml"))

type HomePageData struct {
	BasicPageData
	ShelfTitles []string
}

type BasicPageData struct {
	Title string
}

//Handler registers the route mapping to the URL
func Handler() *mux.Router {
	initialising()
	r := mux.NewRouter()
	r.HandleFunc("/", LandingHandler)
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/addbook", BookHandler).Methods(http.MethodGet)
	r.HandleFunc("/addbook", PostBookHandler).Methods(http.MethodPost)
	r.HandleFunc("/addbook", PostSubmitBookHandler).Methods(http.MethodPost)
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js")))).Methods(http.MethodGet)
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css")))).Methods(http.MethodGet)
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

func initialising() {
	conn = database.ConnectToDB()
	//database.MigrateDB(conn)
	//model.SeedGenres(conn)
	user = model.User{Email: "elisa@elisa.com"}
	model.Retrieve(conn, &user)
}
