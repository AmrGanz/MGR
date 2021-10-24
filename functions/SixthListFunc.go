package functions

func SixthListOnSelect(index int, list_item_name string, second string, run rune) {
	SixthListItem = list_item_name
	TextView.Clear()
	ActivePathBox.SetText(FirstListItem + " -> " + SecondListItem + " -> " + ThirdListItem + " -> " + FourthListItem + " -> " + SixthListItem)
	if FirstListItem == "Projects" && ThirdListItem == "Pods" && SixthListItem == "Logs" {
		TextView.SetText("TBA")
		TextView.ScrollToBeginning()
	}

}
