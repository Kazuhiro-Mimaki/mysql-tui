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

	tui.setEvent()

	if err := tui.App.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		log.Println(err)
	}
}

func NewTui() *TUI {
	tui := &TUI{
		App:   tview.NewApplication(),
		mysql: mysql.NewMySQL(""),
	}

	tui.DatabaseDropDown = ui.NewDatabaseDropDown(tui.mysql)
	tui.TableList = ui.NewTableList(tui.mysql)
	tui.TableGrid = ui.NewTableGrid(tui.mysql)

	// when the databases selected, update the table list
	tui.DatabaseDropDown.DropDown.SetSelectedFunc(func(text string, index int) {
		tui.selectDatabase(text)
		tui.setFocus(tui.TableList.List)
	})

	// when the table selected, update the table records
	tui.TableList.List.SetSelectedFunc(func(int, string, string, rune) {
		selectedTable, _ := tui.TableList.List.GetItemText(tui.TableList.List.GetCurrentItem())
		tui.updateTable(selectedTable)
	})

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

func (tui *TUI) updateTable(newTable string) {
	tui.queueUpdateDraw(func() {
		tui.TableGrid.SetTableGrid(newTable, tui.mysql)
	})
}

func (tui *TUI) setEvent() {
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			tui.setFocus(tui.TableGrid.Records)
		}
		return event
	})
}
