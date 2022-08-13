package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TableRecords struct {
	View *tview.Table
	Data [][]*string
}

/*
====================
Initialize table schema
====================
*/
func NewTableRecords() *TableRecords {
	var tableView = tview.NewTable().SetSelectable(true, true)

	return &TableRecords{
		View: tableView,
	}
}

/*
====================
Set new table data
====================
*/
func (tbr *TableRecords) SetData(data [][]*string) {
	tbr.View.Clear().ScrollToBeginning()

	for i, row := range data {
		for j, col := range row {
			var cellValue string
			var cellColor = tcell.ColorWhite
			var notSelectable = false

			if col != nil {
				cellValue = *col
			}

			// top(column names)
			if i == 0 {
				cellColor = tcell.ColorNavy
			}

			tbr.View.SetCell(
				i, j,
				&tview.TableCell{
					Text:          cellValue,
					Color:         cellColor,
					NotSelectable: notSelectable,
				},
			)
		}
	}
}
