package ui

import (
	"tui-dbms/database/mysql"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type DatabaseDropDown struct {
	DropDown *tview.DropDown
}

// Initialize db dropdown
func NewDatabaseDropDown(database mysql.IDatabaase) *DatabaseDropDown {
	databases := database.ShowDatabases()

	dropdown := tview.NewDropDown()

	dropdown.SetOptions(databases, nil)
	dropdown.SetTitle("Database")
	dropdown.SetBorder(true)
	dropdown.SetFieldTextColor(tcell.ColorBlack)

	return &DatabaseDropDown{
		DropDown: dropdown,
	}
}
