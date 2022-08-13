package ui

import (
	table "tui-dbms/ui/table"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TableGridComponent struct {
	View    *tview.Pages
	Schemas *table.TableSchemas
	Records *table.TableRecords
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
	var tableSchemas = table.NewTableSchemas()
	var tableRecords = table.NewTableRecords()

	pageView.SetBorder(true)
	pageView.AddPage("Schemas", tableSchemas.View, true, false)
	pageView.AddPage("Records", tableRecords.View, true, true)

	var tbg = &TableGridComponent{
		View:    pageView,
		Schemas: tableSchemas,
		Records: tableRecords,
	}

	tbg.setEventKey()

	return tbg
}

/*
====================
Set new table view
====================
*/
func (tbg *TableGridComponent) SetTable(data TableData) {
	tbg.Schemas.SetData(data.Schemas)
	tbg.Records.SetData(data.Records)
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
