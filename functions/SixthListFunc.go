package functions

func SixthListOnSelect(index int, list_item_name string, second string, run rune) {
	SixthListItem = list_item_name
	ActivePathBox.SetText(FirstListItem + " -> " + SecondListItem + " -> " + ThirdListItem + " -> " + FourthListItem + " -> " + SixthListItem)
	if FirstListItem == "Projects" && ThirdListItem == "Pods" && SixthListItem == "YAML" {
		// content, _ := os.ReadFile(BasePath + "namespaces/" + SecondListItem + "/pods/" + FourthListItem + "/" + FourthListItem + ".yaml")
		// TextView.SetText(string(content))
		// TextView.ScrollToBeginning()
		// TextViewData = TextView.GetText(false)
	} else if FirstListItem == "Projects" && ThirdListItem == "Pods" && SixthListItem == "Logs" {
		// TextView.SetText("To Be Implemented")
		// content, _ := os.ReadFile(BasePath + "namespaces/" + SecondListItem + "/pods/" + FourthListItem + "/" + FifthListItem + "/" + FifthListItem + "/logs/current.log")
		// TextView.SetText(string(content))
		// TextView.ScrollToBeginning()
		// TextViewData = TextView.GetText(false)
	}

}
