// Where I create and initialize the TUI elements

package functions

import (
	"io/ioutil"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func OnHelp() {
	Help, err := ioutil.ReadFile("Help.txt")
	if err != nil {
		TextView.SetText("No data in the Help file or I couldn't access it")
	} else {
		TextView.SetText(string(Help))
	}
}

func OnSearch() {
	searchitem := SearchBox.GetText()

	SearchResult = []string{""}
	if searchitem == "" {
		TextView.SetText(TextViewData)
		TextView.ScrollToBeginning()
	} else {
		lines := strings.Split(TextViewData, "\n")
		for i := 0; i < len(lines); i++ {
			if strings.Contains(lines[i], searchitem) {
				SearchResult = append(SearchResult, lines[i])
			}
		}
		if len(SearchResult) > 0 {
			TextView.SetText(strings.Join(SearchResult, "\n"))
		} else {
			TextView.Clear()
		}
		SearchResult = nil
	}
}

func OnCopyMode() {
	App.
		SetRoot(CopyModeGrid, true).
		SetFocus(TextView).EnableMouse(false)
}

func TextViewOnExit(key tcell.Key) {
	App.EnableMouse(true)
	App.
		SetRoot(MainGrid, true).
		SetFocus(TextView)
}

func OnGoBack() {
	App.
		SetRoot(MainGrid, true).
		SetFocus(MainGrid).
		EnableMouse(true)
	TextView.SetText(TextViewData)
	TextView.ScrollToBeginning()
}

func KeyboardKeys(event *tcell.EventKey) *tcell.EventKey {

	if SearchBox.HasFocus() && event.Key() == tcell.KeyEnter {
		OnSearch()
	}
	// if SearchBox.HasFocus() {
	// 	return event
	// } else {
	// 	if event.Key() == tcell.KeyDown || event.Key() == tcell.KeyUp {
	// 		TextView.SetText(event.Name())
	// 		return event
	// 	} else if event.Key() == tcell.KeyTab {
	// 		App.SetFocus(List3)
	// 		return nil
	// 	} else if event.Key() == tcell.KeyCtrlC {
	// 		App.Stop()
	// 		return nil
	// 	} else {
	// 		TextView.Clear()
	// 		return nil
	// 	}
	// }
	return event
}

func Exit() {
	App.Stop()
}

func TextViewOnFocus(action tview.MouseAction, event *tcell.EventMouse) {
	TextView.SetBackgroundColor(tcell.ColorLightGrey)
}

func SetPath() {
	_, SelectedMGType := MGDropDown.GetCurrentOption()
	BasePath = ProvidedDirPath + SelectedMGType + "/"
}

func CreateUI() {
	// Modifying MGDropDown's attributes
	MGDropDown.SetLabel("Select a must-gather type:  ")

	// Modifying buttons' attributes
	ClusterInfoButton.SetBorder(false).SetBackgroundColor(tcell.ColorBlue)
	ClusterStatusButton.SetBorder(false).SetBackgroundColor(tcell.ColorBlue)
	FocusModeButton.SetBorder(false).SetBackgroundColor(tcell.ColorOrangeRed)
	// FocusModeButton.SetSelectedFunc(OnSearchMode)
	HelpButton.SetBorder(false).SetBackgroundColor(tcell.ColorGreen)
	ExitButton.SetBorder(false).SetBackgroundColor(tcell.ColorDarkRed)
	ExitButton.SetSelectedFunc(Exit)
	CopyModeButton.SetBorder(false).SetBackgroundColor(tcell.ColorOrangeRed)
	CopyModeButton.SetSelectedFunc(OnCopyMode)
	SearchButton.SetBorder(false).SetBackgroundColor(tcell.ColorGreen)
	SearchButton.SetSelectedFunc(OnSearch)
	GoBackButton.SetBorder(false).SetBackgroundColor(tcell.ColorDarkRed)
	GoBackButton.SetSelectedFunc(OnGoBack)

	// Setting SearchBox attributes
	SearchBox.SetInputCapture(KeyboardKeys)

	// setting TextView attributes
	TextView.SetDoneFunc(TextViewOnExit)
	TextView.
		SetDynamicColors(true).
		SetWordWrap(true)

	// Setting SearchBox attributes
	HelpButton.SetSelectedFunc(OnHelp)

	// inittializing the List1 widget with common attributes
	List1.
		SetBorder(true).
		SetTitle("Resources")
	List1.SetHighlightFullLine(true)
	List1.
		ShowSecondaryText(false).
		AddItem("Summary", "", 0, nil).
		AddItem("Configurations", "", 0, nil).
		AddItem("Projects", "", 0, nil).
		AddItem("Nodes", "", 0, nil).
		AddItem("Operators", "", 0, nil).
		AddItem("MCP", "", 0, nil).
		AddItem("MC", "", 0, nil).
		AddItem("PV", "", 0, nil).
		AddItem("CSR", "", 0, nil)
	List1.SetSelectedFunc(FirstListOnSelect)
	// inittializing the List4 widget with common attributes
	List2.
		SetBorder(true)
	List2.
		SetSelectedFunc(SecondListOnSelect).
		ShowSecondaryText(false).
		SetHighlightFullLine(true)

	// setting List1 widget common attributes
	List3.
		ShowSecondaryText(false).
		SetBorder(true)
	List3.
		SetSelectedFunc(ThirdListOnSelect).
		SetHighlightFullLine(true)

	// setting List1 widget common attributes
	List4.
		ShowSecondaryText(false).
		SetBorder(true)
	List4.
		SetSelectedFunc(FourthListOnSelect).
		SetHighlightFullLine(true)

	// adding Resources to the List6 and setting a separate "onselect" function for each item
	List5.
		ShowSecondaryText(false).
		SetBorder(true)
	List5.
		SetSelectedFunc(FifthListOnSelect).
		SetHighlightFullLine(true)

	// adding Resources to the List6 and setting a separate "onselect" function for each item
	List6.
		ShowSecondaryText(false).
		SetBorder(true)
	List6.SetHighlightFullLine(true)
	List6.SetSelectedFunc(SixthListOnSelect)

	MessageWindow.SetTitle("Loading")

	// create the MainGrid which will hold the initial widgets
	MainGrid.
		SetRows(rows...).
		SetColumns(columns...).
		SetBorders(true).
		AddItem(MGDropDown, 0, 0, 1, 4, 0, 0, false).
		AddItem(ActivePathBox, 0, 4, 1, 1, 0, 0, false).
		AddItem(HelpButton, 0, 5, 1, 1, 0, 0, false).
		AddItem(CopyModeButton, 0, 6, 1, 1, 0, 0, false).
		AddItem(ExitButton, 0, 7, 1, 1, 0, 0, false).
		AddItem(List1, 1, 0, 1, 1, 0, 0, false).
		AddItem(List2, 1, 1, 1, 3, 0, 0, false).
		AddItem(List3, 2, 0, 3, 1, 0, 0, false).
		AddItem(List4, 2, 1, 1, 2, 0, 0, false).
		AddItem(List5, 3, 1, 2, 2, 0, 0, false).
		AddItem(List6, 2, 3, 3, 1, 0, 0, false).
		AddItem(SearchBox, 4, 4, 1, 3, 0, 0, false).
		AddItem(SearchButton, 4, 7, 1, 1, 0, 0, false).
		AddItem(TextView, 1, 4, 3, 4, 0, 0, false)

	CopyModeGrid.
		SetRows(rows...).
		SetColumns(columns...).
		SetBorders(false).
		AddItem(TextView, 0, 0, 5, 8, 0, 0, false)
}
