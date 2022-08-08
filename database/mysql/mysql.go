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
	ShowDatabases() []string
	ShowTables(database string) []string
	GetRecords(table string) [][]*string
	changeTable(database string)
}

/*
====================
Initialize mysql
====================
*/
func NewMySQL(database string) *MySQL {
	pool, err := sql.Open("mysql", fmt.Sprintf("root@(localhost:3306)/%s", database)) // "dbUser:dbPassword@(dbURL:PORT)/dbName"
	if err != nil {
		log.Println(err)
	}

	return &MySQL{pool: pool}
}

/*
====================
Show database list
====================
*/
func (mysql *MySQL) ShowDatabases() []string {
	row, err := mysql.pool.Query("SHOW DATABASES")
	if err != nil {
		log.Println(err)
	}

	databases := []string{}
	for row.Next() {
		var databaseName string
		err = row.Scan(&databaseName)
		if err != nil {
			log.Println(err)
		}
		databases = append(databases, databaseName)
	}

	return databases
}

/*
====================
Create connection with other table
====================
*/
func (mysql *MySQL) changeTable(database string) {
	pool, err := sql.Open("mysql", fmt.Sprintf("root@(localhost:3306)/%s", database))
	if err != nil {
		log.Println(err)
	}

	mysql.pool = pool
}

/*
====================
Show table list
====================
*/
func (mysql *MySQL) ShowTables(database string) []string {
	mysql.changeTable(database)

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

/*
====================
Get records
====================
*/
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
