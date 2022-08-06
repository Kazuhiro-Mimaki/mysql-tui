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
	App *tview.Application

	// mysql connection
	mysql *mysql.MySQL

	// view components
	DatabaseTree *db_tree.DatabaseTree
	TableGrid    *table_grid.TableGrid
}

func main() {
	tui := NewTui()

	flex := flex.NewFlex(tui.DatabaseTree.TreeView, tview.NewBox().SetBorder(true).SetTitle("Query Input"), tui.TableGrid.Records)

	tui.setEvent()

	if err := tui.App.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		log.Println(err)
	}
}

func NewTui() *TUI {
	tui := &TUI{
		App:   tview.NewApplication(),
		mysql: mysql.NewMySQL(),
	}

	tui.DatabaseTree = db_tree.NewDatabaseTree(tui.mysql)
	tui.TableGrid = table_grid.NewTableGrid(tui.mysql)

	// when the table selected, update the view of table grid
	tui.DatabaseTree.TreeView.SetSelectedFunc(func(node *tview.TreeNode) {
		tui.DatabaseTree.TreeView.SetCurrentNode(node)
		tui.updateTable(tui.DatabaseTree.TreeView.GetCurrentNode().GetText())
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
