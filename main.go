package main

import (
	"log"
	"tui-dbms/database/mysql"
	"tui-dbms/ui"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App *tview.Application

	// mysql connection
	mysql *mysql.MySQL

	// view components
	DBDropDownComponent *ui.DBDropDownComponent
	TableListComponent  *ui.TableListComponent
	TableGridComponent  *ui.TableGridComponent
	Flex                *ui.FlexLayout
}

/*
====================
main
====================
*/
func main() {
	tui := NewTui()

	tui.setEventKey()

	if err := tui.App.SetRoot(tui.Flex.Main, true).EnableMouse(false).Run(); err != nil {
		log.Println(err)
	}
}

/*
====================
Initialize tui
====================
*/
func NewTui() *TUI {
	tui := &TUI{
		App:   tview.NewApplication(),
		mysql: mysql.NewMySQL(""),
	}

	databases := tui.mysql.ShowDatabases()

	tui.DBDropDownComponent = ui.NewDBDropDownComponent(databases)
	tui.TableListComponent = ui.NewTableListComponent()
	tui.TableGridComponent = ui.NewTableGridComponent()

	tui.Flex = ui.NewMainFlex(
		tui.DBDropDownComponent.View,
		tui.TableListComponent.View,
		tview.NewBox().SetBorder(true).SetTitle("Query Input"),
		tui.TableGridComponent.View,
	)

	// when the databases selected, update the table list
	tui.DBDropDownComponent.View.SetSelectedFunc(func(selectedDatabase string, _ int) {
		tui.selectDatabase(selectedDatabase)
		tui.setFocus(tui.TableListComponent.View)
	})

	// when the table selected, update the table records
	tui.TableListComponent.View.SetSelectedFunc(func(_ int, selectedTable, _ string, _ rune) {
		tui.selectTable(selectedTable)
		tui.setFocus(tui.TableGridComponent.View)
	})

	tui.highlightFocusedArea(tui.App.GetFocus())

	return tui
}

/*
====================
Execute by goroutine
====================
*/
func (tui *TUI) queueUpdateDraw(f func()) {
	go func() {
		tui.App.QueueUpdateDraw(f)
	}()
}

/*
====================
Set focus on the target area
====================
*/
func (tui *TUI) setFocus(p tview.Primitive) {
	tui.queueUpdateDraw(func() {
		tui.App.SetFocus(p)
	})
}

/*
====================
Select database
====================
*/
func (tui *TUI) selectDatabase(selectedDB string) {
	tui.queueUpdateDraw(func() {
		tables := tui.mysql.ShowTables(selectedDB)
		tui.TableListComponent.SetTableListView(tables)
	})
}

/*
====================
Select table
====================
*/
func (tui *TUI) selectTable(selectedTable string) {
	tui.queueUpdateDraw(func() {
		records := tui.mysql.GetRecords(selectedTable)
		tui.TableGridComponent.SetTableView(records)
	})
}

/*
====================
Highlight the focused area
====================
*/
func (tui *TUI) highlightFocusedArea(focusedArea tview.Primitive) {
	tui.queueUpdateDraw(func() {
		var highlightColor = tcell.ColorGreen

		switch focusedArea {
		case tui.DBDropDownComponent.View:
			tui.DBDropDownComponent.View.SetBorderColor(highlightColor)
		case tui.TableListComponent.View:
			tui.TableListComponent.View.SetBorderColor(highlightColor)
		case tui.TableGridComponent.View:
			tui.TableGridComponent.View.SetBorderColor(highlightColor)
		}
	})
}

/*
====================
Set event key config
====================
*/
func (tui *TUI) setEventKey() {
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			tui.setFocus(tui.DBDropDownComponent.View)
		case tcell.KeyCtrlS:
			tui.setFocus(tui.TableListComponent.View)
		case tcell.KeyCtrlE:
			tui.setFocus(tui.TableGridComponent.View)
		}
		return event
	})
}
