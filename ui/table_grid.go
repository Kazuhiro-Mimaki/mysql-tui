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
	tables := database.ShowTables()

	table := tview.NewTable()
	table.SetSelectable(true, true).SetTitle("Records").SetBorder(true)

	tableGrid := &TableGrid{
		Records: table,
	}

	tableGrid.SetTableGrid(tables[0], database)

	return tableGrid
}

func (tg *TableGrid) SetTableGrid(table string, database mysql.IDatabaase) {
	tableData := database.GetRecords(table)

	tg.Records.Clear()

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
