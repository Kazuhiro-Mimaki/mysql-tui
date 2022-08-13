package ui

import (
	"tui-dbms/ui/tablemode/read"
	"tui-dbms/ui/tablemode/write"

	"github.com/rivo/tview"
)

type PageComponent struct {
	View        *tview.Pages
	ReadLayout  *read.ReadLayout
	WriteLayout *write.WriteLayout
}

func NewPageComponent() *PageComponent {
	var pageView = tview.NewPages()

	var rLayout = read.NewReadLayout()
	var wLayout = write.NewWriteLayout()

	var rFlex = read.NewReadFlex(rLayout)
	var wFlex = write.NewWriteFlex(wLayout)

	pageView.AddPage("Read", rFlex, true, true)
	pageView.AddPage("Write", wFlex, true, false)

	return &PageComponent{
		View:        pageView,
		ReadLayout:  rLayout,
		WriteLayout: wLayout,
	}
}
