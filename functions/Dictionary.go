package functions

import "github.com/rivo/tview"

var App = tview.NewApplication()
var MainGrid = tview.NewGrid()
var SearchModeGrid = tview.NewGrid()
var CopyModeGrid = tview.NewGrid()
var MGDropDown = tview.NewDropDown()
var ClusterInfoButton = tview.NewButton("Cluster Info")
var ClusterStatusButton = tview.NewButton("Cluster Status")
var HelpButton = tview.NewButton("Help")
var FocusModeButton = tview.NewButton("Focus Window")
var ExitButton = tview.NewButton("Exit")
var SearchBox = tview.NewInputField()
var CopyModeButton = tview.NewButton("Copy Mode")
var SearchButton = tview.NewButton("Search")
var GoBackButton = tview.NewButton("Go Back")
var ActivePathBox = tview.NewTextView()
var List1 = tview.NewList()
var List2 = tview.NewList()
var List3 = tview.NewList()
var List4 = tview.NewList()
var List5 = tview.NewList()
var List6 = tview.NewList()
var TextView = tview.NewTextView()
var MessageWindow = tview.NewModal()
var List1Item string = ""
var List2Item string = ""
var List3Item string = ""
var List4Item string = ""
var List5Item string = ""
var List6Item string = ""
var TextViewData string = ""
var SearchResult []string = []string{""}
var height = 1
var width = 15
var rows []int = []int{height, 20, 10, 0, height}
var columns []int = []int{width, width, width, width, 0, width, width, width}
var ProvidedDirPath = "empty"
var BasePath = "empty"

type Node struct {
	Kind string `yaml:"apiVersion"`
}
