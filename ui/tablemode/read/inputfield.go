package read

import (
	"github.com/rivo/tview"
)

type SQLInputFieldComponent struct {
	View *tview.InputField
}

/*
====================
Initialize db dropdown
====================
*/
func NewSQLInputFieldComponent() *SQLInputFieldComponent {
	inputFieldView := tview.NewInputField()

	inputFieldView.SetTitle("Query Input (SELECT)")
	inputFieldView.SetTitleAlign(tview.AlignLeft)
	inputFieldView.SetBorder(true)

	return &SQLInputFieldComponent{
		View: inputFieldView,
	}
}
