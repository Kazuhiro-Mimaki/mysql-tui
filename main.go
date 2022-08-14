package main

import (
	"log"
	"time"
	"tui-dbms/database/mysql"
	"tui-dbms/ui"

	"tui-dbms/ui/tablemode/read"

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
	PageComponent       *ui.PageComponent
}

/*
====================
main
====================
*/
func main() {
	mysql := mysql.NewMySQL("")

	tui := NewTui(
		tview.NewApplication(),
		mysql,
		ui.NewDBDropDownComponent(mysql.ShowDatabases()),
		ui.NewTableListComponent(),
		ui.NewPageComponent(),
	)

	leftLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tui.DBDropDownComponent.View, 0, 1, true).
		AddItem(tui.TableListComponent.View, 0, 15, false)

	rightLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tui.PageComponent.View, 0, 1, false)

	mainLayout := tview.NewFlex().
		AddItem(leftLayout, 0, 1, true).
		AddItem(rightLayout, 0, 7, true)

	if err := tui.App.SetRoot(mainLayout, true).EnableMouse(false).Run(); err != nil {
		log.Println(err)
	}
}

/*
====================
Initialize tui
====================
*/
func NewTui(
	App *tview.Application,
	mysql *mysql.MySQL,
	DBDropDownComponent *ui.DBDropDownComponent,
	TableListComponent *ui.TableListComponent,
	PageComponent *ui.PageComponent,
) *TUI {
	tui := &TUI{
		App:                 App,
		mysql:               mysql,
		DBDropDownComponent: DBDropDownComponent,
		TableListComponent:  TableListComponent,
		PageComponent:       PageComponent,
	}

	tui.setEventKey()
	tui.setEventFunction()
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
		schemas, _ := tui.mysql.GetSchemas(selectedTable)
		records, _ := tui.mysql.GetRecords(selectedTable)

		var tableData = read.TableData{
			Schemas: schemas,
			Records: records,
		}

		tui.PageComponent.ReadLayout.TableGridComponent.SetTable(tableData)
	})
}

/*
====================
Execute read SQL
====================
*/
func (tui *TUI) executeReadQuery(query string) {
	tui.queueUpdateDraw(func() {
		records, err := tui.mysql.ReadQuery(query)
		if err != nil {
			tui.showReadSQLError(err)
		}

		var tableData = read.TableData{
			Schemas: nil,
			Records: records,
		}

		tui.PageComponent.ReadLayout.TableGridComponent.SetTable(tableData)
	})
}

/*
====================
Execute write SQL
====================
*/
func (tui *TUI) executeWriteQuery(query string) {
	tui.queueUpdateDraw(func() {
		successMsg, err := tui.mysql.WriteQuery(query)
		if err != nil {
			tui.showWriteSQLError(err)
		}

		tui.PageComponent.WriteLayout.SQLOutputFieldComponent.SetSuccessMessage(successMsg)
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
		case tui.PageComponent.ReadLayout.SQLInputFieldComponent.View:
			tui.PageComponent.ReadLayout.SQLInputFieldComponent.View.SetBorderColor(highlightColor)
		case tui.PageComponent.ReadLayout.TableGridComponent.View:
			tui.PageComponent.ReadLayout.TableGridComponent.View.SetBorderColor(highlightColor)
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
		case tcell.KeyCtrlI:
			tui.setFocus(tui.PageComponent.ReadLayout.SQLInputFieldComponent.View)
		case tcell.KeyCtrlE:
			tui.setFocus(tui.PageComponent.ReadLayout.TableGridComponent.View)
		case tcell.KeyCtrlO:
			page, _ := tui.PageComponent.View.GetFrontPage()
			if page == "Read" {
				tui.PageComponent.View.SwitchToPage("Write")
				tui.setFocus(tui.PageComponent.WriteLayout.SQLInputFieldComponent.View)
			} else {
				tui.PageComponent.View.SwitchToPage("Read")
				tui.setFocus(tui.PageComponent.ReadLayout.SQLInputFieldComponent.View)
			}
		}
		return event
	})
}

/*
====================
Set selected functions
====================
*/
func (tui *TUI) setEventFunction() {
	// DB dropdown
	tui.DBDropDownComponent.View.SetSelectedFunc(func(selectedDatabase string, _ int) {
		tui.selectDatabase(selectedDatabase)
		tui.setFocus(tui.TableListComponent.View)
	})

	// Table list
	tui.TableListComponent.View.SetSelectedFunc(func(_ int, selectedTable, _ string, _ rune) {
		tui.selectTable(selectedTable)
		tui.setFocus(tui.PageComponent.ReadLayout.TableGridComponent.View)
	})

	// SQL read input
	tui.PageComponent.ReadLayout.SQLInputFieldComponent.View.SetDoneFunc(func(key tcell.Key) {
		inputQuery := tui.PageComponent.ReadLayout.SQLInputFieldComponent.View.GetText()
		tui.executeReadQuery(inputQuery)
		tui.setFocus(tui.PageComponent.ReadLayout.TableGridComponent.View)
	})

	// SQL write input
	tui.PageComponent.WriteLayout.SQLInputFieldComponent.View.SetDoneFunc(func(key tcell.Key) {
		inputQuery := tui.PageComponent.WriteLayout.SQLInputFieldComponent.View.GetText()
		tui.executeWriteQuery(inputQuery)
	})
}

func (tui *TUI) showReadSQLError(err error) {
	tui.queueUpdateDraw(func() {
		tui.PageComponent.ReadLayout.SQLOutputFieldComponent.SetError(err)
		go time.AfterFunc(3*time.Second, tui.resetMessage)
	})
}

func (tui *TUI) showWriteSQLError(err error) {
	tui.queueUpdateDraw(func() {
		tui.PageComponent.WriteLayout.SQLOutputFieldComponent.SetError(err)
		go time.AfterFunc(3*time.Second, tui.resetMessage)
	})
}

func (tui *TUI) resetMessage() {
	tui.queueUpdateDraw(func() {
		tui.PageComponent.ReadLayout.SQLOutputFieldComponent.Clear()
		tui.PageComponent.WriteLayout.SQLOutputFieldComponent.Clear()
	})
}
