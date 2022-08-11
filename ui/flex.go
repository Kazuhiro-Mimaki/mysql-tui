package ui

import "github.com/rivo/tview"

type FlexLayout struct {
	Main  *tview.Flex
	Left  *LeftFlex
	Right *RightFlex
}

type LeftFlex struct {
	Layout *tview.Flex
}

type RightFlex struct {
	Layout *tview.Flex
}

/*
====================
Initialize flex layout
====================
*/
func NewMainFlex(dbDropDown, tableList, sqlInputField, tableGrid tview.Primitive) *FlexLayout {
	left := NewLeftFlex(dbDropDown, tableList)
	right := NewRightFlex(sqlInputField, tableGrid)

	main := tview.NewFlex()
	main.AddItem(left.Layout, 0, 1, true)
	main.AddItem(right.Layout, 0, 7, true)

	return &FlexLayout{
		Main:  main,
		Left:  left,
		Right: right,
	}
}

/*
====================
Initialize left flex layout
====================
*/
func NewLeftFlex(dbDropDown, tableList tview.Primitive) *LeftFlex {
	leftLayout := tview.NewFlex()
	leftLayout.SetDirection(tview.FlexRow)
	leftLayout.AddItem(dbDropDown, 0, 1, true) // tview.Primitive, fixedSize int, proportion int, focus bool
	leftLayout.AddItem(tableList, 0, 15, false)

	return &LeftFlex{
		Layout: leftLayout,
	}
}

/*
====================
Initialize right flex layout
====================
*/
func NewRightFlex(sqlInputField, tableGrid tview.Primitive) *RightFlex {
	rightLayout := tview.NewFlex()
	rightLayout.SetDirection(tview.FlexRow)
	rightLayout.AddItem(sqlInputField, 0, 1, false)
	rightLayout.AddItem(tableGrid, 0, 15, false)

	return &RightFlex{
		Layout: rightLayout,
	}
}
