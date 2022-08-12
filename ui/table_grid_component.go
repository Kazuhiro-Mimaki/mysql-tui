package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TableGridComponent struct {
	View    *tview.Pages
	Schemas *tview.Table
	Records *tview.Table
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
	var pageView = tview.NewPages()
	var schemaTable = tview.NewTable().SetSelectable(true, true)
	var recordTable = tview.NewTable().SetSelectable(true, true)

	pageView.SetBorder(true)
	pageView.AddPage("Schemas", schemaTable, true, false)
	pageView.AddPage("Records", recordTable, true, true)

	var tbg = &TableGridComponent{
		View:    pageView,
		Schemas: schemaTable,
		Records: recordTable,
	}

	tbg.setEventKey()

	return tbg
}

/*
====================
Set new table view
====================
*/
func (tbg *TableGridComponent) SetTableView(tableData TableData) {
	tbg.SetTableData(tbg.Schemas, tableData.Schemas)
	tbg.SetTableData(tbg.Records, tableData.Records)
}

/*
====================
Set new table data
====================
*/
func (tbg *TableGridComponent) SetTableData(targetGrid *tview.Table, data [][]*string) {
	targetGrid.Clear().ScrollToBeginning()

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

			targetGrid.SetCell(
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
Switch table view mode
====================
*/
func (tbg *TableGridComponent) switchMode() {
	var currentPage, _ = tbg.View.GetFrontPage()
	switch currentPage {
	case "Schemas":
		tbg.View.SwitchToPage("Records")
	case "Records":
		tbg.View.SwitchToPage("Schemas")
	}
}

/*
====================
Set event key config
====================
*/
func (tbg *TableGridComponent) setEventKey() {
	tbg.View.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlP:
			tbg.switchMode()
		}
		return event
	})
}
