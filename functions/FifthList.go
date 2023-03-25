package functions

import (
	"io/ioutil"
	"strings"
)

func FifthListOnSelect(index int, list_item_name string, second string, run rune) {
	TextView.Clear()
	TextViewData = ""
	List6.Clear()

	List6.SetTitle("")
	List5Item = list_item_name
	ActivePathBox.SetText(List1Item + " > " + List2Item + " > " + List3Item + " > " + List4Item + " > " + List5Item)
	List6.SetTitle("Logs")
	Files, _ := ioutil.ReadDir(Namespaces_Path + List2Item + "/pods/" + List4Item + "/" + List5Item + "/" + List5Item + "/logs/")
	for i := range Files {
		List6.AddItem(strings.Split(Files[i].Name(), ".log")[0], "", 0, nil)
	}
}
