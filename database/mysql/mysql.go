package mysql

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Country struct {
	Id        int
	Iso, Name string
}

func Driver() {
	db, err := sqlx.Connect("mysql", "root@(localhost:3306)/sample_country") // "dbUser:dbPassword@(dbURL:PORT)/dbName"
	if err != nil {
		log.Println(err)
		return
	}

	countries := []Country{}

	db.Select(&countries, "SELECT id, iso, name FROM country")

	log.Println(countries)
}
