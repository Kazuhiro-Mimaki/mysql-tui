package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TableSchemas struct {
	View *tview.Table
	Data [][]*string
}

/*
====================
Initialize table schema
====================
*/
func NewTableSchemas() *TableSchemas {
	var tableView = tview.NewTable().SetSelectable(true, true)

	return &TableSchemas{
		View: tableView,
	}
}

/*
====================
Set new table data
====================
*/
func (tbs *TableSchemas) SetData(data [][]*string) {
	tbs.View.Clear().ScrollToBeginning()

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

			tbs.View.SetCell(
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
