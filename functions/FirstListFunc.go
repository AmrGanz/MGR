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
	List2.Clear()
	List3.Clear()
	List4.Clear()
	List5.Clear()
	List6.Clear()
	List2.SetTitle("")
	List3.SetTitle("")
	List4.SetTitle("")
	List5.SetTitle("")
	List6.SetTitle("")

	// Setting the global variable List1Item
	List1Item = list_item_name

	// Adding selection path to the ActivePathBox
	ActivePathBox.SetText(List1Item + " -> ")
	if list_item_name == "Configurations" {
		List2.SetTitle("Cluster Configurations")
		files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/config.openshift.io/")
		for i := range files {
			if !files[i].IsDir() {
				List2.AddItem(strings.Split(files[i].Name(), ".yaml")[0], "", 0, nil)
			}
		}

	} else if list_item_name == "Projects" {
		List2.SetTitle("Projects")
		namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
		if len(namespaces) > 0 {
			List2.AddItem("All Projects", "", 0, nil)
			for _, project := range namespaces {
				List2.AddItem(project.Name(), "", 0, nil)
			}
		} else {
			List2.AddItem("Empty", "", 0, nil)
		}
	} else if list_item_name == "Nodes" {
		List2.SetTitle("Nodes")
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/core/nodes/")
		List2.
			AddItem("Nodes Summary", "", 0, nil).
			AddItem("Nodes Details", "", 0, nil)
		// Populate List2 with nodes names
		for _, node := range files {
			nodeName := strings.Split(node.Name(), ".yaml")
			List2.AddItem(nodeName[0], "", 0, nil)
		}

	} else if list_item_name == "Operators" {
		List2.SetTitle("Operators")
		clusteroperators, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/config.openshift.io/clusteroperators.yaml")
		Output := []string{Colors.Yellow + "NAME" + "|" + "VERSION" + "|" + "AVAILABLE" + "|" + "PROGRESSING" + "|" + "DEGRADED" + "|" + "SINCE" + Colors.White}

		m := make(map[interface{}]interface{})
		yaml.Unmarshal(clusteroperators, m)
		items, _ := m["items"].([]interface{})
		for i := range items {
			operator := items[i].(map[interface{}]interface{})
			name := operator["metadata"].(map[interface{}]interface{})["name"]
			nameS := fmt.Sprintf("%v", name)
			List2.AddItem(nameS, "", 0, nil)

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
			Output = append(Output, Colors.White+nameS+"|"+versionS+"|"+availableS+"|"+progressingS+"|"+degradedS+"|"+availableSince+Colors.White)
		}
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if list_item_name == "MCP" {
		List2.SetTitle("MCP")
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		Output := []string{Colors.Yellow + "NAME" + "|" + "CONFIG" + "|" + "UPDATED" + "|" + "UPDATING" + "|" + "DEGRADED" + "|" + "MACHINECOUNT" + "|" + "READYMACHINECOUNT" + "|" + "UPDATEDMACHINECOUNT" + "|" + "DEGRADEDMACHINECOUNT" + "|" + "AGE" + Colors.White}

		files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigpools/")
		for _, mcp := range files {
			yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigpools/" + mcp.Name())

			m := make(map[string]interface{})
			yaml.Unmarshal(yfile, m)

			name := m["metadata"].(map[interface{}]interface{})["name"]
			nameS := fmt.Sprintf("%v", name)
			List2.AddItem(nameS, "", 0, nil)

			config := m["status"].(map[interface{}]interface{})["configuration"].(map[interface{}]interface{})["name"]
			configS := fmt.Sprintf("%v", config)

			status := m["status"].(map[interface{}]interface{})["conditions"].([]interface{})
			updatedS := ""
			updatingS := ""
			degradedS := ""
			for i := range status {
				if status[i].(map[interface{}]interface{})["type"] == "Updated" {
					if status[i].(map[interface{}]interface{})["status"] == "True" {
						updatedS = "True"
					} else {
						updatedS = "False"
					}

				} else if status[i].(map[interface{}]interface{})["type"] == "Updating" {
					if status[i].(map[interface{}]interface{})["status"] == "True" {
						updatingS = "True"
					} else {
						updatingS = "False"
					}
				} else if status[i].(map[interface{}]interface{})["type"] == "Degraded" {
					if status[i].(map[interface{}]interface{})["status"] == "True" {
						degradedS = "True"
					} else {
						degradedS = "False"
					}
				}

			}
			machineCount := m["status"].(map[interface{}]interface{})["machineCount"]
			machineCountS := fmt.Sprintf("%v", machineCount)

			machineReady := m["status"].(map[interface{}]interface{})["readyMachineCount"]
			machineReadyS := fmt.Sprintf("%v", machineReady)

			machineUpdated := m["status"].(map[interface{}]interface{})["updatedMachineCount"]
			machineUpdatedS := fmt.Sprintf("%v", machineUpdated)

			machineDegraded := m["status"].(map[interface{}]interface{})["degradedMachineCount"]
			machineDegradedS := fmt.Sprintf("%v", machineDegraded)

			now := time.Now().UTC()
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
			Output = append(Output, Colors.White+nameS+"|"+configS+"|"+updatedS+"|"+updatingS+"|"+degradedS+"|"+machineCountS+"|"+machineReadyS+"|"+machineUpdatedS+"|"+machineDegradedS+"|"+age+Colors.White)
		}
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput

	} else if list_item_name == "MC" {
		List2.SetTitle("MC")
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		now := time.Now().UTC()
		Output := []string{Colors.Yellow + "NAME" + "|" + "GENERATEDBYCONTROLLER" + "|" + "IGNITIONVERSION" + "|" + "AGE" + Colors.White + "\n"}
		files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigs/")
		for _, file := range files {
			yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigs/" + file.Name())

			m := make(map[string]interface{})
			yaml.Unmarshal(yfile, m)

			name := m["metadata"].(map[interface{}]interface{})["name"]
			nameS := fmt.Sprintf("%v", name)
			List2.AddItem(nameS, "", 0, nil)

			// TBA
			// ganaratedBy := m["metadata"].(map[interface{}]interface{})["annotations"].(map[interface{}]interface{})["machineconfiguration.openshift.io/generated-by-controller-version"]
			// generatedByS := fmt.Sprintf("%v", ganaratedBy)
			generatedByS := "TBA"

			ignitionVersion := m["spec"].(map[interface{}]interface{})["config"].(map[interface{}]interface{})["ignition"].(map[interface{}]interface{})["version"]
			ignitionVersionS := fmt.Sprintf("%v", ignitionVersion)

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
			Output = append(Output, Colors.White+nameS+"|"+generatedByS+"|"+ignitionVersionS+"|"+age+Colors.White+"\n")
		}
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if list_item_name == "PV" {
		List2.SetTitle("PV")
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		now := time.Now().UTC()
		Output := []string{Colors.Yellow + "NAME" + "|" + "CAPACITY" + "|" + "ACCESS MODE" + "|" + "RECLAIM POLICY" + "|" + "STATUS" + "|" + "CLAIM" + "|" + "STORAGECLASS" + "|" + "AGE" + Colors.White}
		files, _ := ioutil.ReadDir(BasePath + "cluster-scoped-resources/core/persistentvolumes/")
		for _, file := range files {
			yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/core/persistentvolumes/" + file.Name())

			m := make(map[string]interface{})
			yaml.Unmarshal(yfile, m)

			name := m["metadata"].(map[interface{}]interface{})["name"]
			nameS := fmt.Sprintf("%v", name)
			List2.AddItem(nameS, "", 0, nil)
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

			Output = append(Output, Colors.White+nameS+"|"+capacityS+"|"+accessS+"|"+reclaimS+"|"+statusS+"|"+claimS+"|"+storageclassS+"|"+age+Colors.White)
		}
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput

	} else if list_item_name == "CSR" {
		Files, _ = ioutil.ReadDir(BasePath + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/")
		List2.SetTitle("CSR")
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		List2.AddItem("All Certificate Signing Requests", "", 0, nil)
		for _, File := range Files {
			List2.AddItem(strings.Split(File.Name(), ".yaml")[0], "", 0, nil)
		}
	}
}
