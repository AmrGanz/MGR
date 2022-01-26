package functions

import (
	"fmt"
	"io/ioutil"
)

func SecondListOnSelect(index int, list_item_name string, second string, run rune) {
	TextView.Clear()
	TextViewData = ""
	TextView.ScrollToBeginning()
	List3.Clear()
	List4.Clear()
	List5.Clear()
	List6.Clear()

	List3.SetTitle("")
	List4.SetTitle("")
	List5.SetTitle("")
	List6.SetTitle("")
	// List3Item, _ = List1.GetItemText(index)
	List2Item = list_item_name
	ActivePathBox.SetText(List1Item + " -> " + List2Item)
	if List1Item == "Configurations" {
		file, err := ioutil.ReadFile(BasePath + "cluster-scoped-resources/config.openshift.io/" + List2Item + ".yaml")
		if err != nil {
			fmt.Print(err)
		} else {
			TextView.SetText(string(file))
		}
	} else if List1Item == "Projects" {
		if List2Item == "All Projects" {
			List3.Clear()
			List4.Clear()
			List5.Clear()
			List6.Clear()
			List3.
				// AddItem("Summary", "", 0, nil).
				AddItem("Pods", "", 0, nil).
				AddItem("Deployment", "", 0, nil).
				AddItem("DeploymentConfig", "", 0, nil).
				AddItem("Daemonset", "", 0, nil).
				AddItem("Services", "", 0, nil).
				AddItem("Routes", "", 0, nil).
				AddItem("Image Stream", "", 0, nil).
				AddItem("PVC", "", 0, nil).
				AddItem("ConfigMap", "", 0, nil).
				AddItem("Secrets", "", 0, nil).
				AddItem("Subscriptions", "", 0, nil).
				AddItem("Operators", "", 0, nil)
		} else if List2Item == "Empty" {
			// Do nothing
		} else {
			List3.
				// AddItem("Summary", "", 0, nil).
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
				AddItem("Secrets", "", 0, nil).
				AddItem("Subscriptions", "", 0, nil).
				AddItem("Operators", "", 0, nil)
		}
	} else if List1Item == "Nodes" {
		if List2Item == "Nodes Summary" {
			GetNodesInfo("All Nodes", "Summary")
		} else if List2Item == "Nodes Details" {
			GetNodesInfo("All Nodes", "Details")
		} else {
			GetNodesInfo(List2Item, "Summary")
			List3.
				AddItem("Summary", "", 0, nil).
				AddItem("Details", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		}
	} else if List1Item == "Operators" {
		List3.SetTitle("Info")
		List3.
			AddItem("YAML", "", 0, nil).
			AddItem("Summary", "", 0, nil)
	} else if List1Item == "MCP" {
		List3.SetTitle("Info")
		List3.
			AddItem("Info", "", 0, nil).
			AddItem("YAML", "", 0, nil)
	} else if List1Item == "MC" {
		List3.SetTitle("Info")
		List3.
			AddItem("Summary", "", 0, nil).
			AddItem("YAML", "", 0, nil).
			AddItem("Data", "", 0, nil)
	} else if List1Item == "PV" {
		List3.SetTitle("Info")
		List3.
			AddItem("Info", "", 0, nil).
			AddItem("YAML", "", 0, nil)
	} else if List1Item == "CSR" {
		if List2Item == "All Certificate Signing Requests" {
			GetCSRInfo()
		} else {
			File, _ = ioutil.ReadFile(BasePath + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/" + List2Item + ".yaml")
			TextView.SetText(string(File))
			TextView.ScrollToBeginning()
			TextViewData = TextView.GetText(false)
		}
	}
}
