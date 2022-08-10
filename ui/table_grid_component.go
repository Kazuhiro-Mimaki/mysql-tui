package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TableGridComponent struct {
	View *tview.Table
}

/*
====================
Initialize table grid
====================
*/
func NewTableGridComponent() *TableGridComponent {
	tableView := tview.NewTable()

	tableView.SetSelectable(true, true)
	tableView.SetTitle("Records")
	tableView.SetBorder(true)

	// fix column names
	tableView.SetFixed(1, 0)

	return &TableGridComponent{
		View: tableView,
	}
}

/*
====================
Set new table view
====================
*/
func (tc *TableGridComponent) SetTableView(data [][]*string) {
	tc.ResetTableView()

	// Set records as default
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

			tc.View.SetCell(
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

/*
====================
Clear records and scroll to beginning
====================
*/
func (tc *TableGridComponent) ResetTableView() {
	tc.View.Clear().ScrollToBeginning()
}
