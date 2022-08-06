package flex

import "github.com/rivo/tview"

func NewFlex(dbDropDown, dbTreeList, sqlInputField, tableGrid tview.Primitive) *tview.Flex {
	left := tview.NewFlex()
	left.SetDirection(tview.FlexRow)
	left.AddItem(dbDropDown, 0, 1, true) // tview.Primitive, fixedSize int, proportion int, focus bool
	left.AddItem(dbTreeList, 0, 15, false)

	right := tview.NewFlex()
	right.SetDirection(tview.FlexRow)
	right.AddItem(sqlInputField, 0, 1, false)
	right.AddItem(tableGrid, 0, 15, false)

	main := tview.NewFlex()
	main.AddItem(left, 0, 1, true)
	main.AddItem(right, 0, 7, false)

	return main
}
