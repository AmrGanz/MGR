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
	ActivePathBox.SetText(List1Item + " > " + List2Item + " > " + List3Item)
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
			Files, _ = ioutil.ReadDir(Namespaces_Path)
			if len(Files) > 0 {
				// Cleaning TextView and TextViewData
				TextView.Clear()
				TextViewData = ""
				// Getting current timestamp
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + Colors.Yellow + "READY" + Colors.Yellow + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + Colors.Yellow + "RESTARTS" + Colors.Yellow + "|" + "Age" + Colors.White}
				for projectIndex := 0; projectIndex < len(Files); projectIndex++ {
					// Get project's pods "regardless of it's type/status" I can add the owner column if possible
					if _, err := os.Stat(Namespaces_Path + Files[projectIndex].Name() + "/core/pods.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + Files[projectIndex].Name() + "/core/pods.yaml")
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
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/apps/deployments.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/apps/deployments.yaml")
						MyDeployments := DEPLOYMENTS{}
						yaml.Unmarshal(File, &MyDeployments)

						items := MyDeployments.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								Deployment := MyDeployments.Items[i]
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
				//Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/apps.openshift.io/deploymentconfigs.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/apps.openshift.io/deploymentconfigs.yaml")

						MyDCs := DEPLOYMENTCONFIGS{}
						yaml.Unmarshal(File, &MyDCs)

						items := MyDCs.Items
						if len(items) > 0 {
							Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "REVISION" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "TRIGGERED BY" + "|" + "Age" + Colors.White}
							for i := 0; i < len(items); i++ {
								DC := MyDCs.Items[i]
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
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "NODE SELECTOR" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/apps/daemonsets.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/apps/daemonsets.yaml")
						MyDaemonset := DAEMONSETS{}
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
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "TYPE" + "|" + "CLUSTER-IP" + "|" + "EXTERNAL-IP" + "|" + "PORT(S)" + "|" + "AGE" + Colors.White}
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

								externalIP := Service.Spec.ExternalName

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
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "HOST/PORT" + "|" + "PATH" + "|" + "SERVICES" + "|" + "PORT" + "|" + "TERMINATION" + "|" + "WILDCARD" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/route.openshift.io/routes.yaml"); err == nil {
						File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/route.openshift.io/routes.yaml")
						MyRoutes := ROUTES{}
						yaml.Unmarshal(File, &MyRoutes)
						items := MyRoutes.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								name, host, path, services, port, term, wildcard, age := GetRouteDetails(items[i])
								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+host+"|"+path+"|"+services+"|"+port+"|"+term+"|"+wildcard+"|"+age+"|"+"\n")
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
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "TAGS" + "|" + "AGE 	" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/image.openshift.io/imagestreams.yaml"); err == nil {
						File, _ := ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/image.openshift.io/imagestreams.yaml")
						MyImageStreams := IMAGESTREAMS{}
						yaml.Unmarshal(File, &MyImageStreams)
						items := MyImageStreams.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								name := items[i].Metadata.Name

								tags := ""
								if items[i].Spec.Tags != nil {
									all_tags := items[i].Spec.Tags
									for t := 0; t < len(all_tags); t++ {
										tag_name := items[i].Spec.Tags[t].Name
										tags = tags + tag_name
										if i != len(all_tags)-1 {
											tags = tags + ","
										}
									}
								}

								age := GetAge(items[i].Metadata.CreationTimestamp)

								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+tags+"|"+age+Colors.White)
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
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "STATUS" + "|" + "VOLUME" + "|" + "CAPACITY" + "|" + "ACCESS MODES" + "|" + "STORAGECLASS" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml"); err == nil {
						File, _ := ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml")
						MyPVCs := PVCS{}
						yaml.Unmarshal(File, &MyPVCs)
						items := MyPVCs.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								name := items[i].Metadata.Name

								status := items[i].Status.Phase

								volume := items[i].Spec.VolumeName

								capacity := items[i].Status.Capacity.Storage

								access := items[i].Status.AccessModes
								accessS := fmt.Sprintf("%v", access)
								accessS = strings.Replace(accessS, "[", "", -1)
								accessS = strings.Replace(accessS, "]", "", -1)

								storageClassName := items[i].Spec.StorageClassName

								age := GetAge(items[i].Metadata.CreationTimestamp)
								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+status+"|"+volume+"|"+capacity+"|"+accessS+"|"+storageClassName+"|"+age+"\n")
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
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "DATA" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/core/configmaps.yaml"); err == nil {
						File, _ := ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/core/configmaps.yaml")
						MyCM := CONFIGMAPS{}
						yaml.Unmarshal(File, &MyCM)
						items := MyCM.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								name := items[i].Metadata.Name
								dataS := ""
								if items[i].Data != nil {
									data := items[i].Data
									dataN := len(data)
									dataS = fmt.Sprint(dataN)
								} else {
									dataS = "0"
								}

								age := GetAge(items[i].Metadata.CreationTimestamp)
								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+dataS+"|"+age+Colors.White)
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
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "TYPE" + "|" + "DATA" + "|" + "AGE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/core/secrets.yaml"); err == nil {
						File, _ := ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/core/secrets.yaml")
						MySecrets := SECRETS{}
						yaml.Unmarshal(File, &MySecrets)
						items := MySecrets.Items
						if len(items) > 0 {
							for i := 0; i < len(items); i++ {
								SECRET := items[i]
								name := SECRET.Metadata.Name

								dataS := ""
								if SECRET.Data != nil {
									data := SECRET.Data
									dataN := len(data)
									dataS = fmt.Sprint(dataN)
								} else {
									dataS = "0"
								}

								type_key := SECRET.Type

								CreationTime := SECRET.Metadata.CreationTimestamp
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

								Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+type_key+"|"+dataS+"|"+age+"\n")
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
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "PACKAGE" + "|" + "SOURCE" + "|" + "CHANNEL" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					Files, _ := ioutil.ReadDir(Namespaces_Path + namespaces[projectIndex].Name() + "/operators.coreos.com/subscriptions/")

					for _, File := range Files {
						yfile, _ := ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/operators.coreos.com/subscriptions/" + File.Name())
						MySubscription := SUBSCRIPTION{}
						yaml.Unmarshal(yfile, &MySubscription)

						name := MySubscription.Metadata.Name

						Package := MySubscription.Spec.Name

						source := MySubscription.Spec.Source

						channel := MySubscription.Spec.Channel

						Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+Package+"|"+source+"|"+channel+Colors.White)
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
			namespaces, _ := ioutil.ReadDir(Namespaces_Path)
			if len(namespaces) > 0 {
				Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "DISPLAY" + "|" + "VERSION" + "|" + "REPLACES" + "|" + "PHASE" + Colors.White}
				for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
					Files, _ := ioutil.ReadDir(Namespaces_Path + namespaces[projectIndex].Name() + "/operators.coreos.com/clusterserviceversions/")
					for _, File := range Files {
						yfile, _ := ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/operators.coreos.com/clusterserviceversions/" + File.Name())
						MyCSV := CSV{}
						yaml.Unmarshal(yfile, &MyCSV)

						name := MyCSV.Metadata.Name

						displayName := MyCSV.Spec.DisplayName

						version := MyCSV.Spec.Version

						replace := MyCSV.Spec.Replaces

						phase := MyCSV.Status.Phase

						Output = append(Output, Colors.White+namespaces[projectIndex].Name()+"|"+name+"|"+displayName+"|"+version+"|"+replace+"|"+phase+"\n")

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
			fileInfo, _ := os.ReadFile(Namespaces_Path + List2Item + "/" + List2Item + ".yaml")
			TextView.SetText(string(fileInfo))
			TextViewData = TextView.GetText(false)
		} else if List3Item == "Events" {
			// Get project's events "if exists"
			// This should be a table
			yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/core/events.yaml")
			MyEvents := EVENTS{}
			yaml.Unmarshal(yfile, &MyEvents)

			if len(MyEvents.Items) > 0 {
				Output := []string{Colors.Yellow + "Time" + "|" + "Type" + "|" + "Message" + "|" + "Reason" + Colors.White}
				for i := 0; i < len(MyEvents.Items); i++ {
					eventTime := MyEvents.Items[i].LastTimestamp.String()
					eventType := MyEvents.Items[i].Type
					eventReason := MyEvents.Items[i].Reason
					eventMessage := MyEvents.Items[i].Message

					Output = append(Output, Colors.White+eventTime+"|"+eventType+"|"+eventReason+"|"+eventMessage+"\n")
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
			Output := []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "READY" + Colors.Yellow + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + Colors.Yellow + "RESTARTS" + Colors.Yellow + "|" + "AGE" + Colors.White}
			// Get project's pods "regardless of it's type/status" I can add the owner column if possible
			File, _ = ioutil.ReadFile(Namespaces_Path + List2Item + "/core/pods.yaml")
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
			yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/apps/deployments.yaml")
			MyDelployments := DEPLOYMENTS{}
			yaml.Unmarshal(yfile, &MyDelployments)

			if len(MyDelployments.Items) > 0 {
				List4.SetTitle("Deployments")
				Output := []string{Colors.Yellow + "NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE" + Colors.White}

				for i := 0; i < len(MyDelployments.Items); i++ {
					now := time.Now().UTC()

					name := MyDelployments.Items[i].Metadata.Name
					List4.AddItem(name, "", 0, nil)

					ready := strconv.Itoa(MyDelployments.Items[i].Status.ReadyReplicas)
					UpToDate := "TBA"
					// I think I should print Ready/Avilable just like in the output of [# oc get deployment]
					available := strconv.Itoa(MyDelployments.Items[i].Status.AvailableReplicas)

					CreationTime := MyDelployments.Items[i].Metadata.CreationTimestamp
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

					Output = append(Output, Colors.White+name+"|"+ready+"|"+UpToDate+"|"+available+"|"+age+Colors.White)
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
			yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/apps.openshift.io/deploymentconfigs.yaml")
			MyDCs := DEPLOYMENTCONFIGS{}
			yaml.Unmarshal(yfile, &MyDCs)

			if len(MyDCs.Items) > 0 {
				List4.SetTitle("DeploymentConfig")
				Output := []string{Colors.Yellow + "NAME" + "|" + "REVISION" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "TRIGGERED BY" + Colors.White}

				for i := 0; i < len(MyDCs.Items); i++ {

					name := MyDCs.Items[i].Metadata.Name

					List4.AddItem(name, "", 0, nil)

					revision := strconv.Itoa(MyDCs.Items[i].Status.LatestVersion)

					desired := strconv.Itoa(MyDCs.Items[i].Spec.Replicas)

					current := strconv.Itoa(MyDCs.Items[i].Status.AvailableReplicas)

					triggers := MyDCs.Items[i].Spec.Triggers
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

					Output = append(Output, Colors.White+name+"|"+revision+"|"+desired+"|"+current+"|"+triggersType+Colors.White)
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
			yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/apps/daemonsets.yaml")
			MyDSs := DAEMONSETS{}
			yaml.Unmarshal(yfile, &MyDSs)

			if len(MyDSs.Items) > 0 {
				List4.SetTitle("Daemonsets")
				Output := []string{Colors.Yellow + "NAME" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "NODE SELECTOR" + "|" + "AGE" + Colors.White}
				for i := 0; i < len(MyDSs.Items); i++ {
					now := time.Now().UTC()

					name := MyDSs.Items[i].Metadata.Name
					List4.AddItem(name, "", 0, nil)

					desired := strconv.Itoa(MyDSs.Items[i].Status.DesiredNumberScheduled)

					current := strconv.Itoa(MyDSs.Items[i].Status.CurrentNumberScheduled)

					ready := strconv.Itoa(MyDSs.Items[i].Status.NumberReady)

					uptodate := strconv.Itoa(MyDSs.Items[i].Status.UpdatedNumberScheduled)

					available := strconv.Itoa(MyDSs.Items[i].Status.NumberAvailable)

					// To be enhanced
					nodeselector := MyDSs.Items[i].Spec.Template.Spec.NodeSelector
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

					CreationTime := MyDSs.Items[i].Metadata.CreationTimestamp
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

					Output = append(Output, Colors.White+name+"|"+desired+"|"+current+"|"+ready+"|"+uptodate+"|"+available+"|"+nodeselectorS+"|"+age+"|"+Colors.White)
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
			yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/core/services.yaml")

			MyServices := SERVICES{}
			yaml.Unmarshal(yfile, &MyServices)

			Output := []string{Colors.Yellow + "NAME" + "|" + "TYPE" + "|" + "CLUSTER-IP" + "|" + "EXTERNAL-IP" + "|" + "PORT(S)" + "|" + "AGE" + Colors.White}

			if len(MyServices.Items) > 0 {
				List4.SetTitle("Services")
				for i := 0; i < len(MyServices.Items); i++ {

					name := MyServices.Items[i].Metadata.Name
					List4.AddItem(name, "", 0, nil)

					Stype := MyServices.Items[i].Spec.Type

					clusterIP := MyServices.Items[i].Spec.ClusterIP

					externalIP := MyServices.Items[i].Spec.ExternalName

					CreationTime := MyServices.Items[i].Metadata.CreationTimestamp
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
					if MyServices.Items[i].Spec.Ports != nil {
						ports := MyServices.Items[i].Spec.Ports
						for x := 0; x < len(ports); x++ {
							port := strconv.Itoa(MyServices.Items[i].Spec.Ports[x].Port)
							protocol := MyServices.Items[i].Spec.Ports[x].Protocol
							ports_proto = ports_proto + port + "/" + protocol
							if x != len(ports)-1 {
								ports_proto = ports_proto + ","
							}
						}
					} else {
						ports_proto = "None"
					}

					Output = append(Output, Colors.White+name+"|"+Stype+"|"+clusterIP+"|"+externalIP+"|"+ports_proto+"|"+age+"|"+Colors.White)
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
			Output := []string{Colors.Yellow + "NAME" + "|" + "HOST/PORT" + "|" + "PATH" + "|" + "SERVICES" + "|" + "PORT" + "|" + "TERMINATION" + "|" + "WILDCARD" + "|" + "AGE" + Colors.White}
			File, _ = ioutil.ReadFile(Namespaces_Path + List2Item + "/route.openshift.io/routes.yaml")
			MyRoutes := ROUTES{}
			yaml.Unmarshal(File, &MyRoutes)
			items := MyRoutes.Items
			if len(items) > 0 {
				List4.SetTitle("Routes")
				for i := 0; i < len(items); i++ {
					name, host, path, services, port, term, wildcard, age := GetRouteDetails(items[i])
					List4.AddItem(name, "", 0, nil)
					Output = append(Output, Colors.White+name+"|"+host+"|"+path+"|"+services+"|"+port+"|"+term+"|"+wildcard+"|"+age+"|"+Colors.White)
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
			Output := []string{Colors.Yellow + "NAMESPACE" + "|" + "NAME" + "|" + "TAGS" + "|" + "AGE" + Colors.White}
			File, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/image.openshift.io/imagestreams.yaml")
			MyImageStreams := IMAGESTREAMS{}
			yaml.Unmarshal(File, &MyImageStreams)

			if len(MyImageStreams.Items) > 0 {
				List4.SetTitle("Image Streams")
				for i := 0; i < len(MyImageStreams.Items); i++ {

					name := MyImageStreams.Items[i].Metadata.Name
					List4.AddItem(name, "", 0, nil)

					tagsS := ""
					if MyImageStreams.Items[i].Spec.Tags != nil {
						for x := 0; x < len(MyImageStreams.Items[i].Spec.Tags); x++ {
							tag_name := MyImageStreams.Items[i].Spec.Tags[x].Name
							tagsS = tagsS + tag_name
							if i != len(MyImageStreams.Items[i].Spec.Tags)-1 {
								tagsS = tagsS + ","
							}
						}
					}

					CreationTime := MyImageStreams.Items[i].Metadata.CreationTimestamp
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

					Output = append(Output, Colors.White+name+List2Item+"|"+name+"|"+tagsS+"|"+age+Colors.White)
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
			Output := []string{Colors.Yellow + "NAME" + "|" + "STATUS" + "|" + "VOLUME 	" + "|" + "CAPACITY" + "|" + "ACCESS MODES" + "|" + "STORAGECLASS" + "|" + "AGE" + Colors.White}
			yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/core/persistentvolumeclaims.yaml")

			MyPVCs := PVCS{}
			yaml.Unmarshal(yfile, &MyPVCs)

			if len(MyPVCs.Items) > 0 {
				List4.SetTitle("PVCs")
				for i := 0; i < len(MyPVCs.Items); i++ {

					name := MyPVCs.Items[i].Metadata.Name
					List4.AddItem(name, "", 0, nil)

					status := MyPVCs.Items[i].Status.Phase

					volume := MyPVCs.Items[i].Spec.VolumeName

					capacity := MyPVCs.Items[i].Status.Capacity.Storage

					access := MyPVCs.Items[i].Status.AccessModes
					accessS := fmt.Sprintf("%v", access)
					accessS = strings.Replace(accessS, "[", "", -1)
					accessS = strings.Replace(accessS, "]", "", -1)

					storageClass := MyPVCs.Items[i].Spec.StorageClassName

					CreationTime := MyPVCs.Items[i].Metadata.CreationTimestamp
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

					Output = append(Output, Colors.White+name+"|"+status+"|"+volume+"|"+capacity+"|"+accessS+"|"+storageClass+"|"+age+Colors.White)
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
			yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/core/configmaps.yaml")

			MyCMs := CONFIGMAPS{}
			yaml.Unmarshal(yfile, &MyCMs)

			if len(MyCMs.Items) > 0 {
				List4.SetTitle("ConfigMaps")
				for i := 0; i < len(MyCMs.Items); i++ {

					name := MyCMs.Items[i].Metadata.Name
					List4.AddItem(name, "", 0, nil)

					dataCount := ""
					if MyCMs.Items[i].Data != nil {
						dataCount = strconv.Itoa(len(MyCMs.Items[i].Data))
					} else {
						dataCount = "0"
					}

					CreationTime := MyCMs.Items[i].Metadata.CreationTimestamp
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

					Output = append(Output, Colors.White+name+"|"+dataCount+"|"+age+Colors.White)
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
			yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/core/secrets.yaml")

			MySecrets := SECRETS{}
			yaml.Unmarshal(yfile, &MySecrets)

			if len(MySecrets.Items) > 0 {
				List4.SetTitle("Secrets")
				for i := 0; i < len(MySecrets.Items); i++ {

					name := MySecrets.Items[i].Metadata.Name
					List4.AddItem(name, "", 0, nil)

					dataCount := ""
					if MySecrets.Items[i].Data != nil {
						dataCount = strconv.Itoa(len(MySecrets.Items[i].Data))

					} else {
						dataCount = "0"
					}

					typeS := MySecrets.Items[i].Type

					CreationTime := MySecrets.Items[i].Metadata.CreationTimestamp
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

					Output = append(Output, Colors.White+name+"|"+typeS+"|"+dataCount+"|"+age+Colors.White)
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
			files, _ := ioutil.ReadDir(Namespaces_Path + List2Item + "/operators.coreos.com/subscriptions/")
			for _, file := range files {
				List4.SetTitle("Subscriptions")
				yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/operators.coreos.com/subscriptions/" + file.Name())

				MySub := SUBSCRIPTION{}
				yaml.Unmarshal(yfile, &MySub)

				name := MySub.Metadata.Name
				List4.AddItem(strings.Split(file.Name(), ".yaml")[0], "", 0, nil)

				packageS := MySub.Spec.Name

				source := MySub.Spec.Source

				channel := MySub.Spec.Channel

				Output = append(Output, Colors.White+name+"|"+packageS+"|"+source+"|"+channel+Colors.White)

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
			files, _ := ioutil.ReadDir(Namespaces_Path + List2Item + "/operators.coreos.com/clusterserviceversions/")
			for _, file := range files {
				yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/operators.coreos.com/clusterserviceversions/" + file.Name())

				MyOperator := CSV{}
				yaml.Unmarshal(yfile, &MyOperator)

				name := MyOperator.Metadata.Name
				List4.AddItem(name, "", 0, nil)

				displayName := MyOperator.Spec.DisplayName

				version := MyOperator.Spec.Version

				replaceS := MyOperator.Spec.Replaces

				phase := MyOperator.Status.Phase

				Output = append(Output, Colors.White+name+"|"+displayName+"|"+version+"|"+replaceS+"|"+phase+Colors.White)
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
		Output := []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + Colors.White}
		name, status, roles, age, version, _, _, _, _, _, _ := GetNodeDetails(File)
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
		Output := []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "|" + "INTERNAL-IP" + "|" + "EXTERNAL-IP" + "|" + "OS-IMAGE" + "|" + "KERNEL-VERSION" + "|" + "CONTAINER-RUNTIME" + Colors.White + "\n"}
		name, status, roles, age, version, internalIP, externalIP, osImage, kernelVersion, contRuntime, _ := GetNodeDetails(File)
		Output = append(Output, Colors.White+name+"|"+status+"|"+roles+"|"+age+"|"+version+"|"+internalIP+"|"+externalIP+"|"+osImage+"|"+kernelVersion+"|"+contRuntime+Colors.White)
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if List1Item == "Nodes" && List3Item == "YAML" {
		// Get node's YAML
		File, _ = os.ReadFile(Nodes_Path + List2Item + ".yaml")
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
		operator, _ := ioutil.ReadFile(Operator_Path + List2Item + ".yaml")
		TextView.SetText(string(operator))
		TextView.ScrollToBeginning()
	} else if List1Item == "Operators" && List3Item == "Summary" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		yfile, _ := ioutil.ReadFile(Operator_Path + List2Item + ".yaml")

		Output := []string{Colors.Yellow + "NAME" + "|" + "VERSION" + "|" + "AVAILABLE" + "|" + "PROGRESSINS" + "|" + "DEGRADED" + "|" + "SINCE" + Colors.White}

		MyClusterOperator := CLUSTEROPERATOR{}
		yaml.Unmarshal(yfile, &MyClusterOperator)

		name := MyClusterOperator.Metadata.Name

		version := ""
		for i := range MyClusterOperator.Status.Versions {
			if MyClusterOperator.Status.Versions[i].Name == "operator" {
				version = MyClusterOperator.Status.Versions[i].Version
			}
		}
		conditions := MyClusterOperator.Status.Conditions
		available := ""
		progressingS := ""
		degradedS := ""
		availableSince := ""
		for i := range conditions {
			if conditions[i].Type == "Available" {
				if conditions[i].Status == "True" {
					available = "True"
				} else {
					available = "False"
				}

				now := time.Now().UTC()
				statusTime := conditions[i].LastTransitionTime
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

			} else if conditions[i].Type == "Progressing" {
				if conditions[i].Status == "True" {
					progressingS = "True"
				} else {
					progressingS = "False"
				}
			} else if conditions[i].Type == "Degraded" {
				if conditions[i].Status == "True" {
					degradedS = "True"
				} else {
					degradedS = "False"
				}
			}

		}
		// fmt.Print(nameS + "\t" + versionS + "\t" + availableS + "\t" + progressingS + "\t" + degradedS + "\t" + availableSince + Colors.White)
		Output = append(Output, Colors.White+name+"|"+version+"|"+available+"|"+progressingS+"|"+degradedS+"|"+availableSince+Colors.White)

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
		Output := []string{Colors.Yellow + "NAME" + "|" + "AGE" + Colors.White}
		File, _ = ioutil.ReadFile(InstalledOperators_Path + List2Item + ".yaml")
		MyOperator := CLUSTEROPERATOR{}
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
		File, _ = ioutil.ReadFile(MCP_Path + List2Item + ".yaml")
		GetMCPInfo(File)

	} else if List1Item == "MCP" && List3Item == "YAML" {
		mcpFile, _ := ioutil.ReadFile(MCP_Path + List2Item + ".yaml")
		TextView.SetText(string(mcpFile))
		TextView.ScrollToBeginning()
	} else if List1Item == "MC" && List3Item == "Summary" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		now := time.Now().UTC()
		Output := []string{Colors.Yellow + "NAME" + "|" + "GENERATEDBYCONTROLLER" + "|" + "IGNITIONVERSION" + "|" + "AGE" + Colors.White}
		File, _ = ioutil.ReadFile(MC_Path + List2Item + ".yaml")

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
		mcFile, _ := ioutil.ReadFile(MC_Path + List2Item + ".yaml")
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
		File, _ = ioutil.ReadFile(PV_Path + List2Item + ".yaml")
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
		pvFile, _ := ioutil.ReadFile(PV_Path + List2Item + ".yaml")
		TextView.SetText(string(pvFile))
		TextView.ScrollToBeginning()
	}
}
