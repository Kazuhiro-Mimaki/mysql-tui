package ui

import (
	"tui-dbms/database/mysql"

	"github.com/rivo/tview"
)

type TableGrid struct {
	Records *tview.Table
}

// Initialize table grid
func NewTableGrid(database mysql.IDatabaase) *TableGrid {
	table := tview.NewTable()

	table.SetSelectable(true, true)
	table.SetTitle("Records")
	table.SetBorder(true)

	// fix column names
	table.SetFixed(1, 0)

	tableGrid := &TableGrid{
		Records: table,
	}

	return tableGrid
}

func (tg *TableGrid) SetTableGrid(table string, database mysql.IDatabaase) {
	tg.ResetRecords()

	tableData := database.GetRecords(table)

	for i, row := range tableData {
		for j, col := range row {
			var cellValue string

			if col != nil {
				cellValue = *col
			}

			tg.Records.SetCell(
				i, j,
				tview.NewTableCell(cellValue),
			)
		}
	}
}

// clear records and scroll to beginning
func (tg *TableGrid) ResetRecords() {
	tg.Records.Clear().ScrollToBeginning()
}
