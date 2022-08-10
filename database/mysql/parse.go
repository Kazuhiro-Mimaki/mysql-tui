package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
====================
Show full columns
====================
*/
func scanRows(rows *sql.Rows) [][]*string {
	cols, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}

	var colNames []*string
	for _, col := range cols {
		colName := col
		colNames = append(colNames, &colName)
	}

	var fields = [][]*string{}
	// set column names at first
	fields = append(fields, colNames)
	for rows.Next() {
		row := make([]*string, len(cols))
		rowPointers := make([]interface{}, len(cols))
		for i := range row {
			rowPointers[i] = &row[i]
		}

		// scan by a row, and set to pointers
		err = rows.Scan(rowPointers...)
		if err != nil {
			log.Println(err)
		}

		fields = append(fields, row)
	}

	return fields
}
