package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/*
====================
Show full columns
====================
*/
func scanRows(rows *sql.Rows) (data [][]*string, err error) {
	cols, err := rows.Columns()
	if err != nil {
		return
	}

	var colNames []*string
	for _, col := range cols {
		colName := col
		colNames = append(colNames, &colName)
	}

	data = [][]*string{}
	// set column names at first
	data = append(data, colNames)
	for rows.Next() {
		row := make([]*string, len(cols))
		rowPointers := make([]interface{}, len(cols))
		for i := range row {
			rowPointers[i] = &row[i]
		}

		// scan by a row, and set to pointers
		err = rows.Scan(rowPointers...)
		if err != nil {
			return
		}

		data = append(data, row)
	}

	return
}
