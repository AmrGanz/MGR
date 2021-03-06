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

var FormatedOutput string = ""

func ThirdListOnSelect(index int, list_item_name string, second string, run rune) {
	// List4Item, _ = ThirdList.GetItemText(index)
	List3Item = list_item_name
	ActivePathBox.SetText(List1Item + " -> " + List2Item + " -> " + List3Item)
	TextView.Clear()
	TextViewData = ""
	List4.Clear()
	List5.Clear()
	List6.Clear()

	List4.SetTitle("")
	List5.SetTitle("")
	List6.SetTitle("")
	// This section of code is dedicated for the "All Projects" choice of projects list
	if List1Item == "Projects" && List2Item == "All Projects" {
		if List3Item == "Summary" {
			// Summary of all projects
			TextView.SetText("To Be Implemented")
		} else if List3Item == "Pods" {
			Files, _ = ioutil.ReadDir(MG_Path + "namespaces/")
			if len(Files) > 0 {
				// Cleaning TextView and TextViewData
				TextView.Clear()
				TextViewData = ""
				// Getting current timestamp
				Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + Colors.Yellow + "READY" + Colors.Yellow + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + Colors.Yellow + "RESTARTS" + Colors.Yellow + "|" + "Age" + Colors.White}
				for projectIndex := 0; projectIndex < len(Files); projectIndex++ {
					// Get project's pods "regardless of it's type/status" I can add the owner column if possible
					if _, err := os.Stat(MG_Path + "namespaces/" + Files[projectIndex].Name() + "/core/pods.yaml"); err == nil {
						File, _ = ioutil.ReadFile(MG_Path + "namespaces/" + Files[projectIndex].Name() + "/core/pods.yaml")
						MyPods := PODS{}
						yaml.Unmarshal(File, &MyPods)
						items := MyPods.Items
						if len(items) > 0 {
							// Loop between pods
							for i := 0; i < len(items); i++ {
								name := items[i].Metadata.Name
								Status := items[i].Status.Phase
								if Status == "Running" || Status == "Succeeded" || Status == "Completed" {
									Status = Colors.White + Status + Colors.White
								} else {
									Status = Colors.Red + Status + Colors.White
								}

								CreationTime := items[i].Metadata.CreationTimestamp
								age := GetAge(CreationTime)

								// Initializing Ready Container count
								readyCount := 0
								// Initializing number of restarts
								restarts := 0
								containerStatuses := items[i].Status.ContainerStatuses
								for x := 0; x < len(containerStatuses); x++ {
									restartCount := containerStatuses[x].RestartCount

									if containerStatuses[x].Ready { // equal to containerStatuses[x].Ready == true
										readyCount++
									}
									restarts += restartCount
								}
								restartsS := ""
								if restarts > 0 {
									restartsS = Colors.Orange + fmt.Sprintf("%v", restarts) + Colors.White
								} else {
									restartsS = Colors.White + fmt.Sprintf("%v", restarts) + Colors.White
								}
								containers := len(items[i].Spec.Containers)
								containersS := fmt.Sprintf("%v", containers)
								readyConutS := fmt.Sprintf("%v", readyCount)
								READY := ""
								if readyCount != containers {
									READY = Colors.Orange + readyConutS + "/" + containersS + Colors.White
								} else {
									READY = Colors.White + readyConutS + "/" + containersS + Colors.White
								}
								Output = append(Output, Colors.White+Files[projectIndex].Name()+"|"+name+"|"+READY+"|"+Status+"|"+restartsS+"|"+age+Colors.White)
							}
							FormatedOutput = columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}

					}
				}

			}
		} else if List3Item == "Deployment" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/apps/deployments.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/apps/deployments.yaml")
						MyDeployment := DEPLOYMENT{}
						yaml.Unmarshal(File, &MyDeployment)

						items := MyDeployment.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								Deployment := MyDeployment.Items[i]
								name := Deployment.Metadata.Name

								ready := strconv.Itoa(Deployment.Status.ReadyReplicas)

								UpToDate := "TBA"

								// I think I should print Ready/Avilable just like in the output of [# oc get deployment]
								available := strconv.Itoa(Deployment.Status.AvailableReplicas)

								CreationTime := Deployment.Metadata.CreationTimestamp
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
								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+ready+"|"+UpToDate+"|"+available+"|"+age+Colors.White)
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
		} else if List3Item == "DeploymentConfig" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/apps.openshift.io/deploymentconfigs.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/apps.openshift.io/deploymentconfigs.yaml")

						MyDeploymentconfig := DEPLOYMENTCONFIG{}
						yaml.Unmarshal(File, &MyDeploymentconfig)

						items := MyDeploymentconfig.Items
						if len(items) > 0 {
							Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "REVISION" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "TRIGGERED BY" + "|" + "Age" + Colors.White}
							for i := 0; i < len(items); i++ {
								DC := MyDeploymentconfig.Items[i]
								name := DC.Metadata.Name

								revision := strconv.Itoa(DC.Status.LatestVersion)

								desired := strconv.Itoa(DC.Spec.Replicas)

								current := strconv.Itoa(DC.Status.AvailableReplicas)

								triggers := DC.Spec.Triggers
								triggersType := ""
								for i := 0; i < len(triggers); i++ {
									if triggers[i].Type == "ImageChange" {
										image := triggers[i].ImageChangeParams.From.Name
										triggersType = triggersType + "image" + "(" + image + ")"
									} else {
										triggersType = triggersType + "config"
									}
									if i != len(triggers)-1 {
										triggersType = triggersType + ","
									}
								}
								CreationTime := DC.Metadata.CreationTimestamp
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
								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+revision+"|"+desired+"|"+current+"|"+triggersType+"|"+age+Colors.White)
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
		} else if List3Item == "Daemonset" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "NODE SELECTOR" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/apps/daemonsets.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/apps/daemonsets.yaml")
						MyDaemonset := DAEMONSET{}
						yaml.Unmarshal(File, &MyDaemonset)
						items := MyDaemonset.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								Daemonset := MyDaemonset.Items[i]
								name := Daemonset.Metadata.Name

								desired := strconv.Itoa(Daemonset.Status.DesiredNumberScheduled)

								current := strconv.Itoa(Daemonset.Status.CurrentNumberScheduled)

								ready := strconv.Itoa(Daemonset.Status.NumberReady)

								uptodate := strconv.Itoa(Daemonset.Status.UpdatedNumberScheduled)

								available := strconv.Itoa(Daemonset.Status.NumberAvailable)

								// To be enhanced "print key/value only"
								nodeselector := Daemonset.Spec.Template.Spec.NodeSelector
								nodeselectorS := ""
								count := len(nodeselector)
								for key, value := range nodeselector {
									if count > 1 {
										if fmt.Sprintf("%v", value) != "" {
											nodeselectorS = nodeselectorS + fmt.Sprintf("%v", key) + ":" + fmt.Sprintf("%v", value) + ","
										} else {
											nodeselectorS = nodeselectorS + fmt.Sprintf("%v", key) + ","
										}
										count--
									} else {
										if fmt.Sprintf("%v", value) != "" {
											nodeselectorS = nodeselectorS + fmt.Sprintf("%v", key) + ":" + fmt.Sprintf("%v", value)
										} else {
											nodeselectorS = nodeselectorS + fmt.Sprintf("%v", key)
										}
									}

								}

								CreationTime := Daemonset.Metadata.CreationTimestamp
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

								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+desired+"|"+current+"|"+ready+"|"+uptodate+"|"+available+"|"+nodeselectorS+"|"+age+"|"+Colors.White)
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if List3Item == "Services" {
			// Table of Projects/Services
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "TYPE" + "|" + "CLUSTER-IP" + "|" + "EXTERNAL-IP" + "|" + "PORT(S)" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/core/services.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/core/services.yaml")
						MyServices := SERVICES{}
						yaml.Unmarshal(File, &MyServices)
						items := MyServices.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								Service := MyServices.Items[i]
								name := Service.Metadata.Name

								Stype := Service.Spec.Type

								clusterIP := Service.Spec.ClusterIP

								externalIP := "TBA"

								CreationTime := Service.Metadata.CreationTimestamp
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
								if Service.Spec.Ports != nil {
									ports := Service.Spec.Ports
									for i := 0; i < len(ports); i++ {
										port := ports[i].Port
										portS := fmt.Sprintf("%v", port)
										proto := ports[i].Protocol
										protoS := fmt.Sprintf("%v", proto)
										ports_proto = ports_proto + portS + "/" + protoS
										if i != len(ports)-1 {
											ports_proto = ports_proto + ","
										}
									}
								} else {
									ports_proto = "None"
								}

								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+Stype+"|"+clusterIP+"|"+externalIP+"|"+ports_proto+"|"+age+"|"+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if List3Item == "Routes" {
			// Table of Projects/Routes
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "HOST/PORT" + "|" + "PATH" + "|" + "SERVICES" + "|" + "PORT" + "|" + "TERMINATION" + "|" + "WILDCARD" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/route.openshift.io/routes.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/route.openshift.io/routes.yaml")
						MyRoutes := ROUTES{}
						yaml.Unmarshal(File, &MyRoutes)
						items := MyRoutes.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								name, host, services, port, term, wildcard, age := GetRouteDetails(items[i])
								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+host+"|"+"TBA"+"|"+services+"|"+port+"|"+term+"|"+wildcard+"|"+age+"|"+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if List3Item == "Image Stream" {
			// Table of Projects/Image Stream
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(MG_Path + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "TAGS" + "|" + "AGE 	" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/image.openshift.io/imagestreams.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/image.openshift.io/imagestreams.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal(yfile, m)
						items, _ := m["items"].([]interface{})
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								y := items[i].(map[interface{}]interface{})
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

								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+nameS+"|"+tagsS+"|"+age+Colors.White)
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if List3Item == "PVC" {
			// Table of Projects/PVC
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(MG_Path + "namespaces/")
			if len(namespaces) > 0 {
				Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "STATUS" + "|" + "VOLUME" + "|" + "CAPACITY" + "|" + "ACCESS MODES" + "|" + "STORAGECLASS" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal(yfile, m)
						items, _ := m["items"].([]interface{})
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								y := items[i].(map[interface{}]interface{})
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
								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+nameS+"|"+statusS+"|"+volumeS+"|"+capacityS+"|"+accessS+"|"+storageCS+"|"+age+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if List3Item == "ConfigMap" {
			// Table of Projects/ConfigMap
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(MG_Path + "namespaces/")
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "DATA" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/core/configmaps.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/core/configmaps.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal(yfile, m)
						items, _ := m["items"].([]interface{})
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								y := items[i].(map[interface{}]interface{})
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

								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+nameS+"|"+dataS+"|"+age+Colors.White)
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if List3Item == "Secrets" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			namespaces, _ := ioutil.ReadDir(MG_Path + "namespaces/")
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "TYPE" + "|" + "DATA" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/core/secrets.yaml"); err == nil {
						yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/core/secrets.yaml")
						m := make(map[interface{}]interface{})
						yaml.Unmarshal(yfile, m)
						items, _ := m["items"].([]interface{})
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								y := items[i].(map[interface{}]interface{})
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

								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+nameS+"|"+type_keyS+"|"+dataS+"|"+age+"\n")
							}
							FormatedOutput := columnize.SimpleFormat(Output)
							TextView.SetText(FormatedOutput)
							TextView.ScrollToBeginning()
							TextViewData = FormatedOutput
						}
					}

				}
			}
		} else if List3Item == "Subscriptions" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			namespaces, _ := ioutil.ReadDir(MG_Path + "namespaces/")
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "PACKAGE" + "|" + "SOURCE" + "|" + "CHANNEL" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					files, _ := ioutil.ReadDir(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/operators.coreos.com/subscriptions/")

					for _, file := range files {
						yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/operators.coreos.com/subscriptions/" + file.Name())
						m := make(map[string]interface{})
						yaml.Unmarshal(yfile, m)

						name := m["metadata"].(map[interface{}]interface{})["name"]
						nameS := fmt.Sprintf("%v", name)

						Package := m["spec"].(map[interface{}]interface{})["name"]
						packageS := fmt.Sprintf("%v", Package)

						source := m["spec"].(map[interface{}]interface{})["source"]
						sourceS := fmt.Sprintf("%v", source)

						channel := m["spec"].(map[interface{}]interface{})["channel"]
						channelS := fmt.Sprintf("%v", channel)

						Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+nameS+"|"+packageS+"|"+sourceS+"|"+channelS+Colors.White)
					}
					FormatedOutput := columnize.SimpleFormat(Output)
					TextView.SetText(FormatedOutput)
					TextView.ScrollToBeginning()
					TextViewData = FormatedOutput
				}
			}
		} else if List3Item == "Operators" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			namespaces, _ := ioutil.ReadDir(MG_Path + "namespaces/")
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "DISPLAY" + "|" + "VERSION" + "|" + "REPLACES" + "|" + "PHASE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					files, _ := ioutil.ReadDir(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/operators.coreos.com/clusterserviceversions/")
					for _, file := range files {
						yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + namespaces[projectIndex].Name() + "/operators.coreos.com/clusterserviceversions/" + file.Name())
						m := make(map[string]interface{})
						yaml.Unmarshal(yfile, m)

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

						Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+nameS+"|"+displayS+"|"+versionS+"|"+replaceS+"|"+phaseS+"\n")

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
		} else if List3Item == "Install Plans" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "CSV" + "|" + "APPROVAL" + "|" + "APPROVED" + Colors.White}
			Namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			for _, Namespace := range Namespaces {
				Files, _ = ioutil.ReadDir(Namespaces_Path + Namespace.Name() + "/operators.coreos.com/installplans/")
				for i := range Files {
					File, _ = ioutil.ReadFile(Namespaces_Path + Namespace.Name() + "/operators.coreos.com/installplans/" + Files[i].Name())
					MyIP := INSTALLPLAN{}
					yaml.Unmarshal(File, &MyIP)
					name := MyIP.Metadata.Name
					csv := MyIP.Spec.ClusterServiceVersionNames[0]
					approval := MyIP.Spec.Approval
					approved := MyIP.Spec.Approved
					Output = append(Output, Colors.White+Namespace.Name()+"|"+name+"|"+csv+"|"+approval+"|"+fmt.Sprint(approved)+Colors.White)
				}

			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		}

		// The following "esle" statement is when we select a single project
	} else if List1Item == "Projects" && List2Item != "All Projects" {
		if List3Item == "Summary" {
			TextView.SetText("To Be Implemented")
		} else if List3Item == "YAML" {
			// Get project's YAML
			fileInfo, _ := os.ReadFile(MG_Path + "namespaces/" + List2Item + "/" + List2Item + ".yaml")
			TextView.SetText(string(fileInfo))
			TextViewData = TextView.GetText(false)
		} else if List3Item == "Events" {
			// Get project's events "if exists"
			// This should be a table
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/events.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			if len(items) > 0 {
				Output = []string{Colors.Yellow + "Time" + "|" + "Type" + "|" + "Message" + "|" + "Reason" + Colors.White}
				for i := 0; i < len(items); i++ {
					y := items[i].(map[interface{}]interface{})
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
					Output = append(Output, Colors.White+eventTime.(string)+"|"+eventType.(string)+"|"+eventReason.(string)+"|"+eventMessage.(string)+"\n")
				}
				FormatedOutput = columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
			} else {
				TextView.SetText("No Recorded Events")
				TextViewData = TextView.GetText(false)
			}

		} else if List3Item == "Pods" {
			TextView.Clear()
			TextViewData = ""
			List4.SetTitle("Pods")
			Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "READY" + Colors.Yellow + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + Colors.Yellow + "RESTARTS" + Colors.Yellow + "|" + "AGE" + Colors.White}
			// Get project's pods "regardless of it's type/status" I can add the owner column if possible
			File, _ = ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/pods.yaml")
			MyPods := PODS{}
			yaml.Unmarshal(File, &MyPods)
			items := MyPods.Items
			now := time.Now().UTC()
			// Loop between pods
			for i := 0; i < len(items); i++ {
				name := items[i].Metadata.Name

				List4.AddItem(name, "", 0, nil)
				Status := items[i].Status.Phase
				if Status == "Running" || Status == "Succeeded" || Status == "Completed" {
					Status = Colors.White + Status + Colors.White
				} else {
					Status = Colors.Red + Status + Colors.White
				}
				// Initializing Ready Container count and number of restarts
				readyCount := 0
				restartsCount := 0
				containerStatuses := items[i].Status.ContainerStatuses
				for x := 0; x < len(containerStatuses); x++ {
					restartCount := containerStatuses[x].RestartCount

					if containerStatuses[x].Ready { // containerStatuses[x].Ready == true
						readyCount++
					}
					restartsCount += restartCount
				}
				restartsS := ""
				if restartsCount > 0 {
					restartsS = Colors.Orange + fmt.Sprintf("%v", restartsCount) + Colors.White
				} else {
					restartsS = Colors.White + fmt.Sprintf("%v", restartsCount) + Colors.White
				}
				containers := len(items[i].Spec.Containers)
				containersS := fmt.Sprintf("%v", containers)
				readyConutS := fmt.Sprintf("%v", readyCount)
				READY := ""
				if readyCount != containers {
					READY = Colors.Orange + readyConutS + "/" + containersS + Colors.White
				} else {
					READY = Colors.White + readyConutS + "/" + containersS + Colors.White
				}
				CreationTime := items[i].Metadata.CreationTimestamp
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
				Output = append(Output, Colors.White+name+"|"+READY+"|"+Status+"|"+restartsS+"|"+age+Colors.White)
			}
			FormatedOutput = columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput

		} else if List3Item == "Deployment" {
			// Get projects deployments "if exists"
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps/deployments.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			if len(items) > 0 {
				List4.SetTitle("Deployments")
				Output = []string{Colors.Yellow + "NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE" + Colors.White}
				for i := 0; i < len(items); i++ {
					y := items[i].(map[interface{}]interface{})
					now := time.Now().UTC()

					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)
					List4.AddItem(nameS, "", 0, nil)

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

					Output = append(Output, Colors.White+nameS+"|"+readyS+"|"+UpToDate+"|"+availableS+"|"+age+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			} else {
				TextView.SetText("No Deployment resources found")
			}
		} else if List3Item == "DeploymentConfig" {
			// Get projects deploymentconfigs "if exists"
			// Get projects deployments "if exists"
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps.openshift.io/deploymentconfigs.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			if len(items) > 0 {
				List4.SetTitle("DeploymentConfig")
				Output = []string{Colors.Yellow + "NAME" + "|" + "REVISION" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "TRIGGERED BY" + Colors.White}
				for i := 0; i < len(items); i++ {
					y := items[i].(map[interface{}]interface{})
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

					Output = append(Output, Colors.White+nameS+"|"+revisionS+"|"+desiredS+"|"+currentS+"|"+triggersType+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			} else {
				TextView.SetText("No DeploymentConfig resources found")
			}
		} else if List3Item == "Daemonset" {
			// Get projects daemonsets "if exists"
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/apps/daemonsets.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			if len(items) > 0 {
				List4.SetTitle("Daemonsets")
				Output = []string{Colors.Yellow + "NAME" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "NODE SELECTOR" + "|" + "AGE" + Colors.White}
				for i := 0; i < len(items); i++ {
					now := time.Now().UTC()
					y := items[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)
					List4.AddItem(nameS, "", 0, nil)

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
					nodeselector := y["spec"].(map[interface{}]interface{})["template"].(map[interface{}]interface{})["spec"].(map[interface{}]interface{})["nodeSelector"].(map[interface{}]interface{})
					// nodeselectorS := fmt.Sprintf("%v", nodeselector)
					nodeselectorS := ""
					count := len(nodeselector)
					for key, value := range nodeselector {
						if count > 1 {
							if fmt.Sprintf("%v", value) != "" {
								nodeselectorS = nodeselectorS + fmt.Sprintf("%v", key) + ":" + fmt.Sprintf("%v", value) + ", "
							} else {
								nodeselectorS = nodeselectorS + fmt.Sprintf("%v", key) + ", "
							}

							count--
						} else {
							if fmt.Sprintf("%v", value) != "" {
								nodeselectorS = nodeselectorS + fmt.Sprintf("%v", key) + ":" + fmt.Sprintf("%v", value)
							} else {
								nodeselectorS = nodeselectorS + fmt.Sprintf("%v", key)
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

					Output = append(Output, Colors.White+nameS+"|"+desiredS+"|"+currentS+"|"+readyS+"|"+uptodateS+"|"+availableS+"|"+nodeselectorS+"|"+age+"|"+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			} else {
				TextView.SetText("No Daemonset resources found")
			}
		} else if List3Item == "Services" {
			// Get projects services "if exists"
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/services.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			Output = []string{Colors.Yellow + "NAME" + "|" + "TYPE" + "|" + "CLUSTER-IP" + "|" + "EXTERNAL-IP" + "|" + "PORT(S)" + "|" + "AGE" + Colors.White}
			yaml.Unmarshal(yfile, m)
			if len(items) > 0 {
				List4.SetTitle("Services")
				for i := 0; i < len(items); i++ {
					y := items[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)
					List4.AddItem(nameS, "", 0, nil)

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

					Output = append(Output, Colors.White+nameS+"|"+StypeS+"|"+clusterIPS+"|"+externalIPS+"|"+ports_proto+"|"+age+"|"+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}
		} else if List3Item == "Routes" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			Output = []string{Colors.Yellow + "NAME" + "|" + "HOST/PORT" + "|" + "PATH" + "|" + "SERVICES" + "|" + "PORT" + "|" + "TERMINATION" + "|" + "WILDCARD" + "|" + "AGE" + Colors.White}
			File, _ = ioutil.ReadFile(Namespaces_Path + List2Item + "/route.openshift.io/routes.yaml")
			MyRoutes := ROUTES{}
			yaml.Unmarshal(File, &MyRoutes)
			items := MyRoutes.Items
			if len(items) > 0 {
				List4.SetTitle("Routes")
				for i := 0; i < len(items); i++ {
					name, host, services, port, term, wildcard, age := GetRouteDetails(items[i])

					Output = append(Output, Colors.White+name+"|"+host+"|"+"TBA"+"|"+services+"|"+port+"|"+term+"|"+wildcard+"|"+age+"|"+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			} else {
				TextView.SetText("No Available Data")
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}
		} else if List3Item == "Image Stream" {
			// Get projects IS "if exists"
			// Table of Projects/Image Stream
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output = []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "TAGS" + "|" + "AGE 	" + Colors.White}
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/image.openshift.io/imagestreams.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			if len(items) > 0 {
				List4.SetTitle("Image Streams")
				for i := 0; i < len(items); i++ {
					y := items[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)
					List4.AddItem(nameS, "", 0, nil)

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

					Output = append(Output, Colors.White+nameS+List2Item+"|"+nameS+"|"+tagsS+"|"+age+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if List3Item == "PVC" {
			// Get projects PVC "if exists"
			// Table of Projects/PVC
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output = []string{Colors.Yellow + "NAME" + "|" + "STATUS" + "|" + "VOLUME 	" + "|" + "CAPACITY" + "|" + "ACCESS MODES" + "|" + "STORAGECLASS" + "|" + "AGE" + Colors.White}
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/persistentvolumeclaims.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			if len(items) > 0 {
				List4.SetTitle("PVCs")
				for i := 0; i < len(items); i++ {
					y := items[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)
					List4.AddItem(nameS, "", 0, nil)

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

					Output = append(Output, Colors.White+nameS+"|"+statusS+"|"+volumeS+"|"+capacityS+"|"+accessS+"|"+storageCS+"|"+age+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if List3Item == "ConfigMap" {
			// Get projects CM "if exists"
			// Table of Projects/ConfigMap
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output := []string{Colors.Yellow + "NAME" + "|" + "DATA" + "|" + "AGE" + Colors.White}
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/configmaps.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			if len(items) > 0 {
				List4.SetTitle("ConfigMaps")
				for i := 0; i < len(items); i++ {
					y := items[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)
					List4.AddItem(nameS, "", 0, nil)

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

					Output = append(Output, Colors.White+nameS+"|"+dataS+"|"+age+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if List3Item == "Secrets" {
			// Get projects secrets  "if exists"
			// Table of Projects/Secrets
			// Table of Projects/ConfigMap
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			// Getting current timestamp
			now := time.Now().UTC()
			Output := []string{Colors.Yellow + "NAME" + "|" + "TYPE" + "|" + "DATA" + "|" + "AGE" + Colors.White}
			yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/core/secrets.yaml")
			m := make(map[interface{}]interface{})
			yaml.Unmarshal(yfile, m)
			items, _ := m["items"].([]interface{})
			if len(items) > 0 {
				List4.SetTitle("Secrets")
				for i := 0; i < len(items); i++ {
					y := items[i].(map[interface{}]interface{})
					name := y["metadata"].(map[interface{}]interface{})["name"]
					nameS := fmt.Sprintf("%v", name)
					List4.AddItem(nameS, "", 0, nil)

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

					Output = append(Output, Colors.White+nameS+"|"+type_keyS+"|"+dataS+"|"+age+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput

			}
		} else if List3Item == "Subscriptions" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			Output := []string{Colors.Yellow + "NAME" + "|" + "PACKAGE" + "|" + "SOURCE" + "|" + "CHANNEL" + Colors.White}
			files, _ := ioutil.ReadDir(MG_Path + "namespaces/" + List2Item + "/operators.coreos.com/subscriptions/")
			for _, file := range files {
				List4.SetTitle("Subscriptions")
				yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/operators.coreos.com/subscriptions/" + file.Name())
				m := make(map[string]interface{})
				yaml.Unmarshal(yfile, m)

				name := m["metadata"].(map[interface{}]interface{})["name"]
				nameS := fmt.Sprintf("%v", name)
				List4.AddItem(strings.Split(file.Name(), ".yaml")[0], "", 0, nil)

				Package := m["spec"].(map[interface{}]interface{})["name"]
				packageS := fmt.Sprintf("%v", Package)

				source := m["spec"].(map[interface{}]interface{})["source"]
				sourceS := fmt.Sprintf("%v", source)

				channel := m["spec"].(map[interface{}]interface{})["channel"]
				channelS := fmt.Sprintf("%v", channel)

				Output = append(Output, Colors.White+nameS+"|"+packageS+"|"+sourceS+"|"+channelS+Colors.White)

			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else if List3Item == "Operators" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			List4.SetTitle("Operators")
			Output := []string{Colors.Yellow + "NAME" + "|" + "DISPLAY" + "|" + "VERSION" + "|" + "REPLACES" + "|" + "PHASE" + Colors.White}
			files, _ := ioutil.ReadDir(MG_Path + "namespaces/" + List2Item + "/operators.coreos.com/clusterserviceversions/")
			for _, file := range files {
				yfile, _ := ioutil.ReadFile(MG_Path + "namespaces/" + List2Item + "/operators.coreos.com/clusterserviceversions/" + file.Name())
				m := make(map[string]interface{})
				yaml.Unmarshal(yfile, m)

				name := m["metadata"].(map[interface{}]interface{})["name"]
				nameS := fmt.Sprintf("%v", name)
				List4.AddItem(nameS, "", 0, nil)

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

				Output = append(Output, Colors.White+nameS+"|"+displayS+"|"+versionS+"|"+replaceS+"|"+phaseS+Colors.White)
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else if List3Item == "Install Plans" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			List4.SetTitle("Install Plans")
			Output := []string{Colors.Yellow + "NAME" + "|" + "CSV" + "|" + "APPROVAL" + "|" + "APPROVED" + Colors.White}
			Files, _ = ioutil.ReadDir(Namespaces_Path + List2Item + "/operators.coreos.com/installplans/")
			for i := range Files {
				File, _ = ioutil.ReadFile(Namespaces_Path + List2Item + "/operators.coreos.com/installplans/" + Files[i].Name())
				MyIP := INSTALLPLAN{}
				yaml.Unmarshal(File, &MyIP)
				name := MyIP.Metadata.Name
				csv := MyIP.Spec.ClusterServiceVersionNames[0]
				approval := MyIP.Spec.Approval
				approved := MyIP.Spec.Approved
				Output = append(Output, Colors.White+name+"|"+csv+"|"+approval+"|"+fmt.Sprint(approved)+Colors.White)
				List4.AddItem(strings.Split(Files[i].Name(), ".yaml")[0], "", 0, nil)
			}

			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		}
	} else if List1Item == "Nodes" && List3Item == "Summary" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		File, _ = ioutil.ReadFile(Nodes_Path + List2Item + ".yaml")
		Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + Colors.White}
		MyNode := NODE{}
		yaml.Unmarshal(File, &MyNode)
		name, status, roles, age, version, _, _, _, _, _, _ := GetNodeDetails(MyNode)
		Output = append(Output, Colors.White+name+"|"+status+"|"+roles+"|"+age+"|"+version+Colors.White)
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if List1Item == "Nodes" && List3Item == "Details" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		File, _ = ioutil.ReadFile(Nodes_Path + List2Item + ".yaml")
		Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "|" + "INTERNAL-IP" + "|" + "EXTERNAL-IP" + "|" + "OS-IMAGE" + "|" + "KERNEL-VERSION" + "|" + "CONTAINER-RUNTIME" + Colors.White + "\n"}
		MyNode := NODE{}
		yaml.Unmarshal(File, &MyNode)
		name, status, roles, age, version, internalIP, externalIP, osImage, kernelVersion, contRuntime, _ := GetNodeDetails(MyNode)
		Output = append(Output, Colors.White+name+"|"+status+"|"+roles+"|"+age+"|"+version+"|"+internalIP+"|"+externalIP+"|"+osImage+"|"+kernelVersion+"|"+contRuntime+Colors.White)
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if List1Item == "Nodes" && List3Item == "YAML" {
		// Get node's YAML
		File, _ = os.ReadFile(MG_Path + "cluster-scoped-resources/core/nodes/" + List2Item + ".yaml")
		TextView.SetText(string(File))
		TextViewData = TextView.GetText(false)
		TextView.ScrollToBeginning()
		List6.AddItem("Metadata", "", 0, nil).
			AddItem("Spec", "", 0, nil).
			AddItem("Status", "", 0, nil).
			AddItem("HW Specs", "", 0, nil).
			AddItem("Images", "", 0, nil).
			AddItem("nodeInfo", "", 0, nil)
	} else if List1Item == "Operators" && List3Item == "YAML" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		operator, _ := ioutil.ReadFile(MG_Path + "cluster-scoped-resources/config.openshift.io/clusteroperators/" + List2Item + ".yaml")
		TextView.SetText(string(operator))
		TextView.ScrollToBeginning()
	} else if List1Item == "Operators" && List3Item == "Summary" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		operatorFile, _ := ioutil.ReadFile(MG_Path + "cluster-scoped-resources/config.openshift.io/clusteroperators/" + List2Item + ".yaml")

		Output := []string{Colors.Yellow + "NAME" + "|" + "VERSION" + "|" + "AVAILABLE" + "|" + "PROGRESSINS" + "|" + "DEGRADED" + "|" + "SINCE" + Colors.White}
		operator := make(map[interface{}]interface{})
		yaml.Unmarshal([]byte(operatorFile), operator)

		name := operator["metadata"].(map[interface{}]interface{})["name"]
		nameS := fmt.Sprintf("%v", name)

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
		// fmt.Print(nameS + "\t" + versionS + "\t" + availableS + "\t" + progressingS + "\t" + degradedS + "\t" + availableSince + Colors.White)
		Output = append(Output, Colors.White+nameS+"|"+versionS+"|"+availableS+"|"+progressingS+"|"+degradedS+"|"+availableSince+Colors.White)

		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if List1Item == "Installed Operators" && List3Item == "YAML" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		File, _ = ioutil.ReadFile(InstalledOperators_Path + List2Item + ".yaml")
		TextView.SetText(string(File))
		TextView.ScrollToBeginning()
	} else if List1Item == "Installed Operators" && List3Item == "Summary" {
		TextView.Clear()
		TextViewData = ""
		// Get installed operators file
		Output = []string{Colors.Yellow + "NAME" + "|" + "AGE" + Colors.White}
		File, _ = ioutil.ReadFile(InstalledOperators_Path + List2Item + ".yaml")
		MyOperator := OPERATOR{}
		yaml.Unmarshal(File, &MyOperator)
		name := MyOperator.Metadata.Name
		age := GetAge(MyOperator.Metadata.CreationTimestamp)
		Output = append(Output, Colors.White+name+"|"+age+Colors.White)

		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if List1Item == "MCP" && List3Item == "Info" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		File, _ = ioutil.ReadFile(MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigpools/" + List2Item + ".yaml")
		GetMCPInfo(File)

	} else if List1Item == "MCP" && List3Item == "YAML" {
		mcpFile, _ := ioutil.ReadFile(MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigpools/" + List2Item + ".yaml")
		TextView.SetText(string(mcpFile))
		TextView.ScrollToBeginning()
	} else if List1Item == "MC" && List3Item == "Summary" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		now := time.Now().UTC()
		Output := []string{Colors.Yellow + "NAME" + "|" + "GENERATEDBYCONTROLLER" + "|" + "IGNITIONVERSION" + "|" + "AGE" + Colors.White}
		File, _ = ioutil.ReadFile(MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigs/" + List2Item + ".yaml")

		MyMC := MC{}
		yaml.Unmarshal(File, &MyMC)

		name := MyMC.Metadata.Name
		List2.AddItem(name, "", 0, nil)

		generatedByMap := MyMC.Metadata.Annotations
		generatedBy := generatedByMap["machineconfiguration.openshift.io/generated-by-controller-version"]

		ignitionVersion := MyMC.Spec.Config.Ignition.Version

		CreationTime := MyMC.Metadata.CreationTimestamp
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
		Output = append(Output, Colors.White+name+"|"+generatedBy+"|"+ignitionVersion+"|"+age+Colors.White)

		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if List1Item == "MC" && List3Item == "YAML" {
		mcFile, _ := ioutil.ReadFile(MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigs/" + List2Item + ".yaml")
		TextView.SetText(string(mcFile))
		TextView.ScrollToBeginning()

	} else if List1Item == "MC" && List3Item == "Data" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""

		List4.SetTitle("Files")

		File, _ = ioutil.ReadFile(MC_Path + List2Item + ".yaml")

		MyMC := MC{}
		yaml.Unmarshal(File, &MyMC)

		paths := MyMC.Spec.Config.Storage.Files
		for i := range paths {
			path := fmt.Sprintf("%v", paths[i].Path)
			List4.AddItem(path, "", 0, nil)
		}

	} else if List1Item == "PV" && List3Item == "Info" {

		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		now := time.Now().UTC()
		Output := []string{Colors.Yellow + "NAME" + "|" + "CAPACITY" + "|" + "ACCESS MODE" + "|" + "RECLAIM POLICY" + "|" + "STATUS" + "|" + "CLAIM" + "|" + "STORAGECLASS" + "|" + "AGE" + Colors.White}
		File, _ = ioutil.ReadFile(MG_Path + "cluster-scoped-resources/core/persistentvolumes/" + List2Item + ".yaml")
		MyPV := PV{}
		yaml.Unmarshal(File, &MyPV)

		name := MyPV.Metadata.Name
		List2.AddItem(name, "", 0, nil)

		capacity := MyPV.Spec.Capacity.Storage

		accessArray := MyPV.Spec.AccessModes
		access := fmt.Sprintf("%v", accessArray)
		access = strings.Replace(access, "[", "", -1)
		access = strings.Replace(access, "]", "", -1)

		reclaim := MyPV.Spec.PersistentVolumeReclaimPolicy

		status := MyPV.Status.Phase

		claim := MyPV.Spec.ClaimRef.Name

		storageclass := MyPV.Spec.StorageClassName

		CreationTime := MyPV.Metadata.CreationTimestamp
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

		Output = append(Output, Colors.White+name+"|"+capacity+"|"+access+"|"+reclaim+"|"+status+"|"+claim+"|"+storageclass+"|"+age+Colors.White)

		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput

	} else if List1Item == "PV" && List3Item == "YAML" {
		pvFile, _ := ioutil.ReadFile(MG_Path + "cluster-scoped-resources/core/persistentvolumes/" + List2Item + ".yaml")
		TextView.SetText(string(pvFile))
		TextView.ScrollToBeginning()
	}
}
