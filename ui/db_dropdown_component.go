package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type DBDropDownComponent struct {
	View *tview.DropDown
}

/*
====================
Initialize db dropdown
====================
*/
func NewDBDropDownComponent(databases []string) *DBDropDownComponent {
	var firstOptionIndex = 0

	dropdownView := tview.NewDropDown()

	dropdownView.SetOptions(databases, nil)
	dropdownView.SetTitle("Database")
	dropdownView.SetTitleAlign(tview.AlignLeft)
	dropdownView.SetBorder(true)
	dropdownView.SetFieldTextColor(tcell.ColorBlack)
	dropdownView.SetCurrentOption(firstOptionIndex)

	return &DBDropDownComponent{
		View: dropdownView,
	}
}
