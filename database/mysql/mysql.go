package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetTables() []string {
	db, err := sql.Open("mysql", "root@(localhost:3306)/world") // "dbUser:dbPassword@(dbURL:PORT)/dbName"
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	query := "SHOW TABLES"

	row, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	tables := []string{}
	for row.Next() {
		var tableName string
		err = row.Scan(&tableName)
		if err != nil {
			log.Println(err)
		}
		tables = append(tables, tableName)
	}

	return tables
}

func GetRecords() [][]*string {
	db, err := sql.Open("mysql", "root@(localhost:3306)/world") // "dbUser:dbPassword@(dbURL:PORT)/dbName"
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	query := "SELECT * FROM country"

	row, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	cols, err := row.Columns()
	if err != nil {
		log.Println(err)
	}

	var colsNames []*string
	for _, col := range cols {
		colName := col
		colsNames = append(colsNames, &colName)
	}

	var tableData = [][]*string{}
	tableData = append(tableData, colsNames)
	for row.Next() {
		fields := make([]*string, len(cols))
		fieldsPointers := make([]interface{}, len(cols))
		for i := range fields {
			fieldsPointers[i] = &fields[i]
		}

		// scan by a row, and set to pointers
		err = row.Scan(fieldsPointers...)
		if err != nil {
			log.Println(err)
		}

		tableData = append(tableData, fields)
	}

	return tableData
}
