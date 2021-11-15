package functions

func FifthListOnSelect(index int, list_item_name string, second string, run rune) {
	TextView.Clear()
	TextViewData = ""
	List6.Clear()

	List6.SetTitle("")
	List5Item = list_item_name
	ActivePathBox.SetText(List1Item + " -> " + List2Item + " -> " + List3Item + " -> " + List4Item + " -> " + List5Item)
	List6.AddItem("Logs", "", 0, nil)
}
