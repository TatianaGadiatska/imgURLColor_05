package app

import (
	"c/GoExam/imgUrlColor_05/internal/parser"
	"c/GoExam/imgUrlColor_05/internal/store"
	m "c/GoExam/imgUrlColor_05/model"
	"c/GoExam/imgUrlColor_05/repo"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//URLColors ...
var URLColors []m.URLImage = parser.GetImagesLinks()

var db *sql.DB = store.CreateDB()

//RunBD ...
func RunBD() {

	store.CreateTable(db)

	repo.InsertURL(db, URLColors)

	//defer db.Close()
}

//RunAPI ...
func RunAPI() {

	router := mux.NewRouter()
	log.Print("Router starting...")

	router.HandleFunc("/", handleURLImages(db))

	err := http.ListenAndServe(":8181", router)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("starting api server ...")
	}
}
