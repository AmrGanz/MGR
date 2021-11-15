package functions

import (
	"io/ioutil"
	"strings"

	"github.com/rivo/tview"
)

func AddDirItemsToList(Item string, list *tview.List) {

	if Item == "Projects" {
		fileInfo, _ := ioutil.ReadDir(BasePath + "namespaces/")
		if len(fileInfo) > 0 {
			List2.AddItem("All", "", 0, nil)
			for _, info := range fileInfo {
				projectname := strings.Split(info.Name(), ".yaml")
				list.AddItem(projectname[0], "", 0, nil)

			}
		} else {
			list.AddItem("Empty", "", 0, nil)
		}
	}

}
