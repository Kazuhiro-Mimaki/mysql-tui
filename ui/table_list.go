package ui

import (
	"tui-dbms/database/mysql"

	"github.com/rivo/tview"
)

type TableList struct {
	List *tview.List
}

// Initialize table list
func NewTableList(database mysql.IDatabaase) *TableList {
	list := tview.NewList()

	list.ShowSecondaryText(false)
	list.SetTitle("Tables")
	list.SetBorder(true)

	tbl := &TableList{
		List: list,
	}

	return tbl
}

func (tbl *TableList) SetTableList(dbName string, database mysql.IDatabaase) {
	tbl.List.Clear()

	tables := database.ShowTables(dbName)

	for _, tableName := range tables {
		tbl.List.AddItem(tableName, tableName, 0, nil)
	}
}
