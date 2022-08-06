package db_tree

import (
	"tui-dbms/database/mysql"

	"github.com/rivo/tview"
)

type DatabaseTree struct {
	TreeView *tview.TreeView
}

// Initialize database tree
func NewDatabaseTree(database mysql.IDatabaase) *DatabaseTree {
	tables := database.ShowTables()

	rootNode := tview.NewTreeNode("default database")
	treeView := tview.NewTreeView()

	for _, tableName := range tables {
		childNode := tview.NewTreeNode(tableName)
		rootNode.AddChild(childNode)
	}

	treeView.SetRoot(rootNode)
	treeView.SetCurrentNode(rootNode)
	treeView.SetTitle("Database trees")
	treeView.SetBorder(true)

	dbTree := &DatabaseTree{
		TreeView: treeView,
	}

	return dbTree
}
