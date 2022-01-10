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

	if List1Item == "Projects" {
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
				yaml.Unmarshal(yfile, m)

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
		} else if List2Item == "Nodes Details" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/core/nodes/")

			now := time.Now().UTC()

			Output := []string{"NAME" + "|" + "STATUS" + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "|" + "INTERNAL-IP" + "|" + "EXTERNAL-IP" + "|" + "OS-IMAGE" + "|" + "KERNEL-VERSION" + "|" + "CONTAINER-RUNTIME" + "\n"}
			for _, file := range files {
				yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/core/nodes/" + file.Name())
				m := make(map[string]interface{})
				yaml.Unmarshal(yfile, m)
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

				IPs := m["status"].(map[interface{}]interface{})["addresses"].([]interface{})
				internalIPS := ""
				externalIPS := ""
				for i := range IPs {
					if IPs[i].(map[interface{}]interface{})["type"] == "InternalIP" {
						internalIP := IPs[i].(map[interface{}]interface{})["address"]
						internalIPS = fmt.Sprintf("%v", internalIP)
					} else if IPs[i].(map[interface{}]interface{})["type"] == "ExternalIP" {
						externalIP := IPs[i].(map[interface{}]interface{})["address"]
						externalIPS = fmt.Sprintf("%v", externalIP)

					}
				}

				osImage := m["status"].(map[interface{}]interface{})["nodeInfo"].(map[interface{}]interface{})["osImage"]
				osImageS := fmt.Sprintf("%v", osImage)

				kernelVersion := m["status"].(map[interface{}]interface{})["nodeInfo"].(map[interface{}]interface{})["kernelVersion"]
				kernelVersionS := fmt.Sprintf("%v", kernelVersion)

				contRuntime := m["status"].(map[interface{}]interface{})["nodeInfo"].(map[interface{}]interface{})["containerRuntimeVersion"]
				contRuntimeS := fmt.Sprintf("%v", contRuntime)

				Output = append(Output, nameS+"|"+statusS+"|"+rolesS+"|"+age+"|"+versionS+"|"+internalIPS+"|"+externalIPS+"|"+osImageS+"|"+kernelVersionS+"|"+contRuntimeS+"\n")
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else {
			List3.
				AddItem("YAML", "", 0, nil).
				AddItem("Summary", "", 0, nil)
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
	}
}
