package write

import "github.com/rivo/tview"

type WriteLayout struct {
	SQLInputFieldComponent  *SQLInputFieldComponent
	SQLOutputFieldComponent *SQLOutputFieldComponent
}

func NewWriteLayout() *WriteLayout {
	return &WriteLayout{
		SQLInputFieldComponent:  NewSQLInputFieldComponent(),
		SQLOutputFieldComponent: NewSQLOutputFieldComponent(),
	}
}

func NewWriteFlex(wLayout *WriteLayout) *tview.Flex {
	var wFlex = tview.NewFlex()

	wFlex.SetDirection(tview.FlexRow)
	wFlex.AddItem(wLayout.SQLInputFieldComponent.View, 0, 1, false)
	wFlex.AddItem(wLayout.SQLOutputFieldComponent.View, 0, 1, false)

	return wFlex
}
