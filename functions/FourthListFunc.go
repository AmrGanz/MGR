package functions

import (
	"io/ioutil"
	"os"
)

func FourthListOnSelect(index int, list_item_name string, second string, run rune) {
	TextView.Clear()
	TextViewData = ""
	FifthList.Clear()
	SixthList.Clear()

	FifthList.SetTitle("")
	SixthList.SetTitle("")
	FourthListItem = list_item_name
	ActivePathBox.SetText(FirstListItem + " -> " + SecondListItem + " -> " + ThirdListItem + " -> " + FourthListItem)
	if FirstListItem == "Projects" && SecondListItem == "All Projects" {
		if ThirdListItem == "Summary" {

		} else if ThirdListItem == "Pods" {
			// Table of Projects/Pods
		} else if ThirdListItem == "Deployment" {
			// Table of Projects/Deployment
		} else if ThirdListItem == "DeploymentConfig" {
			// Table of Projects/DeploymentConfig
		} else if ThirdListItem == "Daemonset" {
			// Table of Projects/Daemonset
		} else if ThirdListItem == "Services" {
			// Table of Projects/Services
		} else if ThirdListItem == "Routes" {
			// Table of Projects/Routes
		} else if ThirdListItem == "Image Stream" {
			// Table of Projects/Image Stream
		} else if ThirdListItem == "PVC" {
			// Table of Projects/PVC
		} else if ThirdListItem == "ConfigMap" {
			// Table of Projects/ConfigMap
		} else if ThirdListItem == "Secrets" {
			// Table of Projects/Secrets
		}
	} else if FirstListItem == "Projects" && SecondListItem != "All Projects" {
		if ThirdListItem == "Summary" {

		} else if ThirdListItem == "YAML" {
			// Table of Projects/YAML
		} else if ThirdListItem == "Events" {
			// Table of Projects/Events
		} else if ThirdListItem == "Pods" {
			// Add pod's containers to the Fifth list
			FifthList.Clear()
			filesList, _ := ioutil.ReadDir(BasePath + "namespaces/" + SecondListItem + "/pods/" + FourthListItem)
			if len(filesList) > 0 {
				for _, file := range filesList {
					if file.IsDir() {
						FifthList.AddItem(file.Name(), "", 0, nil)
					}
				}
			} else {
				FifthList.AddItem("Empty", "", 0, nil)
			}
			// Print Pod's YAML definition
			content, _ := os.ReadFile(BasePath + "namespaces/" + SecondListItem + "/pods/" + FourthListItem + "/" + FourthListItem + ".yaml")
			TextView.SetText(string(content))
			TextView.ScrollToBeginning()
			TextViewData = TextView.GetText(false)

		} else if ThirdListItem == "Deployment" {
			// Table of Projects/Deployment
		} else if ThirdListItem == "DeploymentConfig" {
			// Table of Projects/DeploymentConfig
		} else if ThirdListItem == "Daemonset" {
			// Table of Projects/Daemonset
		} else if ThirdListItem == "Services" {
			// Table of Projects/Services
		} else if ThirdListItem == "Routes" {
			// Table of Projects/Routes
		} else if ThirdListItem == "Image Stream" {
			// Table of Projects/Image Stream
		} else if ThirdListItem == "PVC" {
			// Table of Projects/PVC
		} else if ThirdListItem == "ConfigMap" {
			// Table of Projects/ConfigMap
		} else if ThirdListItem == "Secrets" {
			// Table of Projects/Secrets
		}
	}
}
