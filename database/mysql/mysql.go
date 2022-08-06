package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	pool *sql.DB
}

type IDatabaase interface {
	ShowTables() []string
	GetRecords(table string) [][]*string
}

// Initialize mysql
func NewMySQL() *MySQL {
	pool, err := sql.Open("mysql", "root@(localhost:3306)/world") // "dbUser:dbPassword@(dbURL:PORT)/dbName"
	if err != nil {
		log.Println(err)
	}

	return &MySQL{pool: pool}
}

func (mysql *MySQL) ShowTables() []string {
	row, err := mysql.pool.Query("SHOW TABLES")
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

func (mysql *MySQL) GetRecords(table string) [][]*string {
	row, err := mysql.pool.Query(fmt.Sprintf("SELECT * FROM %s", table))
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

	var records = [][]*string{}
	// set column names at first
	records = append(records, colsNames)
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

		records = append(records, fields)
	}

	return records
}
