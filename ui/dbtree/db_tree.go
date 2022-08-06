package db_tree

import (
	"github.com/rivo/tview"
)

type DatabaseTree struct {
	TreeView *tview.TreeView
}

// Initialize database tree
func NewDatabaseTree(tables []string) *DatabaseTree {
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

	return &DatabaseTree{
		TreeView: treeView,
	}
}
