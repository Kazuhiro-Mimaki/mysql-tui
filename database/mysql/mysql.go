package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rivo/tview"
)

func Driver() [][]*string {
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
		}

		tableData = append(tableData, fields)
	}

	return tableData
}

func ShowData(viewTable *tview.Table, tableData [][]*string) {
	for i, row := range tableData {
		for j, col := range row {
			var cellValue string

			if col != nil {
				cellValue = *col
			}

			viewTable.SetCell(
				i, j,
				tview.NewTableCell(cellValue),
			)
		}
	}
	viewTable.SetSelectable(true, true)
}
