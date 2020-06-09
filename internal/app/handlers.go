package app

import (
	"c/GoExam/imgUrlColor_05/model"
	"c/GoExam/imgUrlColor_05/repo"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func handleURLImages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		d := repo.GetImg(db)

		data := model.Data{
			URLImgData: d,
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Print(err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Print(err)
		} else {
			log.Print("Ok...")
		}
	}
}
