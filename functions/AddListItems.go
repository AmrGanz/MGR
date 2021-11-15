package functions

import (
	"github.com/rivo/tview"
)

func AddListItems(list *tview.List) {
	if list == List3 {
		List3.Clear()
		List4.Clear()
		List5.Clear()
		List6.Clear()
		if List1Item == "Projects" {
			list.
				AddItem("YAML", "", 0, nil).
				AddItem("Events", "", 0, nil).
				AddItem("Pods", "", 0, nil).
				AddItem("Deployment", "", 0, nil).
				AddItem("DeploymentConfig", "", 0, nil).
				AddItem("Daemonset", "", 0, nil).
				AddItem("Services", "", 0, nil).
				AddItem("Routes", "", 0, nil).
				AddItem("Image Stream", "", 0, nil).
				AddItem("PVC", "", 0, nil).
				AddItem("ConfigMap", "", 0, nil).
				AddItem("Secrets", "", 0, nil)
		} else if List1Item == "Nodes" {
			list.
				AddItem("Nodes Status", "", 0, nil).
				AddItem("Utilization", "", 0, nil).
				AddItem("Nodes Info", "", 0, nil).
				AddItem("Used MC", "", 0, nil)
		}
	} else if list == List6 {
		if List3Item == "Pods" {
			list.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil).
				AddItem("JSON", "", 0, nil).
				AddItem("Logs", "", 0, nil)
		}
	}
}
