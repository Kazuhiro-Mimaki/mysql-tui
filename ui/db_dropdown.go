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
	var firstOptionIndex = 0

	databases := database.ShowDatabases()

	dropdown := tview.NewDropDown()

	dropdown.SetOptions(databases, nil)
	dropdown.SetTitle("Database")
	dropdown.SetBorder(true)
	dropdown.SetFieldTextColor(tcell.ColorBlack)
	dropdown.SetCurrentOption(firstOptionIndex)

	return &DatabaseDropDown{
		DropDown: dropdown,
	}
}
