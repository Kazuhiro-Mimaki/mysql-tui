package main

import (
	"log"
	"tui-dbms/database/mysql"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App           *tview.Application
	DatabaseTrees *tview.TreeView
	TableRecords  *tview.Table
}

func newBox(title string) *tview.Box {
	return tview.NewBox().SetBorder(true).SetTitle(title)
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

func setDatabaseTrees(tables []string) *tview.TreeView {
	root := tview.NewTreeNode("root")
	databaseTrees := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	for _, table := range tables {
		node := tview.NewTreeNode(table)
		root.AddChild(node)
	}

	return databaseTrees
}

func setTableRecords(tableData [][]*string) *tview.Table {
	table := tview.NewTable()

	for i, row := range tableData {
		for j, col := range row {
			var cellValue string

			if col != nil {
				cellValue = *col
			}

			table.SetCell(
				i, j,
				tview.NewTableCell(cellValue),
			)
		}
	}
	table.SetSelectable(true, true)

	return table
}

func main() {
	tables := mysql.GetTables()
	tableData := mysql.GetRecords()

	tui := &TUI{
		App:           tview.NewApplication(),
		DatabaseTrees: setDatabaseTrees(tables),
		TableRecords:  setTableRecords(tableData),
	}

	rightFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(newBox("Query Input"), 0, 1, false).
		AddItem(tui.TableRecords, 0, 11, false)

	flex := tview.NewFlex().
		// tview.Primitive, fixedSize int, proportion int, focus bool
		AddItem(tui.DatabaseTrees, 0, 1, false).
		AddItem(rightFlex, 0, 7, false)

	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			tui.setFocus(tui.TableRecords)
		}
		return event
	})

	if err := tui.App.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		log.Println(err)
	}
}
