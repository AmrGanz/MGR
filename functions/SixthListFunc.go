package functions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ryanuber/columnize"
	"gopkg.in/yaml.v2"
)

func SixthListOnSelect(index int, list_item_name string, second string, run rune) {
	List6Item = list_item_name
	TextView.Clear()
	ActivePathBox.SetText(List1Item + " -> " + List2Item + " -> " + List3Item + " -> " + List4Item + " -> " + List5Item + " -> " + List6Item)
	if List1Item == "Projects" && List3Item == "Pods" && List6Item == "Logs" {
		TextView.SetText("TBA")
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Deployment" && List6Item == "Info" {
		// Get projects deployments "if exists"
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps/deployments.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})

		if len(x) > 0 {
			Output = []string{"NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE"}
			for i := 0; i < len(x); i++ {
				if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
					y := x[i].(map[interface{}]interface{})
					now := time.Now().UTC()

					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)

					ready := y["status"].(map[interface{}]interface{})["readyReplicas"]
					readyS := fmt.Sprintf("%v", ready)
					UpToDate := "TBA"
					// I think I should print Ready/Avilable just like in the output of [# oc get deployment]
					available := y["status"].(map[interface{}]interface{})["availableReplicas"]
					availableS := fmt.Sprintf("%v", available)

					CreationTime := y["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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

					Output = append(Output, nameS+"|"+readyS+"|"+UpToDate+"|"+availableS+"|"+age+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else {
			TextView.SetText("No Deployment resources found")
		}
	} else if List1Item == "Projects" && List3Item == "Deployment" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps/deployments.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "DeploymentConfig" && List6Item == "Info" {
		// Get projects deploymentconfigs "if exists"
		// Get projects deployments "if exists"
		yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps.openshift.io/deploymentconfigs.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		if len(x) > 0 {
			Output = []string{"NAME" + "|" + "REVISION" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "TRIGGERED BY" + "\n"}
			for i := 0; i < len(x); i++ {
				if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
					y := x[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)
					List4.AddItem(nameS, "", 0, nil)

					revision := y["status"].(map[interface{}]interface{})["latestVersion"]
					revisionS := fmt.Sprintf("%v", revision)

					desired := y["spec"].(map[interface{}]interface{})["replicas"]
					desiredS := fmt.Sprintf("%v", desired)

					current := y["status"].(map[interface{}]interface{})["availableReplicas"]
					currentS := fmt.Sprintf("%v", current)

					triggers := y["spec"].(map[interface{}]interface{})["triggers"].([]interface{})
					triggersType := ""
					for i := 0; i < len(triggers); i++ {
						if triggers[i].(map[interface{}]interface{})["type"].(string) == "ImageChange" {
							image := triggers[i].(map[interface{}]interface{})["imageChangeParams"].(map[interface{}]interface{})["from"].(map[interface{}]interface{})["name"]
							triggersType = triggersType + "image" + "(" + image.(string) + ")"
						} else {
							triggersType = triggersType + "config"
						}
						if i != len(triggers)-1 {
							triggersType = triggersType + ","
						}
					}

					Output = append(Output, nameS+"|"+revisionS+"|"+desiredS+"|"+currentS+"|"+triggersType+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else {
			TextView.SetText("No DeploymentConfig resources found")
		}
	} else if List1Item == "Projects" && List3Item == "DeploymentConfig" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps.openshift.io/deploymentconfigs.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Daemonset" && List6Item == "Info" {
		// Get projects daemonsets "if exists"
		yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps/daemonsets.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		if len(x) > 0 {
			Output = []string{"NAME" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "NODE SELECTOR" + "|" + "AGE" + "\n"}
			for i := 0; i < len(x); i++ {
				if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
					now := time.Now().UTC()
					y := x[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)

					desired := y["status"].(map[interface{}]interface{})["desiredNumberScheduled"]
					desiredS := fmt.Sprintf("%v", desired)

					current := y["status"].(map[interface{}]interface{})["currentNumberScheduled"]
					currentS := fmt.Sprintf("%v", current)

					ready := y["status"].(map[interface{}]interface{})["numberReady"]
					readyS := fmt.Sprintf("%v", ready)

					uptodate := y["status"].(map[interface{}]interface{})["updatedNumberScheduled"]
					uptodateS := fmt.Sprintf("%v", uptodate)

					available := y["status"].(map[interface{}]interface{})["numberAvailable"]
					availableS := fmt.Sprintf("%v", available)

					// To be enhanced
					nodeselector := y["spec"].(map[interface{}]interface{})["template"].(map[interface{}]interface{})["spec"].(map[interface{}]interface{})["nodeSelector"]
					nodeselectorS := fmt.Sprintf("%v", nodeselector)

					CreationTime := y["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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

					Output = append(Output, nameS+"|"+desiredS+"|"+currentS+"|"+readyS+"|"+uptodateS+"|"+availableS+"|"+nodeselectorS+"|"+age+"|"+"\n")
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}
		} else {
			TextView.SetText("No Daemonset resources found")
		}
	} else if List1Item == "Projects" && List3Item == "Daemonset" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps/daemonsets.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Services" && List6Item == "Info" {
		// Get projects services "if exists"
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		// Getting current timestamp
		now := time.Now().UTC()
		yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/services.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		Output = []string{"NAME" + "|" + "TYPE" + "|" + "CLUSTER-IP" + "|" + "EXTERNAL-IP" + "|" + "PORT(S)" + "|" + "AGE" + "\n"}
		yaml.Unmarshal(yfile, m)
		if len(x) > 0 {
			for i := 0; i < len(x); i++ {
				if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
					y := x[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)

					Stype := y["spec"].(map[interface{}]interface{})["type"]
					StypeS := fmt.Sprintf("%v", Stype)

					clusterIP := y["spec"].(map[interface{}]interface{})["clusterIP"]
					clusterIPS := ""
					if clusterIP == nil {
						clusterIPS = "None"
					} else {
						clusterIPS = fmt.Sprintf("%v", clusterIP)
					}

					externalIP := y["spec"].(map[interface{}]interface{})["externalName"]
					externalIPS := ""
					if externalIP == nil {
						externalIPS = "None"
					} else {
						externalIPS = fmt.Sprintf("%v", externalIP)
					}

					CreationTime := y["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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
					ports_proto := ""
					if y["spec"].(map[interface{}]interface{})["ports"] != nil {
						ports := y["spec"].(map[interface{}]interface{})["ports"].([]interface{})
						for i := 0; i < len(ports); i++ {
							port := ports[i].(map[interface{}]interface{})["port"]
							portS := fmt.Sprintf("%v", port)
							proto := ports[i].(map[interface{}]interface{})["protocol"]
							protoS := fmt.Sprintf("%v", proto)
							ports_proto = ports_proto + portS + "/" + protoS
							if i != len(ports)-1 {
								ports_proto = ports_proto + ","
							}
						}
					} else {
						ports_proto = "None"
					}
					Output = append(Output, nameS+"|"+StypeS+"|"+clusterIPS+"|"+externalIPS+"|"+ports_proto+"|"+age+"|"+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		}
	} else if List1Item == "Projects" && List3Item == "Services" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/services.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Routes" && List6Item == "Info" {
		// Get projects routes "if exists"
		// Table of Projects/Routes
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		// Getting current timestamp
		now := time.Now().UTC()
		Output = []string{"NAME" + "|" + "HOST/PORT" + "|" + "PATH" + "|" + "SERVICES" + "|" + "PORT" + "|" + "TERMINATION" + "|" + "WILDCARD" + "|" + "AGE" + "\n"}
		yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/route.openshift.io/routes.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		if len(x) > 0 {
			for i := 0; i < len(x); i++ {
				if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
					y := x[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)

					host := y["spec"].(map[interface{}]interface{})["host"]
					hostS := fmt.Sprintf("%v", host)

					services := y["spec"].(map[interface{}]interface{})["to"].(map[interface{}]interface{})["name"]
					servicesS := fmt.Sprintf("%v", services)

					portS := ""
					if y["spec"].(map[interface{}]interface{})["port"] != nil {
						port := y["spec"].(map[interface{}]interface{})["port"].(map[interface{}]interface{})["targetPort"]
						portS = fmt.Sprintf("%v", port)
					} else {
						portS = "nil"
					}

					insecureEdgeTerminationPolicy := y["spec"].(map[interface{}]interface{})["tls"].(map[interface{}]interface{})["insecureEdgeTerminationPolicy"]
					termination := y["spec"].(map[interface{}]interface{})["tls"].(map[interface{}]interface{})["termination"]
					term := ""
					if termination == nil && insecureEdgeTerminationPolicy == nil {
						term = "None"
					} else {
						term = fmt.Sprintf("%v", termination) + "/" + fmt.Sprintf("%v", insecureEdgeTerminationPolicy)
					}
					wildcard := y["spec"].(map[interface{}]interface{})["wildcardPolicy"]
					wildcardS := fmt.Sprintf("%v", wildcard)
					CreationTime := y["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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

					Output = append(Output, nameS+"|"+hostS+"|"+""+"|"+servicesS+"|"+portS+"|"+term+"|"+wildcardS+"|"+age+"|"+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput

		}
	} else if List1Item == "Projects" && List3Item == "Routes" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/route.openshift.io/routes.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Image Stream" && List6Item == "Info" {
		// Get projects IS "if exists"
		// Table of Projects/Image Stream
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		// Getting current timestamp
		now := time.Now().UTC()
		Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "TAGS" + "|" + "AGE 	" + "\n"}
		yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/image.openshift.io/imagestreams.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		if len(x) > 0 {
			for i := 0; i < len(x); i++ {
				if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
					y := x[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)

					tagsS := ""
					if y["spec"].(map[interface{}]interface{})["tags"] != nil {
						all_tags := y["spec"].(map[interface{}]interface{})["tags"].([]interface{})
						for i := 0; i < len(all_tags); i++ {
							tag_name := all_tags[i].(map[interface{}]interface{})["name"]
							tag_nameS := fmt.Sprintf("%v", tag_name)
							tagsS = tagsS + tag_nameS
							if i != len(all_tags)-1 {
								tagsS = tagsS + ","
							}
						}
					}

					CreationTime := y["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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

					Output = append(Output, List2Item+"|"+nameS+"|"+tagsS+"|"+age+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		}
	} else if List1Item == "Projects" && List3Item == "Image Stream" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/image.openshift.io/imagestreams.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "PVC" && List6Item == "Info" {
		// Table of Projects/PVC
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		// Getting current timestamp
		now := time.Now().UTC()
		namespaces, _ := ioutil.ReadDir(MG_Path + "namespaces/")
		if len(namespaces) > 0 {
			Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "STATUS" + "|" + "VOLUME" + "|" + "CAPACITY" + "|" + "ACCESS MODES" + "|" + "STORAGECLASS" + "|" + "AGE" + "\n"}
			for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
				if _, err := os.Stat(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml"); err == nil {
					yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml")
					m := make(map[interface{}]interface{})
					yaml.Unmarshal(yfile, m)
					x, _ := m["items"].([]interface{})
					if len(x) > 0 {
						for i := 0; i < len(x); i++ {
							if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
								y := x[i].(map[interface{}]interface{})
								name := y["metadata"].(map[interface{}]interface{})["name"]
								nameS := fmt.Sprintf("%v", name)

								status := y["status"].(map[interface{}]interface{})["phase"]
								statusS := fmt.Sprintf("%v", status)

								volume := y["spec"].(map[interface{}]interface{})["volumeName"]
								volumeS := fmt.Sprintf("%v", volume)

								capacity := y["status"].(map[interface{}]interface{})["capacity"].(map[interface{}]interface{})["storage"]
								capacityS := fmt.Sprintf("%v", capacity)

								access := y["status"].(map[interface{}]interface{})["accessModes"]
								accessS := fmt.Sprintf("%v", access)
								accessS = strings.Replace(accessS, "[", "", -1)
								accessS = strings.Replace(accessS, "]", "", -1)

								storageC := y["spec"].(map[interface{}]interface{})["storageClassName"]
								storageCS := fmt.Sprintf("%v", storageC)

								CreationTime := y["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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
								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+statusS+"|"+volumeS+"|"+capacityS+"|"+accessS+"|"+storageCS+"|"+age+"\n")
							}
						}
						FormatedOutput := columnize.SimpleFormat(Output)
						TextView.SetText(FormatedOutput)
						TextView.ScrollToBeginning()
						TextViewData = FormatedOutput
					}
				}
			}
		}
	} else if List1Item == "Projects" && List3Item == "PVC" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/persistentvolumeclaims.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "ConfigMap" && List6Item == "Info" {
		// Get projects CM "if exists"
		// Table of Projects/ConfigMap
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		// Getting current timestamp
		now := time.Now().UTC()
		Output := []string{"NAME" + "|" + "DATA" + "|" + "AGE" + "\n"}
		yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/configmaps.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		if len(x) > 0 {
			for i := 0; i < len(x); i++ {
				if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
					y := x[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)

					dataS := ""
					if y["data"] != nil {
						data := y["data"].(map[interface{}]interface{})
						dataN := len(data)
						dataS = fmt.Sprint(dataN)
					} else {
						dataS = "0"
					}

					CreationTime := y["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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
					Output = append(Output, nameS+"|"+dataS+"|"+age+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput

		}
	} else if List1Item == "Projects" && List3Item == "ConfigMap" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/configmaps.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Secrets" && List6Item == "Info" {
		// Get projects secrets  "if exists"
		// Table of Projects/Secrets
		// Table of Projects/ConfigMap
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		// Getting current timestamp
		now := time.Now().UTC()
		Output := []string{"NAME" + "|" + "TYPE" + "|" + "DATA" + "|" + "AGE" + "\n"}
		yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/secrets.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		if len(x) > 0 {
			for i := 0; i < len(x); i++ {
				if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
					y := x[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)

					dataS := ""
					if y["data"] != nil {
						data := y["data"].(map[interface{}]interface{})
						dataN := len(data)
						dataS = fmt.Sprint(dataN)
					} else {
						dataS = "0"
					}

					type_key := y["type"]
					type_keyS := fmt.Sprintf("%v", type_key)

					CreationTime := y["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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

					Output = append(Output, nameS+"|"+type_keyS+"|"+dataS+"|"+age+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput

		}
	} else if List1Item == "Projects" && List3Item == "Secrets" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/secrets.yaml")
		m := make(map[interface{}]interface{})
		yaml.Unmarshal(yfile, m)
		x, _ := m["items"].([]interface{})
		for i := range x {
			if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
				yaml, _ := yaml.Marshal(x[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Subscriptions" && List6Item == "Info" {
		//TBA
	} else if List1Item == "Projects" && List3Item == "Supscriptions" && List6Item == "YAML" {
		// yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps/deployments.yaml")
		// m := make(map[interface{}]interface{})
		// yaml.Unmarshal(yfile, m)
		// x, _ := m["items"].([]interface{})
		// for i := range x {
		// 	if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
		// 		yaml, _ := yaml.Marshal(x[i])
		// 		TextView.SetText(fmt.Sprintf(string(yaml)))
		// 	}
		// }
		TextView.SetText("TBA")
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Operators" && List6Item == "Info" {

	} else if List1Item == "Projects" && List3Item == "Operators" && List6Item == "YAML" {
		// yfile, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps/deployments.yaml")
		// m := make(map[interface{}]interface{})
		// yaml.Unmarshal(yfile, m)
		// x, _ := m["items"].([]interface{})
		// for i := range x {
		// 	if x[i].(map[interface{}]interface{})["metadata"].(map[interface{}]interface{})["name"] == List4Item {
		// 		yaml, _ := yaml.Marshal(x[i])
		// 		TextView.SetText(fmt.Sprintf(string(yaml)))
		// 	}
		// }
		TextView.SetText("TBA")
		TextView.ScrollToBeginning()
	} else if List1Item == "Nodes" && List3Item == "YAML" {
		if List6Item == "Metadata" {
			Metadta, _ := yaml.Marshal(MyNode_Public.Metadata)
			MetadtaS := Colors.Orange + "Metadta:\n" + Colors.White + string(Metadta)
			TextView.Clear()
			TextView.SetText(MetadtaS)
			TextView.ScrollToBeginning()
		} else if List6Item == "Spec" {
			Spec, _ := yaml.Marshal(MyNode_Public.Spec)
			SpecS := Colors.Orange + "Spec:\n" + Colors.White + string(Spec)
			TextView.Clear()
			TextView.SetText(SpecS)
			TextView.ScrollToBeginning()
		} else if List6Item == "Status" {
			Status, _ := yaml.Marshal(MyNode_Public.Status)
			StatusS := Colors.Orange + "Status:\n" + Colors.White + string(Status)
			TextView.Clear()
			TextView.SetText(StatusS)
			TextView.ScrollToBeginning()
		} else if List6Item == "HW Spec" {
			HWSpec := ""
			Addresses, _ := yaml.Marshal(MyNode_Public.Status.Addresses)
			Allocatable, _ := yaml.Marshal(MyNode_Public.Status.Allocatable)
			Capacity, _ := yaml.Marshal(MyNode_Public.Status.Capacity)
			HWSpec += Colors.Orange + "Addresses:\n" + Colors.White + string(Addresses) + Colors.Orange + "Allocatable:\n" + Colors.White + string(Allocatable) + Colors.Orange + "Capacity:\n" + Colors.White + string(Capacity)
			TextView.Clear()
			TextView.SetText(HWSpec)
			TextView.ScrollToBeginning()
		} else if List6Item == "Images" {
			Images, _ := yaml.Marshal(MyNode_Public.Status.Images)
			ImagesS := Colors.Orange + "Images:\n" + Colors.White + string(Images)
			TextView.Clear()
			TextView.SetText(ImagesS)
			TextView.ScrollToBeginning()
		}
	}
}
