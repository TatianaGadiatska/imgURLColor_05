package repo

import (
	m "c/GoExam/imgUrlColor_05/model"
	"database/sql"
	"log"
)

//InsertURL ...
func InsertURL(db *sql.DB, urlColorImages []m.URLImage) {

	for _, u := range urlColorImages {
		_, err := db.Exec("insert into img_url_color values(default, $1, $2)", u.URLImg, u.Color)

		if err != nil {
			log.Print(err)
		}
	}
	log.Print("Insert OK")
}

//GetImg ...
func GetImg(db *sql.DB) []m.URLImage {
	var imgURLs []m.URLImage
	var imgURL m.URLImage
	rows, err := db.Query("select * from img_url_color")
	if err != nil {
		log.Print(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&imgURL.ID, &imgURL.URLImg, &imgURL.Color)
		if err != nil {
			log.Print(err)
		}

		imgURLs = append(imgURLs, imgURL)
	}
	return imgURLs
}
