package table_grid

import "github.com/rivo/tview"

type TableGrid struct {
	Records *tview.Table
}

// Initialize table grid
func NewTableGrid(tableData [][]*string) *TableGrid {
	table := tview.NewTable()

	for i, row := range tableData {
		for j, col := range row {
			var cellValue string

			if col != nil {
				cellValue = *col
			}

			table.SetCell(
				i, j,
				tview.NewTableCell(cellValue),
			)
		}
	}

	table.SetTitle("Records")
	table.SetBorder(true)
	table.SetSelectable(true, true)

	return &TableGrid{
		Records: table,
	}
}
