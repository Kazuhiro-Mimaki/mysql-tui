package flex

import "github.com/rivo/tview"

func NewFlex(dbTreeView, sqlInputField, tableGrid tview.Primitive) *tview.Flex {
	rightFlex := tview.NewFlex()
	rightFlex.SetDirection(tview.FlexRow)
	rightFlex.AddItem(sqlInputField, 0, 1, false) // tview.Primitive, fixedSize int, proportion int, focus bool
	rightFlex.AddItem(tableGrid, 0, 15, false)

	flex := tview.NewFlex()
	flex.AddItem(dbTreeView, 0, 1, false)
	flex.AddItem(rightFlex, 0, 7, false)

	return flex
}
