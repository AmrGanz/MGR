package functions

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var DirGrid = tview.NewGrid()
var DirTree = tview.NewTreeView()
var DirText = tview.NewTextView()
var DirButton = tview.NewButton("Select")
var DirApp = tview.NewApplication()
var rootDir = "/home"
var root = tview.NewTreeNode(rootDir).
	SetColor(tcell.ColorRed)

func DirSelect() {
	path := DirTree.GetCurrentNode()
	DirText.SetText(path.GetText())
}

func CreateDirUI() {
	DirTree.
		SetRoot(root).
		SetCurrentNode(root)

	DirButton.SetSelectedFunc(DirSelect)
	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		err = err

		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				SetSelectable(file.IsDir())
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	DirTree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	DirGrid.
		SetRows(1, 0, 1).
		SetColumns(10, 0, 10).
		SetBorders(true).
		AddItem(DirText, 0, 0, 1, 2, 0, 0, false).
		AddItem(DirButton, 0, 2, 1, 1, 0, 0, false).
		AddItem(DirTree, 1, 0, 2, 3, 0, 0, false)

	DirApp.SetRoot(DirGrid, true).EnableMouse(true).Run()

}
