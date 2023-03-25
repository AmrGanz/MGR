// Where I create and initialize the TUI elements

package functions

import (
	"io/ioutil"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MGDropDownOnSelect() {
	TextView.Clear()
	List1.Clear()
	List2.Clear()
	List3.Clear()
	List4.Clear()
	List5.Clear()
	List6.Clear()
	TextViewData = ""
	// IsMGDir := 0
	if _, value := MGDropDown.GetCurrentOption(); value != "" {
		_, SelectedMGType := MGDropDown.GetCurrentOption()

		if strings.Contains(SelectedMGType, "openshift-release-dev-ocp") {
			List1.
				ShowSecondaryText(false).
				// AddItem("Navigate", "", 0, nil).
				AddItem("Summary", "", 0, nil).
				AddItem("OCP Version", "", 0, nil).
				AddItem("Configurations", "", 0, nil).
				AddItem("Projects", "", 0, nil).
				AddItem("Nodes", "", 0, nil).
				AddItem("Operators", "", 0, nil).
				AddItem("Installed Operators", "", 0, nil).
				AddItem("MCP", "", 0, nil).
				AddItem("MC", "", 0, nil).
				AddItem("PV", "", 0, nil).
				AddItem("CSR", "", 0, nil).
				AddItem("ETCD", "", 0, nil)
		} else if strings.Contains(SelectedMGType, "openshift4-ose-cluster-logging-operator") {
			List1.
				ShowSecondaryText(false).
				AddItem("Summary", "", 0, nil).
				AddItem("Projects", "", 0, nil).
				AddItem("Nodes", "", 0, nil).
				AddItem("PV", "", 0, nil).
				AddItem(Colors.Yellow+"Logging"+Colors.White, "", 0, nil)
		} else {
			TextView.SetText(Colors.Red + "Please make sure you have selected the correct MG Dir And/Or the supported MG image" + "\n\n" + "Click on the Help Buttion to get more details." + Colors.White)
		}
		MG_Path = ProvidedDirPath + SelectedMGType + "/"
		SetResourcesPath()
		/////////////////////////////
		// TimeStampFile, _ := os.Open(ProvidedDirPath + "/timestamp")
		// // Create new Scanner.
		// scanner := bufio.NewScanner(TimeStampFile)
		// result := []string{}
		// // Use Scan.
		// for scanner.Scan() {
		// 	line := scanner.Text()
		// 	// Append line to result.
		// 	result = append(result, line)
		// }
		// FirstLine := strings.Split(result[0], " ")
		// SecondLine := strings.Split(result[1], " ")
		// MGStartDate := strings.Join(FirstLine[0:2], " ")
		// MGEndDate := strings.Join(SecondLine[0:2], " ")
		// TextView.SetText(MGStartDate + "\n" + MGEndDate)
		/////////////////////////////

		// MG_Path = ProvidedDirPath + SelectedMGType + "/"
		// Files, _ = ioutil.ReadDir(MG_Path)
		// for i := range Files {
		// 	if Files[i].Name() == "namespaces" {
		// 		IsMGDir++
		// 	} else if Files[i].Name() == "cluster-scoped-resources" {
		// 		IsMGDir++
		// 	}
		// }
		// if IsMGDir > 1 {
		// 	// Initialize resource Dir/File paths after getting MG_Path value
		// 	SetResourcesPath()
		// 	ItIsClusterLoggingMG := "false"
		// 	Files, _ = ioutil.ReadDir(MG_Path)
		// 	for i := range Files {
		// 		if Files[i].Name() == "cluster-logging" {
		// 			ItIsClusterLoggingMG = "true"
		// 		}
		// 	}
		// 	if ItIsClusterLoggingMG == "true" {
		// 		List1.
		// 			ShowSecondaryText(false).
		// 			AddItem("Summary", "", 0, nil).
		// 			AddItem("Projects", "", 0, nil).
		// 			AddItem("Nodes", "", 0, nil).
		// 			AddItem("PV", "", 0, nil).
		// 			AddItem(Colors.Yellow+"Logging"+Colors.White, "", 0, nil)

		// 	} else {
		// 		List1.
		// 			ShowSecondaryText(false).
		// 			// AddItem("Navigate", "", 0, nil).
		// 			AddItem("Summary", "", 0, nil).
		// 			AddItem("OCP Version", "", 0, nil).
		// 			AddItem("Configurations", "", 0, nil).
		// 			AddItem("Projects", "", 0, nil).
		// 			AddItem("Nodes", "", 0, nil).
		// 			AddItem("Operators", "", 0, nil).
		// 			AddItem("Installed Operators", "", 0, nil).
		// 			AddItem("MCP", "", 0, nil).
		// 			AddItem("MC", "", 0, nil).
		// 			AddItem("PV", "", 0, nil).
		// 			AddItem("CSR", "", 0, nil)
		// 	}

		// }

	}
}

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
				line := strings.Replace(lines[i], searchitem, Colors.Red+searchitem+Colors.White, -1)
				SearchResult = append(SearchResult, line)
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
	MG_Path = ProvidedDirPath + SelectedMGType + "/"
}

func CreateUI() {
	// Modifying MGDropDown's attributes
	MGDropDown.SetLabel("Select MG type:  ")

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
		// SetWordWrap(true).
		SetWrap(true).
		SetChangedFunc(func() {
			App.Draw()
		})

	// Setting ActivePathBox attributes
	ActivePathBox.SetDynamicColors(true)

	// Setting SearchBox attributes
	HelpButton.SetSelectedFunc(OnHelp)

	// Set MGdropDown On Select Function
	// MGDropDown.SetSelectedFunc(MGDropDownOnSelect)

	// inittializing the List1 widget with common attributes
	List1.
		SetBorder(true).
		SetTitle("Resources")
	List1.SetHighlightFullLine(true)
	// List1.
	// 	ShowSecondaryText(false).
	// 	AddItem("Summary", "", 0, nil).
	// 	AddItem("OCP Version", "", 0, nil).
	// 	AddItem("Configurations", "", 0, nil).
	// 	AddItem("Projects", "", 0, nil).
	// 	AddItem("Nodes", "", 0, nil).
	// 	AddItem("Operators", "", 0, nil).
	// 	AddItem("MCP", "", 0, nil).
	// 	AddItem("MC", "", 0, nil).
	// 	AddItem("PV", "", 0, nil).
	// 	AddItem("CSR", "", 0, nil)
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

	TableGrid.
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
		AddItem(Table, 1, 4, 3, 4, 0, 0, false)

	CopyModeGrid.
		SetRows(rows...).
		SetColumns(columns...).
		SetBorders(false).
		AddItem(TextView, 0, 0, 5, 8, 0, 0, false)
}
