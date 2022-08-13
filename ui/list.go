package ui

import (
	"github.com/rivo/tview"
)

type TableListComponent struct {
	View *tview.List
}

/*
====================
Initialize table list
====================
*/
func NewTableListComponent() *TableListComponent {
	listView := tview.NewList()

	listView.ShowSecondaryText(false)
	listView.SetTitle("Tables")
	listView.SetTitleAlign(tview.AlignLeft)
	listView.SetBorder(true)

	return &TableListComponent{
		View: listView,
	}
}

/*
====================
Set table list
====================
*/
func (tbl *TableListComponent) SetTableListView(tables []string) {
	tbl.ResetTableListView()

	for _, table := range tables {
		tbl.View.AddItem(table, table, 0, nil)
	}
}

/*
====================
Clear table list
====================
*/
func (tbl *TableListComponent) ResetTableListView() {
	tbl.View.Clear()
}
