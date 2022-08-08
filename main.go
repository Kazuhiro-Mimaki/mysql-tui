package main

import (
	"log"
	"tui-dbms/database/mysql"
	"tui-dbms/ui"
	"tui-dbms/ui/flex"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App *tview.Application

	// mysql connection
	mysql *mysql.MySQL

	// view components
	DatabaseDropDown *ui.DatabaseDropDown
	TableList        *ui.TableList
	TableGrid        *ui.TableGrid
}

func main() {
	tui := NewTui()

	flex := flex.NewFlex(tui.DatabaseDropDown.DropDown, tui.TableList.List, tview.NewBox().SetBorder(true).SetTitle("Query Input"), tui.TableGrid.Records)

	tui.setEventKey()

	if err := tui.App.SetRoot(flex, true).EnableMouse(false).Run(); err != nil {
		log.Println(err)
	}
}

func NewTui() *TUI {
	tui := &TUI{
		App:   tview.NewApplication(),
		mysql: mysql.NewMySQL(""),
	}

	databases := tui.mysql.ShowDatabases()

	tui.DatabaseDropDown = ui.NewDatabaseDropDown(databases)
	tui.TableList = ui.NewTableList()
	tui.TableGrid = ui.NewTableGrid()

	// when the databases selected, update the table list
	tui.DatabaseDropDown.DropDown.SetSelectedFunc(func(text string, index int) {
		tui.selectDatabase(text)
		tui.setFocus(tui.TableList.List)
	})

	// when the table selected, update the table records
	tui.TableList.List.SetSelectedFunc(func(int, string, string, rune) {
		selectedTable, _ := tui.TableList.List.GetItemText(tui.TableList.List.GetCurrentItem())
		tui.updateTable(selectedTable)
		tui.setFocus(tui.TableGrid.Records)
	})

	tui.highlightFocusedArea(tui.App.GetFocus())

	return tui
}

func (tui *TUI) queueUpdateDraw(f func()) {
	go func() {
		tui.App.QueueUpdateDraw(f)
	}()
}

func (tui *TUI) setFocus(p tview.Primitive) {
	tui.queueUpdateDraw(func() {
		tui.App.SetFocus(p)
	})
}

func (tui *TUI) selectDatabase(selectedDB string) {
	tui.queueUpdateDraw(func() {
		tui.TableList.SetTableList(selectedDB, tui.mysql)
	})
}

func (tui *TUI) updateTable(selectedTable string) {
	tui.queueUpdateDraw(func() {
		tui.TableGrid.SetTableGrid(selectedTable, tui.mysql)
	})
}

func (tui *TUI) highlightFocusedArea(focusedArea tview.Primitive) {
	tui.queueUpdateDraw(func() {
		var highlightColor = tcell.ColorGreen

		switch focusedArea {
		case tui.DatabaseDropDown.DropDown:
			tui.DatabaseDropDown.DropDown.SetBorderColor(highlightColor)
		case tui.TableList.List:
			tui.TableList.List.SetBorderColor(highlightColor)
		case tui.TableGrid.Records:
			tui.TableGrid.Records.SetBorderColor(highlightColor)
		}
	})
}

func (tui *TUI) setEventKey() {
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			tui.setFocus(tui.DatabaseDropDown.DropDown)
		case tcell.KeyCtrlS:
			tui.setFocus(tui.TableList.List)
		case tcell.KeyCtrlE:
			tui.setFocus(tui.TableGrid.Records)
		}
		return event
	})
}
