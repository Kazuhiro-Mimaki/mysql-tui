package main

import (
	"tui-dbms/ui/treeview"

	"github.com/rivo/tview"
)

func FlexUi(viewTable *tview.Table) {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		// tview.Primitive, fixedSize int, proportion int, focus bool
		AddItem(newBox("database/table tree"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(newBox("some query"), 0, 1, false).
			AddItem(viewTable, 0, 11, false),
			0, 7, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	treeview.TreeViewUi()
}

func newBox(title string) *tview.Box {
	return tview.NewBox().SetBorder(true).SetTitle(title)
}

func main() {
	table := tview.NewTable()
	FlexUi(table)
}
