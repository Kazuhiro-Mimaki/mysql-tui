package main

import (
	"log"
	"tui-dbms/database/mysql"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App *tview.Application
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

func main() {
	tui := &TUI{
		App: tview.NewApplication(),
	}

	tableData := mysql.Driver()
	viewTable := tview.NewTable()

	mysql.ShowData(viewTable, tableData)

	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			tui.App.Stop()
		})

	rightFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(newBox("some query"), 0, 1, false).
		AddItem(viewTable, 0, 11, false)

	flex := tview.NewFlex().
		// tview.Primitive, fixedSize int, proportion int, focus bool
		AddItem(list, 0, 1, false).
		AddItem(rightFlex, 0, 7, false)

	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			tui.setFocus(viewTable)
		}
		return event
	})

	if err := tui.App.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		log.Println(err)
	}
}
