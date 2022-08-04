package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type TableColumn struct {
	columnNames []string
	records     [][]*string
}

func Driver() {
	db, err := sql.Open("mysql", "root@(localhost:3306)/world") // "dbUser:dbPassword@(dbURL:PORT)/dbName"
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	query := "SELECT * FROM country"

	queryResult, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	columns, err := queryResult.Columns()
	if err != nil {
		log.Println(err)
		return
	}

	var records = [][]*string{}
	for queryResult.Next() {
		rows := make([]*string, len(columns))
		rowPointers := make([]interface{}, len(columns))
		for i := range rows {
			rowPointers[i] = &rows[i]
		}

		err = queryResult.Scan(rowPointers...)
		if err != nil {
			return
		}

		records = append(records, rows)
	}

	tableColumn := TableColumn{
		columnNames: columns,
		records:     records,
	}

	log.Println(tableColumn)
}
