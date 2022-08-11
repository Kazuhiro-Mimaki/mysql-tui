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
	GetSchemas(table string) [][]*string
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
	rows, err := mysql.pool.Query(fmt.Sprintf("SELECT * FROM %s", table))
	if err != nil {
		log.Println(err)
	}

	data, err := scanRows(rows)
	if err != nil {
		log.Println(err)
	}

	return data
}

/*
====================
Show full columns
====================
*/
func (mysql *MySQL) GetSchemas(table string) [][]*string {
	rows, err := mysql.pool.Query(fmt.Sprintf("SHOW FULL COLUMNS FROM %s", table))
	if err != nil {
		log.Println(err)
	}

	data, err := scanRows(rows)
	if err != nil {
		log.Println(err)
	}

	return data
}

/*
====================
Custom query
====================
*/
func (mysql *MySQL) CustomQuery(query string) (data [][]*string, err error) {
	rows, err := mysql.pool.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	data, err = scanRows(rows)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
