package functions

func FifthListOnSelect(index int, list_item_name string, second string, run rune) {
	TextView.Clear()
	TextViewData = ""
	SixthList.Clear()

	SixthList.SetTitle("")
	FifthListItem = list_item_name
	ActivePathBox.SetText(FirstListItem + " -> " + SecondListItem + " -> " + ThirdListItem + " -> " + FourthListItem + " -> " + FifthListItem)
	SixthList.AddItem("Logs", "", 0, nil)
}
