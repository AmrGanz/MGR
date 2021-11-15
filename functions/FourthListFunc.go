package functions

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func FourthListOnSelect(index int, list_item_name string, second string, run rune) {
	TextView.Clear()
	TextViewData = ""
	List5.Clear()
	List6.Clear()

	List5.SetTitle("")
	List6.SetTitle("")
	List4Item = list_item_name
	ActivePathBox.SetText(List1Item + " -> " + List2Item + " -> " + List3Item + " -> " + List4Item)
	if List1Item == "Projects" && List2Item == "All Projects" {
		if List3Item == "Summary" {

		} else if List3Item == "Pods" {
			// Table of Projects/Pods
		} else if List3Item == "Deployment" {
			// Table of Projects/Deployment
		} else if List3Item == "DeploymentConfig" {
			// Table of Projects/DeploymentConfig
		} else if List3Item == "Daemonset" {
			// Table of Projects/Daemonset
		} else if List3Item == "Services" {
			// Table of Projects/Services
		} else if List3Item == "Routes" {
			// Table of Projects/Routes
		} else if List3Item == "Image Stream" {
			// Table of Projects/Image Stream
		} else if List3Item == "PVC" {
			// Table of Projects/PVC
		} else if List3Item == "ConfigMap" {
			// Table of Projects/ConfigMap
		} else if List3Item == "Secrets" {
			// Table of Projects/Secrets
		}
	} else if List1Item == "Projects" && List2Item != "All Projects" {
		if List3Item == "Summary" {

		} else if List3Item == "YAML" {
			// Table of Projects/YAML
		} else if List3Item == "Events" {
			// Table of Projects/Events
		} else if List3Item == "Pods" {
			// Add pod's containers to the Fifth list
			List5.Clear()
			filesList, _ := ioutil.ReadDir(BasePath + "namespaces/" + List2Item + "/pods/" + List4Item)
			if len(filesList) > 0 {
				for _, file := range filesList {
					if file.IsDir() {
						List5.AddItem(file.Name(), "", 0, nil)
					}
				}
			} else {
				List5.AddItem("Empty", "", 0, nil)
			}
			// Print Pod's YAML definition
			content, _ := os.ReadFile(BasePath + "namespaces/" + List2Item + "/pods/" + List4Item + "/" + List4Item + ".yaml")
			TextView.SetText(string(content))
			TextView.ScrollToBeginning()
			TextViewData = TextView.GetText(false)

		} else if List3Item == "Deployment" {
			// Table of Projects/Deployment
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		} else if List3Item == "DeploymentConfig" {
			// Table of Projects/DeploymentConfig
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		} else if List3Item == "Daemonset" {
			// Table of Projects/Daemonset
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		} else if List3Item == "Services" {
			// Table of Projects/Services
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		} else if List3Item == "Routes" {
			// Table of Projects/Routes
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		} else if List3Item == "Image Stream" {
			// Table of Projects/Image Stream
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		} else if List3Item == "PVC" {
			// Table of Projects/PVC
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		} else if List3Item == "ConfigMap" {
			// Table of Projects/ConfigMap
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		} else if List3Item == "Secrets" {
			// Table of Projects/Secrets
			List6.
				AddItem("Info", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		}
	} else if List1Item == "MC" && List3Item == "Data" {
		TextView.Clear()
		TextViewData = ""

		yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigs/" + List2Item + ".yaml")

		m := make(map[string]interface{})
		yaml.Unmarshal([]byte(yfile), m)

		paths := m["spec"].(map[interface{}]interface{})["config"].(map[interface{}]interface{})["storage"].(map[interface{}]interface{})["files"].([]interface{})
		for i := range paths {
			path := fmt.Sprintf("%v", paths[i].(map[interface{}]interface{})["path"])
			if path == List4Item {
				constents := fmt.Sprintf("%v", paths[i].(map[interface{}]interface{})["contents"].(map[interface{}]interface{})["source"])
				constents = strings.Trim(constents, "data:,")
				constents, _ = url.QueryUnescape(constents)
				TextView.SetText(constents)
				TextView.ScrollToBeginning()
				TextViewData = constents
				break
			}
		}
	}
}
