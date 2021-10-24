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

var Output []string
var FormatedOutput string = ""

func ThirdListOnSelect(index int, list_item_name string, second string, run rune) {
	// FourthListItem, _ = ThirdList.GetItemText(index)
	ThirdListItem = list_item_name
	ActivePathBox.SetText(FirstListItem + " -> " + SecondListItem + " -> " + ThirdListItem)
	TextView.Clear()
	TextViewData = ""
	FourthList.Clear()
	FifthList.Clear()
	SixthList.Clear()

	FourthList.SetTitle("")
	FifthList.SetTitle("")
	SixthList.SetTitle("")
	// This section of code is dedicated for the "All Projects" choice of projects list
	if FirstListItem == "Projects" && SecondListItem == "All Projects" {
		if ThirdListItem == "Summary" {
			// Summary of all projects
			TextView.SetText("To Be Implemented")
		} else if ThirdListItem == "Pods" {
			fileInfo, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(fileInfo) > 0 {
				// Cleaning TextView and TextViewData
				TextView.Clear()
				TextViewData = ""
				// Getting current timestamp
				now := time.Now().UTC()
				Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "READY" + "|" + "STATUS" + "|" + "RESTARTS" + "|" + "Age" + "\n"}
				for projectIndex := 0; projectIndex < len(fileInfo); projectIndex++ {
					// Get project's pods "regardless of it's type/status" I can add the owner column if possible
					if _, err := os.Stat(BasePath + "namespaces/" + fileInfo[projectIndex].Name() + "/core/pods.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + fileInfo[projectIndex].Name() + "/core/pods.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							// Loop between pods
							for i := 0; i < len(x); i++ {
								pod := x[i].(map[interface{}]interface{})
								name := pod["metadata"].(map[interface{}]interface{})["name"]
								nameS := fmt.Sprintf("%v", name)
								Status := pod["status"].(map[interface{}]interface{})["phase"]
								StatusS := fmt.Sprintf("%v", Status)
								CreationTime := pod["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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
								// Initializing Ready Container count
								readyCont := 0
								// Initializing number of restarts
								restarts := 0
								containerStatuses := pod["status"].(map[interface{}]interface{})["containerStatuses"].([]interface{})
								for i := 0; i < len(containerStatuses); i++ {
									restartCount := containerStatuses[i].(map[interface{}]interface{})["restartCount"]

									if containerStatuses[i].(map[interface{}]interface{})["ready"] == true {
										readyCont++
									}
									restarts += restartCount.(int)
								}
								containers := fmt.Sprintf("%v", len(containerStatuses))
								restartsS := fmt.Sprintf("%v", restarts)
								readyContS := fmt.Sprintf("%v", readyCont)
								Output = append(Output, fileInfo[projectIndex].Name()+"|"+nameS+"|"+readyContS+"/"+containers+"|"+StatusS+"|"+restartsS+"|"+age+"\n")
							}
							FormatedOutput = columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}

					}
				}

			}
		} else if ThirdListItem == "Deployment" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/apps/deployments.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/apps/deployments.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							for i := 0; i < len(x); i++ {
								y := x[i].(map[interface{}]interface{})
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
								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+readyS+"|"+UpToDate+"|"+availableS+"|"+age+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
						}

					}
				}

			} else {
				TextView.SetText("Couldn't find namespaces list")
			}
		} else if ThirdListItem == "DeploymentConfig" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/apps.openshift.io/deploymentconfigs.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/apps.openshift.io/deploymentconfigs.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "REVISION" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "TRIGGERED BY" + "|" + "Age" + "\n"}
							for i := 0; i < len(x); i++ {
								y := x[i].(map[interface{}]interface{})
								name := y["metadata"].(map[interface{}]interface{})["name"]
								nameS := fmt.Sprintf("%v", name)

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
								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+revisionS+"|"+desiredS+"|"+currentS+"|"+triggersType+"|"+age+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			} else {
				TextView.SetText("Couldn't find namespaces list")
			}
		} else if ThirdListItem == "Daemonset" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "NODE SELECTOR" + "|" + "AGE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/apps/daemonsets.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/apps/daemonsets.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							for i := 0; i < len(x); i++ {
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

								// To be enhanced "print key/value only"
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

								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+desiredS+"|"+currentS+"|"+readyS+"|"+uptodateS+"|"+availableS+"|"+nodeselectorS+"|"+age+"|"+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if ThirdListItem == "Services" {
			// Table of Projects/Services
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "TYPE" + "|" + "CLUSTER-IP" + "|" + "EXTERNAL-IP" + "|" + "PORT(S)" + "|" + "AGE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/core/services.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/core/services.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							for i := 0; i < len(x); i++ {
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

								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+StypeS+"|"+clusterIPS+"|"+externalIPS+"|"+ports_proto+"|"+age+"|"+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if ThirdListItem == "Routes" {
			// Table of Projects/Routes
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "HOST/PORT" + "|" + "PATH" + "|" + "SERVICES" + "|" + "PORT" + "|" + "TERMINATION" + "|" + "WILDCARD" + "|" + "AGE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/route.openshift.io/routes.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/route.openshift.io/routes.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							for i := 0; i < len(x); i++ {
								y := x[i].(map[interface{}]interface{})
								name := y["metadata"].(map[interface{}]interface{})["name"]
								nameS := fmt.Sprintf("%v", name)

								host := y["spec"].(map[interface{}]interface{})["host"]
								hostS := fmt.Sprintf("%v", host)

								services := y["spec"].(map[interface{}]interface{})["to"].(map[interface{}]interface{})["name"]
								servicesS := fmt.Sprintf("%v", services)

								// port := y["spec"].(map[interface{}]interface{})["port"].(map[interface{}]interface{})["targetPort"]
								// portS := fmt.Sprintf("%v", port)
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

								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+hostS+"|"+""+"|"+servicesS+"|"+portS+"|"+term+"|"+wildcardS+"|"+age+"|"+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if ThirdListItem == "Image Stream" {
			// Table of Projects/Image Stream
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "TAGS" + "|" + "AGE 	" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/image.openshift.io/imagestreams.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/image.openshift.io/imagestreams.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							for i := 0; i < len(x); i++ {
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

								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+tagsS+"|"+age+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if ThirdListItem == "PVC" {
			// Table of Projects/PVC
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "STATUS" + "|" + "VOLUME" + "|" + "CAPACITY" + "|" + "ACCESS MODES" + "|" + "STORAGECLASS" + "|" + "AGE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							for i := 0; i < len(x); i++ {
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
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if ThirdListItem == "ConfigMap" {
			// Table of Projects/ConfigMap
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output := []string{"NAMESPACE" + "|" + "NAME" + "|" + "DATA" + "|" + "AGE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/core/configmaps.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/core/configmaps.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							for i := 0; i < len(x); i++ {
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

								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+dataS+"|"+age+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if ThirdListItem == "Secrets" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output := []string{"NAMESPACE" + "|" + "NAME" + "|" + "TYPE" + "|" + "DATA" + "|" + "AGE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/core/secrets.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/core/secrets.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal([]byte(yfile), m)
						x, _ := m["items"].([]interface{})
						if len(x) > 0 {
							for i := 0; i < len(x); i++ {
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

								Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+type_keyS+"|"+dataS+"|"+age+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if ThirdListItem == "Subscriptions" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output := []string{"NAMESPACE" + "|" + "NAME" + "|" + "PACKAGE" + "|" + "SOURCE" + "|" + "CHANNEL" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					files, _ := ioutil.ReadDir(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/operators.coreos.com/subscriptions/")

					for _, file := range files {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/operators.coreos.com/subscriptions/" + file.Name())
						m := make(map[string]interface{})
						yaml.Unmarshal([]byte(yfile), m)

						name := m["metadata"].(map[interface{}]interface{})["name"]
						nameS := fmt.Sprintf("%v", name)

						Package := m["spec"].(map[interface{}]interface{})["name"]
						packageS := fmt.Sprintf("%v", Package)

						source := m["spec"].(map[interface{}]interface{})["source"]
						sourceS := fmt.Sprintf("%v", source)

						channel := m["spec"].(map[interface{}]interface{})["channel"]
						channelS := fmt.Sprintf("%v", channel)

						Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+packageS+"|"+sourceS+"|"+channelS+"\n")
					}
					FormatedOutput := columnize.SimpleFormat(Output)
					TextView.SetText(FormatedOutput)
					TextView.ScrollToBeginning()
					TextViewData = FormatedOutput
				}
			}
		} else if ThirdListItem == "Operators" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			namespaces, _ := ioutil.ReadDir(BasePath + "namespaces/")
			if len(namespaces) > 0 {
				Output := []string{"NAMESPACE" + "|" + "NAME" + "|" + "DISPLAY" + "|" + "VERSION" + "|" + "REPLACES" + "|" + "PHASE" + "\n"}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					files, _ := ioutil.ReadDir(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/operators.coreos.com/clusterserviceversions/")
					for _, file := range files {
						yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + namespaces[projectIndex].Name() + "/operators.coreos.com/clusterserviceversions/" + file.Name())
						m := make(map[string]interface{})
						yaml.Unmarshal([]byte(yfile), m)

						name := m["metadata"].(map[interface{}]interface{})["name"]
						nameS := fmt.Sprintf("%v", name)

						// display := m["metadata"].(map[interface{}]interface{})["name"]
						// displayS := fmt.Sprintf("%v", display)
						displayS := "TBA"

						version := m["spec"].(map[interface{}]interface{})["version"]
						versionS := fmt.Sprintf("%v", version)

						// replace := m["metadata"].(map[interface{}]interface{})["name"]
						// replaceS := fmt.Sprintf("%v", replace)
						replaceS := "TBA"

						phase := m["status"].(map[interface{}]interface{})["phase"]
						phaseS := fmt.Sprintf("%v", phase)

						Output = append(Output, namespaces[projectIndex].Name()+"|"+nameS+"|"+displayS+"|"+versionS+"|"+replaceS+"|"+phaseS+"\n")

						FormatedOutput := columnize.SimpleFormat(Output)
						TextView.SetText(FormatedOutput)
						TextView.ScrollToBeginning()
						TextViewData = FormatedOutput
					}
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}
		}

		// The following "esle" statement is when we select a single project
	} else if FirstListItem == "Projects" && SecondListItem != "All Projects" {
		if ThirdListItem == "Summary" {
			TextView.SetText("To Be Implemented")
		} else if ThirdListItem == "YAML" {
			// Get project's YAML
			fileInfo, _ := os.ReadFile(BasePath + "namespaces/" + SecondListItem + "/" + SecondListItem + ".yaml")
			TextView.SetText(string(fileInfo))
			TextViewData = TextView.GetText(false)
		} else if ThirdListItem == "Events" {
			// Get project's events "if exists"
			// This should be a table
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/core/events.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				Output = []string{"Time" + "|" + "Type" + "|" + "Message" + "|" + "Reason" + "\n"}
				for i := 0; i < len(x); i++ {
					y := x[i].(map[interface{}]interface{})
					eventTime := y["eventTime"]
					eventType := y["type"]
					eventReason := y["reason"]
					eventMessage := y["message"]

					if eventTime == nil {
						eventTime = "N/A"
					} else if eventType == nil {
						eventType = "N/A"
					} else if eventMessage == nil {
						eventMessage = "N/A"
					} else if eventReason == nil {
						eventReason = "N/A"
					}
					Output = append(Output, eventTime.(string)+"|"+eventType.(string)+"|"+eventReason.(string)+"|"+eventMessage.(string)+"\n")
				}
				FormatedOutput = columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
			} else {
				TextView.SetText("No Recorded Events")
				TextViewData = TextView.GetText(false)
			}

		} else if ThirdListItem == "Pods" {
			TextView.Clear()
			TextViewData = ""
			Output = []string{"NAME" + "|" + "READY" + "|" + "STATUS" + "|" + "RESTARTS" + "|" + "AGE" + "\n"}
			// Get project's pods "regardless of it's type/status" I can add the owner column if possible
			if _, err := os.Stat(BasePath + "namespaces/" + SecondListItem + "/core/pods.yaml"); err == nil {
				yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/core/pods.yaml")
				m := make(map[interface{}]interface{})
				yaml.Unmarshal([]byte(yfile), m)
				x, _ := m["items"].([]interface{})
				now := time.Now().UTC()
				if len(x) > 0 {
					// Loop between pods
					for i := 0; i < len(x); i++ {
						pod := x[i].(map[interface{}]interface{})
						name := pod["metadata"].(map[interface{}]interface{})["name"]
						nameS := fmt.Sprintf("%v", name)
						FourthList.AddItem(nameS, "", 0, nil)
						Status := pod["status"].(map[interface{}]interface{})["phase"]
						StatusS := fmt.Sprintf("%v", Status)
						// Initializing Ready Container count
						readyCont := 0
						// Initializing number of restarts
						restarts := 0
						containerStatuses := pod["status"].(map[interface{}]interface{})["containerStatuses"].([]interface{})
						for i := 0; i < len(containerStatuses); i++ {
							restartCount := containerStatuses[i].(map[interface{}]interface{})["restartCount"]

							if containerStatuses[i].(map[interface{}]interface{})["ready"] == true {
								readyCont++
							}
							restarts += restartCount.(int)
						}
						containers := fmt.Sprintf("%v", len(containerStatuses))
						restartsS := fmt.Sprintf("%v", restarts)
						readyContS := fmt.Sprintf("%v", readyCont)

						CreationTime := pod["metadata"].(map[interface{}]interface{})["creationTimestamp"]
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
						Output = append(Output, nameS+"|"+readyContS+"/"+containers+"|"+StatusS+"|"+restartsS+"|"+age+"\n")
					}
					FormatedOutput = columnize.SimpleFormat(Output)
					TextView.SetText(FormatedOutput)
					TextView.ScrollToBeginning()
					TextViewData = FormatedOutput
				}

			} else {
				// I have to read Pods files from under [BasePath + "namespaces/" + SecondListItem + "/pods/"] and loop between them
			}

		} else if ThirdListItem == "Deployment" {
			// Get projects deployments "if exists"
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/apps/deployments.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				Output = []string{"NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE"}
				for i := 0; i < len(x); i++ {
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
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			} else {
				TextView.SetText("No Deployment resources found")
			}
		} else if ThirdListItem == "DeploymentConfig" {
			// Get projects deploymentconfigs "if exists"
			// Get projects deployments "if exists"
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/apps.openshift.io/deploymentconfigs.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				Output = []string{"NAME" + "|" + "REVISION" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "TRIGGERED BY" + "\n"}
				for i := 0; i < len(x); i++ {
					y := x[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)

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
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			} else {
				TextView.SetText("No DeploymentConfig resources found")
			}
		} else if ThirdListItem == "Daemonset" {
			// Get projects daemonsets "if exists"
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/apps/daemonsets.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				Output = []string{"NAME" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "NODE SELECTOR" + "|" + "AGE" + "\n"}
				for i := 0; i < len(x); i++ {
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
			} else {
				TextView.SetText("No Daemonset resources found")
			}
		} else if ThirdListItem == "Services" {
			// Get projects services "if exists"
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/core/services.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			Output = []string{"NAME" + "|" + "TYPE" + "|" + "CLUSTER-IP" + "|" + "EXTERNAL-IP" + "|" + "PORT(S)" + "|" + "AGE" + "\n"}
			yaml.Unmarshal([]byte(yfile), m)
			if len(x) > 0 {
				for i := 0; i < len(x); i++ {
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
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}
		} else if ThirdListItem == "Routes" {
			// Get projects routes "if exists"
			// Table of Projects/Routes
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output = []string{"NAME" + "|" + "HOST/PORT" + "|" + "PATH" + "|" + "SERVICES" + "|" + "PORT" + "|" + "TERMINATION" + "|" + "WILDCARD" + "|" + "AGE" + "\n"}
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/route.openshift.io/routes.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				for i := 0; i < len(x); i++ {
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
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if ThirdListItem == "Image Stream" {
			// Get projects IS "if exists"
			// Table of Projects/Image Stream
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "TAGS" + "|" + "AGE 	" + "\n"}
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/image.openshift.io/imagestreams.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				for i := 0; i < len(x); i++ {
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

					Output = append(Output, SecondListItem+"|"+nameS+"|"+tagsS+"|"+age+"\n")
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if ThirdListItem == "PVC" {
			// Get projects PVC "if exists"
			// Table of Projects/PVC
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output = []string{"NAME" + "|" + "STATUS" + "|" + "VOLUME 	" + "|" + "CAPACITY" + "|" + "ACCESS MODES" + "|" + "STORAGECLASS" + "|" + "AGE" + "\n"}
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/core/persistentvolumeclaims.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				for i := 0; i < len(x); i++ {
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

					Output = append(Output, nameS+"|"+statusS+"|"+volumeS+"|"+capacityS+"|"+accessS+"|"+storageCS+"|"+age+"\n")
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if ThirdListItem == "ConfigMap" {
			// Get projects CM "if exists"
			// Table of Projects/ConfigMap
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output := []string{"NAME" + "|" + "DATA" + "|" + "AGE" + "\n"}
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/core/configmaps.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				for i := 0; i < len(x); i++ {
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
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if ThirdListItem == "Secrets" {
			// Get projects secrets  "if exists"
			// Table of Projects/Secrets
			// Table of Projects/ConfigMap
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output := []string{"NAME" + "|" + "TYPE" + "|" + "DATA" + "|" + "AGE" + "\n"}
			yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/core/secrets.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal([]byte(yfile), m)
			x, _ := m["items"].([]interface{})
			if len(x) > 0 {
				for i := 0; i < len(x); i++ {
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
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if ThirdListItem == "Subscriptions" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			Output := []string{"NAME" + "|" + "PACKAGE" + "|" + "SOURCE" + "|" + "CHANNEL" + "\n"}
			files, _ := ioutil.ReadDir(BasePath + "namespaces/" + SecondListItem + "/operators.coreos.com/subscriptions/")
			for _, file := range files {
				yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/operators.coreos.com/subscriptions/" + file.Name())
				m := make(map[string]interface{})
				yaml.Unmarshal([]byte(yfile), m)

				name := m["metadata"].(map[interface{}]interface{})["name"]
				nameS := fmt.Sprintf("%v", name)

				Package := m["spec"].(map[interface{}]interface{})["name"]
				packageS := fmt.Sprintf("%v", Package)

				source := m["spec"].(map[interface{}]interface{})["source"]
				sourceS := fmt.Sprintf("%v", source)

				channel := m["spec"].(map[interface{}]interface{})["channel"]
				channelS := fmt.Sprintf("%v", channel)

				Output = append(Output, nameS+"|"+packageS+"|"+sourceS+"|"+channelS+"\n")

			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else if ThirdListItem == "Operators" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			Output := []string{"NAME" + "|" + "DISPLAY" + "|" + "VERSION" + "|" + "REPLACES" + "|" + "PHASE" + "\n"}
			files, _ := ioutil.ReadDir(BasePath + "namespaces/" + SecondListItem + "/operators.coreos.com/clusterserviceversions/")
			for _, file := range files {
				yfile, _ := ioutil.ReadFile(BasePath + "namespaces/" + SecondListItem + "/operators.coreos.com/clusterserviceversions/" + file.Name())
				m := make(map[string]interface{})
				yaml.Unmarshal([]byte(yfile), m)

				name := m["metadata"].(map[interface{}]interface{})["name"]
				nameS := fmt.Sprintf("%v", name)

				// display := m["metadata"].(map[interface{}]interface{})["name"]
				// displayS := fmt.Sprintf("%v", display)
				displayS := "TBA"

				version := m["spec"].(map[interface{}]interface{})["version"]
				versionS := fmt.Sprintf("%v", version)

				// replace := m["metadata"].(map[interface{}]interface{})["name"]
				// replaceS := fmt.Sprintf("%v", replace)
				replaceS := "TBA"

				phase := m["status"].(map[interface{}]interface{})["phase"]
				phaseS := fmt.Sprintf("%v", phase)

				Output = append(Output, nameS+"|"+displayS+"|"+versionS+"|"+replaceS+"|"+phaseS+"\n")
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		}
	} else if FirstListItem == "Nodes" && ThirdListItem == "YAML" {
		// Get node's YAML
		fileInfo, _ := os.ReadFile(BasePath + "cluster-scoped-resources/core/nodes/" + SecondListItem + ".yaml")
		TextView.SetText(string(fileInfo))
		TextViewData = TextView.GetText(false)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if FirstListItem == "Nodes" && ThirdListItem == "Summary" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""

		// Print a summerized nodes status
		now := time.Now().UTC()
		Output := []string{"NAME" + "|" + "STATUS" + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "\n"}
		yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/core/nodes/" + SecondListItem + ".yaml")
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

		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if FirstListItem == "Nodes" && ThirdListItem == "Details" {
		// TBA
	} else if FirstListItem == "Operators" && ThirdListItem == "YAML" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		operator, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/config.openshift.io/clusteroperators/" + SecondListItem + ".yaml")
		TextView.SetText(string(operator))
		TextView.ScrollToBeginning()
	} else if FirstListItem == "Operators" && ThirdListItem == "Summary" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		operatorFile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/config.openshift.io/clusteroperators/" + SecondListItem + ".yaml")

		Output := []string{"NAME" + "|" + "VERSION" + "|" + "AVAILABLE" + "|" + "PROGRESSINS" + "|" + "DEGRADED" + "|" + "SINCE" + "\n"}
		operator := make(map[interface{}]interface{})
		yaml.Unmarshal([]byte(operatorFile), operator)

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

		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if FirstListItem == "PV" && ThirdListItem == "YAML" {
		pvFile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/core/persistentvolumes/" + SecondListItem + ".yaml")
		TextView.SetText(string(pvFile))
		TextView.ScrollToBeginning()
	}
}
