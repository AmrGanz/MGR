package functions

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/ryanuber/columnize"
	"gopkg.in/yaml.v2"
)

func SecondListOnSelect(index int, list_item_name string, second string, run rune) {
	TextView.Clear()
	TextViewData = ""
	TextView.ScrollToBeginning()
	ThirdList.Clear()
	FourthList.Clear()
	SixthList.Clear()

	ThirdList.SetTitle("")
	FourthList.SetTitle("")
	FifthList.SetTitle("")
	SixthList.SetTitle("")
	// ThirdListItem, _ = FirstList.GetItemText(index)
	SecondListItem = list_item_name
	ActivePathBox.SetText(FirstListItem + " -> " + SecondListItem)

	if FirstListItem == "Projects" {
		if SecondListItem == "All Projects" {
			ThirdList.Clear()
			FourthList.Clear()
			FifthList.Clear()
			SixthList.Clear()
			ThirdList.
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
		} else {
			ThirdList.
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
	} else if FirstListItem == "Nodes" {
		if SecondListItem == "Nodes Summary" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/core/nodes/")

			// Print a summerized nodes status
			now := time.Now().UTC()
			Output := []string{"NAME" + "|" + "STATUS" + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "\n"}
			for _, file := range files {
				yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/core/nodes/" + file.Name())
				m := make(map[string]interface{})
				yaml.Unmarshal([]byte(yfile), m)

				name := m["metadata"].(map[interface{}]interface{})["name"]
				nameS := fmt.Sprintf("%v", name)

				status := m["status"].(map[interface{}]interface{})["conditions"].([]interface{})
				statusS := ""
				for i := 0; i < len(status); i++ {
					if status[i].(map[interface{}]interface{})["type"] == "Ready" {
						if status[i].(map[interface{}]interface{})["status"] == "True" {
							statusS = "Ready"
						} else {
							statusS = "NotReady"
						}

					}
				}
				schedulable := m["spec"].(map[interface{}]interface{})["unschedulable"]
				schedulableS := fmt.Sprintf("%v", schedulable)
				if schedulableS == "true" {
					statusS += ",SchedulingDisabled"
				}
				rolesS := ""
				roles := m["metadata"].(map[interface{}]interface{})["labels"].(map[interface{}]interface{})
				for role := range roles {
					rolesStr := fmt.Sprintf("%v", role)
					if strings.Contains(rolesStr, "node-role.kubernetes.io") {
						label := fmt.Sprintf("%v", role)
						rolesS += strings.Split(label, "/")[1]
						rolesS += " "
					}
				}

				CreationTime := m["metadata"].(map[interface{}]interface{})["creationTimestamp"]
				CreationTimeS := fmt.Sprintf("%v", CreationTime)
				t1, _ := time.Parse(time.RFC3339, CreationTimeS)
				diff := now.Sub(t1).Seconds()
				diffI := int(diff)
				seconds := strconv.Itoa((diffI % 60))
				minutes := strconv.Itoa((diffI / 60) % 60)
				hours := strconv.Itoa((diffI / 360) % 24)
				days := strconv.Itoa((diffI / 86400))
				age := ""
				if days != "0" {
					age = days + "d" + hours + "h"
				} else if days == "0" && hours != "" {
					age = hours + "h" + minutes + "m"
				} else if hours == "0" {
					age = minutes + "m" + seconds + "s"
				}

				version := m["status"].(map[interface{}]interface{})["nodeInfo"].(map[interface{}]interface{})["kubeletVersion"]
				versionS := fmt.Sprintf("%v", version)
				Output = append(Output, nameS+"|"+statusS+"|"+rolesS+"|"+age+"|"+versionS+"\n")
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else if SecondListItem == "Nodes Details" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/core/nodes/")

			// Print a summerized nodes status
			now := time.Now().UTC()
			Output := []string{"NAME" + "|" + "STATUS" + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "|" + "INTERNAL-IP" + "|" + "EXTERNAL-IP" + "|" + "OS-IMAGE" + "|" + "KERNEL-VERSION" + "|" + "CONTAINER-RUNTIME" + "\n"}
			for _, file := range files {
				yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/core/nodes/" + file.Name())
				m := make(map[string]interface{})
				yaml.Unmarshal([]byte(yfile), m)

				name := m["metadata"].(map[interface{}]interface{})["name"]
				nameS := fmt.Sprintf("%v", name)

				status := m["status"].(map[interface{}]interface{})["conditions"].([]interface{})
				statusS := ""
				for i := 0; i < len(status); i++ {
					if status[i].(map[interface{}]interface{})["type"] == "Ready" {
						if status[i].(map[interface{}]interface{})["status"] == "True" {
							statusS = "Ready"
						} else {
							statusS = "NotReady"
						}

					}
				}
				schedulable := m["spec"].(map[interface{}]interface{})["unschedulable"]
				schedulableS := fmt.Sprintf("%v", schedulable)
				if schedulableS == "true" {
					statusS += ",SchedulingDisabled"
				}

				rolesS := ""
				roles := m["metadata"].(map[interface{}]interface{})["labels"].(map[interface{}]interface{})
				for role := range roles {
					rolesStr := fmt.Sprintf("%v", role)
					if strings.Contains(rolesStr, "node-role.kubernetes.io") {
						label := fmt.Sprintf("%v", role)
						rolesS += strings.Split(label, "/")[1]
						rolesS += " "
					}
				}

				CreationTime := m["metadata"].(map[interface{}]interface{})["creationTimestamp"]
				CreationTimeS := fmt.Sprintf("%v", CreationTime)
				t1, _ := time.Parse(time.RFC3339, CreationTimeS)
				diff := now.Sub(t1).Seconds()
				diffI := int(diff)
				seconds := strconv.Itoa((diffI % 60))
				minutes := strconv.Itoa((diffI / 60) % 60)
				hours := strconv.Itoa((diffI / 360) % 24)
				days := strconv.Itoa((diffI / 86400))
				age := ""
				if days != "0" {
					age = days + "d" + hours + "h"
				} else if days == "0" && hours != "" {
					age = hours + "h" + minutes + "m"
				} else if hours == "0" {
					age = minutes + "m" + seconds + "s"
				}

				version := m["status"].(map[interface{}]interface{})["nodeInfo"].(map[interface{}]interface{})["kubeletVersion"]
				versionS := fmt.Sprintf("%v", version)
				Output = append(Output, nameS+"|"+statusS+"|"+rolesS+"|"+age+"|"+versionS+"\n")
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else {
			ThirdList.
				AddItem("YAML", "", 0, nil).
				AddItem("Summary", "", 0, nil)
		}
	} else if FirstListItem == "Operators" {
		ThirdList.SetTitle("Info")
		ThirdList.
			AddItem("YAML", "", 0, nil).
			AddItem("Summary", "", 0, nil)
	} else if FirstListItem == "MCP" {
		ThirdList.SetTitle("Info")
		ThirdList.
			AddItem("Info", "", 0, nil).
			AddItem("YAML", "", 0, nil)
	} else if FirstListItem == "MC" {
		ThirdList.SetTitle("Info")
		ThirdList.
			AddItem("Info", "", 0, nil).
			AddItem("YAML", "", 0, nil)
	} else if FirstListItem == "PV" {
		ThirdList.SetTitle("Info")
		ThirdList.
			AddItem("Info", "", 0, nil).
			AddItem("YAML", "", 0, nil)
	}
}
