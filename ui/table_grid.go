package ui

import (
	"tui-dbms/database/mysql"

	"github.com/gdamore/tcell/v2"
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
			var cellColor = tcell.ColorWhite
			var notSelectable = false

			if col != nil {
				cellValue = *col
			}

			// top(column names)
			if i == 0 {
				cellColor = tcell.ColorNavy
			}

			tg.Records.SetCell(
				i, j,
				&tview.TableCell{
					Text:          cellValue,
					Color:         cellColor,
					NotSelectable: notSelectable,
				},
			)
		}
	}
}

// clear records and scroll to beginning
func (tg *TableGrid) ResetRecords() {
	tg.Records.Clear().ScrollToBeginning()
}
