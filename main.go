package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Application struct {
	DB *sql.DB
}

//Database connection const
const (
	DBHost = "127.0.0.1"
	DBPort = ":3306"
	DBUser = "root"
	DBPass = ""
	DBName = "gozero"
)

//Database is a site wide database connection to Db
var database *sql.DB

func main() {
	//database initialization
	database, err := connection()
	if err != nil {
		fmt.Println("Response", err.Error())
		return
	}
	/////////////////////////////////////////////
	router(database)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

////////////////////////////////////////////////////////////
//Connecting to mysql DB
///////////////////////////////////////////////////////////

func connection() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

////////////////////////////////////////////////////////////////////////////////////
//using mux, db and struct
////////////////////////////////////////////////////////////////////////////////////

//Page struct matching db  schema
type Page struct {
	Title   string
	Content string
	Date    string
}

//var app *Application

func (app *Application) pageHandler(w http.ResponseWriter, r *http.Request) {

	guid, _ := mux.Vars(r)["guid"]

	thisPage := Page{}

	//query a row of db where id is 1
	err := app.DB.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?", guid).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)
	if err != nil {

		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	//Execute on web page
	//fmt.Fprint(w, "Response now is", thisPage.Content) or
	tpl, _ := template.New("errMsg").Parse("<h4> {{.Content}}</h4>")
	tpl.Execute(w, thisPage)

}

//Router
func router(db *sql.DB) {
	app := &Application{db}
	router := mux.NewRouter()
	router.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", app.pageHandler)
	http.Handle("/", router)
}
