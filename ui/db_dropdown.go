package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type DatabaseDropDown struct {
	DropDown *tview.DropDown
}

/*
====================
Initialize db dropdown
====================
*/
func NewDatabaseDropDown(databases []string) *DatabaseDropDown {
	var firstOptionIndex = 0

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
