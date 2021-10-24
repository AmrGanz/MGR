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

func FirstListOnSelect(index int, list_item_name string, second string, run rune) {
	// Setting current choice in the ActivePathBox
	TextView.Clear()
	TextViewData = ""
	SecondList.Clear()
	ThirdList.Clear()
	FourthList.Clear()
	FifthList.Clear()
	SixthList.Clear()
	SecondList.SetTitle("")
	ThirdList.SetTitle("")
	FourthList.SetTitle("")
	FifthList.SetTitle("")
	SixthList.SetTitle("")

	// Setting the global variable FirstListItem
	FirstListItem = list_item_name

	// Adding selection path to the ActivePathBox
	ActivePathBox.SetText(FirstListItem + " -> ")
	if list_item_name == "Projects" {
		SecondList.SetTitle("Projects")
		namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
		if len(namespaces) > 0 {
			SecondList.AddItem("All Projects", "", 0, nil)
			for _, project := range namespaces {
				SecondList.AddItem(project.Name(), "", 0, nil)
			}
		} else {
			SecondList.AddItem("Empty", "", 0, nil)
		}
	} else if list_item_name == "Nodes" {
		SecondList.SetTitle("Nodes")
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/core/nodes/")
		SecondList.
			AddItem("Nodes Summary", "", 0, nil).
			AddItem("Nodes Details", "", 0, nil)
		// Populate SecondList with nodes names
		for _, node := range files {
			nodeName := strings.Split(node.Name(), ".yaml")
			SecondList.AddItem(nodeName[0], "", 0, nil)
		}

	} else if list_item_name == "Operators" {
		SecondList.SetTitle("Operators")
		clusteroperators, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/config.openshift.io/clusteroperators.yaml")
		Output := []string{"NAME" + "|" + "VERSION" + "|" + "AVAILABLE" + "|" + "PROGRESSINS" + "|" + "DEGRADED" + "|" + "SINCE" + "\n"}
		m := make(map[interface{}]interface{})
		yaml.Unmarshal([]byte(clusteroperators), m)
		items, _ := m["items"].([]interface{})
		for i := range items {
			operator := items[i].(map[interface{}]interface{})
			name := operator["metadata"].(map[interface{}]interface{})["name"]
			nameS := fmt.Sprintf("%v", name)
			SecondList.AddItem(nameS, "", 0, nil)

			versions := operator["status"].(map[interface{}]interface{})["versions"].([]interface{})
			versionS := ""
			for i := range versions {
				if versions[i].(map[interface{}]interface{})["name"] == "operator" {
					version := versions[i].(map[interface{}]interface{})["version"]
					versionS = fmt.Sprintf("%v", version)
				}
			}
			status := operator["status"].(map[interface{}]interface{})["conditions"].([]interface{})
			availableS := ""
			progressingS := ""
			degradedS := ""
			availableSince := ""
			for i := range status {
				if status[i].(map[interface{}]interface{})["type"] == "Available" {
					if status[i].(map[interface{}]interface{})["status"] == "True" {
						availableS = "True"
					} else {
						availableS = "False"
					}

					now := time.Now().UTC()
					statusTime := status[i].(map[interface{}]interface{})["lastTransitionTime"]
					statusTimeS := fmt.Sprintf("%v", statusTime)
					t1, _ := time.Parse(time.RFC3339, statusTimeS)
					diff := now.Sub(t1).Seconds()
					diffI := int(diff)
					seconds := strconv.Itoa((diffI % 60))
					minutes := strconv.Itoa((diffI / 60) % 60)
					hours := strconv.Itoa((diffI / 360) % 24)
					days := strconv.Itoa((diffI / 86400))

					if days != "0" {
						availableSince = days + "d" + hours + "h"
					} else if days == "0" && hours != "" {
						availableSince = hours + "h" + minutes + "m"
					} else if hours == "0" {
						availableSince = minutes + "m" + seconds + "s"
					}

				} else if status[i].(map[interface{}]interface{})["type"] == "Progressing" {
					if status[i].(map[interface{}]interface{})["status"] == "True" {
						progressingS = "True"
					} else {
						progressingS = "False"
					}
				} else if status[i].(map[interface{}]interface{})["type"] == "Degraded" {
					if status[i].(map[interface{}]interface{})["status"] == "True" {
						degradedS = "True"
					} else {
						degradedS = "False"
					}
				}

			}
			// fmt.Print(nameS + "\t" + versionS + "\t" + availableS + "\t" + progressingS + "\t" + degradedS + "\t" + availableSince + "\n")
			Output = append(Output, nameS+"|"+versionS+"|"+availableS+"|"+progressingS+"|"+degradedS+"|"+availableSince+"\n")
		}
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if list_item_name == "MCP" {
		SecondList.SetTitle("MCP")
	} else if list_item_name == "PV" {
		SecondList.SetTitle("PV")
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		now := time.Now().UTC()
		Output := []string{"NAME" + "|" + "CAPACITY" + "|" + "ACCESS MODE" + "|" + "RECLAIM POLICY" + "|" + "STATUS" + "|" + "CLAIM" + "|" + "STORAGECLASS" + "|" + "AGE" + "\n"}
		files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/core/persistentvolumes/")
		for _, file := range files {
			yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/core/persistentvolumes/" + file.Name())

			m := make(map[string]interface{})
			yaml.Unmarshal([]byte(yfile), m)

			name := m["metadata"].(map[interface{}]interface{})["name"]
			nameS := fmt.Sprintf("%v", name)
			SecondList.AddItem(nameS, "", 0, nil)
			capacity := m["spec"].(map[interface{}]interface{})["capacity"].(map[interface{}]interface{})["storage"]
			capacityS := fmt.Sprintf("%v", capacity)

			access := m["spec"].(map[interface{}]interface{})["accessModes"]
			accessS := fmt.Sprintf("%v", access)
			accessS = strings.Replace(accessS, "[", "", -1)
			accessS = strings.Replace(accessS, "]", "", -1)

			reclaim := m["spec"].(map[interface{}]interface{})["claimRef"].(map[interface{}]interface{})["name"]
			reclaimS := fmt.Sprintf("%v", reclaim)

			status := m["status"].(map[interface{}]interface{})["phase"]
			statusS := fmt.Sprintf("%v", status)

			claim := m["metadata"].(map[interface{}]interface{})["name"]
			claimS := fmt.Sprintf("%v", claim)

			storageclass := m["spec"].(map[interface{}]interface{})["storageClassName"]
			storageclassS := fmt.Sprintf("%v", storageclass)

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

			Output = append(Output, nameS+"|"+capacityS+"|"+accessS+"|"+reclaimS+"|"+statusS+"|"+claimS+"|"+storageclassS+"|"+age+"\n")
		}
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput

	}

}
