package write

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

	return &SQLOutputFieldComponent{
		View: outputTextView,
	}
}

func (output *SQLOutputFieldComponent) SetError(err error) {
	output.View.SetTextColor(tcell.ColorRed)
	output.View.SetText(err.Error())
}

func (output *SQLOutputFieldComponent) SetSuccessMessage(msg string) {
	output.View.SetTextColor(tcell.ColorGreen)
	output.View.SetText(msg)
}

func (output *SQLOutputFieldComponent) Clear() {
	output.View.Clear()
}
