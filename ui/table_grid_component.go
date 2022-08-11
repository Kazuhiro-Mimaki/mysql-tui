package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type currentMode uint8

const (
	Record currentMode = iota + 1
	Schema
)

type TableGridComponent struct {
	CurrentMode currentMode
	View        *tview.Table
	Data        TableData
}

type TableData struct {
	Schemas [][]*string
	Records [][]*string
}

/*
====================
Initialize table grid
====================
*/
func NewTableGridComponent() *TableGridComponent {
	var defaultMode = Record

	tableView := tview.NewTable()

	tableView.SetSelectable(true, true)
	tableView.SetTitle("Table view")
	tableView.SetTitleAlign(tview.AlignLeft)
	tableView.SetBorder(true)
	// fix column names
	tableView.SetFixed(1, 0)

	tc := &TableGridComponent{
		CurrentMode: defaultMode,
		View:        tableView,
	}

	tc.setEventKey()

	return tc
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

/*
====================
Switch table view mode
====================
*/
func (tc *TableGridComponent) switchMode() {
	switch tc.CurrentMode {
	case Record:
		tc.CurrentMode = Schema
		tc.SetTableView(tc.Data.Schemas)
	case Schema:
		tc.CurrentMode = Record
		tc.SetTableView(tc.Data.Records)
	}
}

/*
====================
Set event key config
====================
*/
func (tc *TableGridComponent) setEventKey() {
	tc.View.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlP:
			tc.switchMode()
		}
		return event
	})
}
