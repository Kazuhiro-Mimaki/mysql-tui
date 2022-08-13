package read

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SQLOutputFieldComponent struct {
	View *tview.TextView
}

/*
====================
Initialize sql output text view
====================
*/
func NewSQLOutputFieldComponent() *SQLOutputFieldComponent {
	outputTextView := tview.NewTextView()

	outputTextView.SetTitle("Query Result")
	outputTextView.SetTitleAlign(tview.AlignLeft)
	outputTextView.SetBorder(true)
	outputTextView.SetTextColor(tcell.ColorRed)

	return &SQLOutputFieldComponent{
		View: outputTextView,
	}
}

func (output *SQLOutputFieldComponent) SetError(err error) {
	output.View.SetText(err.Error())
}

func (output *SQLOutputFieldComponent) Clear() {
	output.View.Clear()
}
