package main

import (
	"log"
	"tui-dbms/database/mysql"
	db_tree "tui-dbms/ui/dbtree"
	"tui-dbms/ui/flex"
	table_grid "tui-dbms/ui/tablegrid"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App          *tview.Application
	DatabaseTree *db_tree.DatabaseTree
	TableGrid    *table_grid.TableGrid
}

func main() {
	mysql := mysql.NewMySQL()
	tables := mysql.ShowTables()
	tableData := mysql.GetRecords(tables[0])

	tui := &TUI{
		App:          tview.NewApplication(),
		DatabaseTree: db_tree.NewDatabaseTree(tables),
		TableGrid:    table_grid.NewTableGrid(tableData),
	}

	flex := flex.NewFlex(tui.DatabaseTree.TreeView, tview.NewBox().SetBorder(true).SetTitle("Query Input"), tui.TableGrid.Records)

	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			tui.setFocus(tui.TableGrid.Records)
		}
		return event
	})

	if err := tui.App.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		log.Println(err)
	}
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
