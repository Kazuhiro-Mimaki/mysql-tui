package write

import "github.com/rivo/tview"

type WriteLayout struct {
	Box1 *tview.Box
	Box2 *tview.Box
}

func NewWriteLayout() *WriteLayout {
	return &WriteLayout{
		Box1: tview.NewBox(),
		Box2: tview.NewBox(),
	}
}

func NewWriteFlex(rWrite *WriteLayout) *tview.Flex {
	var wFlex = tview.NewFlex()

	wFlex.SetDirection(tview.FlexRow)
	wFlex.AddItem(rWrite.Box1, 0, 1, false)
	wFlex.AddItem(rWrite.Box2, 0, 15, false)

	return wFlex
}
