package read

import "github.com/rivo/tview"

type ReadLayout struct {
	SQLInputFieldComponent  *SQLInputFieldComponent
	SQLOutputFieldComponent *SQLOutputFieldComponent
	TableGridComponent      *TableGridComponent
}

func NewReadLayout() *ReadLayout {
	return &ReadLayout{
		SQLInputFieldComponent:  NewSQLInputFieldComponent(),
		SQLOutputFieldComponent: NewSQLOutputFieldComponent(),
		TableGridComponent:      NewTableGridComponent(),
	}
}

func NewReadFlex(rLayout *ReadLayout) *tview.Flex {
	var rFlex = tview.NewFlex()

	rFlex.SetDirection(tview.FlexRow)
	rFlex.AddItem(rLayout.SQLInputFieldComponent.View, 0, 1, false)
	rFlex.AddItem(rLayout.SQLOutputFieldComponent.View, 0, 1, false)
	rFlex.AddItem(rLayout.TableGridComponent.View, 0, 14, false)

	return rFlex
}
