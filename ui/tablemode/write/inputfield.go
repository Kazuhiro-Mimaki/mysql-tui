package write

import (
	"github.com/rivo/tview"
)

type SQLInputFieldComponent struct {
	View *tview.InputField
}

/*
====================
Initialize sql input field
====================
*/
func NewSQLInputFieldComponent() *SQLInputFieldComponent {
	inputFieldView := tview.NewInputField()

	inputFieldView.SetTitle("Query Input (INSERT / UPDATE / DELETE)")
	inputFieldView.SetTitleAlign(tview.AlignLeft)
	inputFieldView.SetBorder(true)

	return &SQLInputFieldComponent{
		View: inputFieldView,
	}
}
