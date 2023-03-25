package functions

import (
	"encoding/json"
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

	ActivePathBox.SetText(List1Item + " > " + List2Item + " > " + List3Item + " > " + List4Item + " > " + List5Item + " > " + List6Item)
	if List1Item == "Projects" && List3Item == "Pods" {
		// Print Container's logs
		TextView.Clear()
		File, _ = ioutil.ReadFile(Namespaces_Path + List2Item + "/pods/" + List4Item + "/" + List5Item + "/" + List5Item + "/logs/" + List6Item + ".log")
		// [reminder] Word Wrapping is disabled in "Grids.go" for a better performance
		TextView.SetText(string(File))
		TextViewData = string(File)

	} else if List1Item == "Projects" && List3Item == "Deployment" && List6Item == "Info" {
		// Get a project's deployments "if exists"
		yfile, _ := os.ReadFile(Namespaces_Path + List2Item + "/apps/deployments.yaml")

		MyDeployments := DEPLOYMENTS{}
		yaml.Unmarshal(yfile, &MyDeployments)

		if len(MyDeployments.Items) > 0 {
			Output = []string{"NAME" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "AGE"}
			for i := 0; i < len(MyDeployments.Items); i++ {
				if MyDeployments.Items[i].Metadata.Name == List4Item {

					now := time.Now().UTC()

					name := MyDeployments.Items[i].Metadata.Name

					readyCount := strconv.Itoa(MyDeployments.Items[i].Status.ReadyReplicas)

					UpToDate := "TBA"
					// I think I should print Ready/Avilable just like in the output of [# oc get deployment]
					availableCount := strconv.Itoa(MyDeployments.Items[i].Status.AvailableReplicas)

					CreationTime := MyDeployments.Items[i].Metadata.CreationTimestamp
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

					Output = append(Output, name+"|"+readyCount+"|"+UpToDate+"|"+availableCount+"|"+age+"\n")
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
		yfile, _ := os.ReadFile(Namespaces_Path + List2Item + "/apps/deployments.yaml")
		MyDeployments := DEPLOYMENTS{}
		yaml.Unmarshal(yfile, &MyDeployments)

		for i := 0; i < len(MyDeployments.Items); i++ {
			if MyDeployments.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MyDeployments.Items[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "DeploymentConfig" && List6Item == "Info" {
		// Get projects deploymentconfigs "if exists"
		// Get projects deployments "if exists"
		yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/apps.openshift.io/deploymentconfigs.yaml")
		MyDCs := DEPLOYMENTCONFIGS{}
		yaml.Unmarshal(yfile, &MyDCs)

		if len(MyDCs.Items) > 0 {
			Output = []string{"NAME" + "|" + "REVISION" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "TRIGGERED BY" + "\n"}
			for i := 0; i < len(MyDCs.Items); i++ {
				if MyDCs.Items[i].Metadata.Name == List4Item {
					name := MyDCs.Items[i].Metadata.Name
					List4.AddItem(name, "", 0, nil)

					revision := strconv.Itoa(MyDCs.Items[i].Status.LatestVersion)

					desired := strconv.Itoa(MyDCs.Items[i].Spec.Replicas)

					current := strconv.Itoa(MyDCs.Items[i].Status.AvailableReplicas)

					triggers := MyDCs.Items[i].Spec.Triggers
					triggersType := ""
					for i := 0; i < len(triggers); i++ {
						if MyDCs.Items[i].Spec.Triggers[i].Type == "ImageChange" {
							triggersType = triggersType + "image" + "(" + MyDCs.Items[i].Spec.Triggers[i].ImageChangeParams.From.Name + ")"
						} else {
							triggersType = triggersType + "config"
						}
						if i != len(triggers)-1 {
							triggersType = triggersType + ","
						}
					}

					Output = append(Output, name+"|"+revision+"|"+desired+"|"+current+"|"+triggersType+"\n")
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
		yfile, _ := os.ReadFile(Namespaces_Path + List2Item + "/apps.openshift.io/deploymentconfigs.yaml")
		MyDCs := DEPLOYMENTCONFIGS{}
		yaml.Unmarshal(yfile, &MyDCs)

		for i := range MyDCs.Items {
			if MyDCs.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MyDCs.Items[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Daemonset" && List6Item == "Info" {
		// Get projects daemonsets "if exists"
		yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/apps/daemonsets.yaml")
		MyDSs := DAEMONSETS{}
		yaml.Unmarshal(yfile, &MyDSs)
		if len(MyDSs.Items) > 0 {
			Output = []string{"NAME" + "|" + "DESIRED" + "|" + "CURRENT" + "|" + "READY" + "|" + "UP-TO-DATE" + "|" + "AVAILABLE" + "|" + "NODE SELECTOR" + "|" + "AGE" + "\n"}
			for i := 0; i < len(MyDSs.Items); i++ {
				if MyDSs.Items[i].Metadata.Name == List4Item {
					now := time.Now().UTC()
					name := MyDSs.Items[i].Metadata.Name

					desired := strconv.Itoa(MyDSs.Items[i].Status.DesiredNumberScheduled)

					current := strconv.Itoa(MyDSs.Items[i].Status.CurrentNumberScheduled)

					ready := strconv.Itoa(MyDSs.Items[i].Status.NumberReady)

					uptodate := strconv.Itoa(MyDSs.Items[i].Status.UpdatedNumberScheduled)

					available := strconv.Itoa(MyDSs.Items[i].Status.NumberAvailable)

					// To be enhanced
					nodeSelectorS := ""
					nodeSelector := MyDSs.Items[i].Spec.Template.Spec.NodeSelector
					for key, value := range nodeSelector {
						nodeSelectorS += (key + ":" + value)
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

					Output = append(Output, name+"|"+desired+"|"+current+"|"+ready+"|"+uptodate+"|"+available+"|"+nodeSelectorS+"|"+age+"|"+"\n")
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
		yfile, _ := os.ReadFile(Namespaces_Path + List2Item + "/apps/daemonsets.yaml")
		MyDSs := DAEMONSETS{}
		yaml.Unmarshal(yfile, &MyDSs)
		for i := range MyDSs.Items {
			if MyDSs.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MyDSs.Items[i])
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
		yfile, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/core/services.yaml")
		MyServices := SERVICES{}
		yaml.Unmarshal(yfile, &MyServices)
		Output = []string{"NAME" + "|" + "TYPE" + "|" + "CLUSTER-IP" + "|" + "EXTERNAL-IP" + "|" + "PORT(S)" + "|" + "AGE" + "\n"}
		if len(MyServices.Items) > 0 {
			for i := 0; i < len(MyServices.Items); i++ {
				if MyServices.Items[i].Metadata.Name == List4Item {
					name := MyServices.Items[i].Metadata.Name

					Stype := MyServices.Items[i].Spec.Type

					clusterIP := MyServices.Items[i].Spec.ClusterIP

					externalName := MyServices.Items[i].Spec.ExternalName

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
					Output = append(Output, name+"|"+Stype+"|"+clusterIP+"|"+externalName+"|"+ports_proto+"|"+age+"|"+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		}
	} else if List1Item == "Projects" && List3Item == "Services" && List6Item == "YAML" {
		yfile, _ := os.ReadFile(Namespaces_Path + List2Item + "/core/services.yaml")
		MyServices := SERVICES{}
		yaml.Unmarshal(yfile, MyServices)
		for i := range MyServices.Items {
			if MyServices.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MyServices.Items[i])
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
		File, _ := ioutil.ReadFile(Namespaces_Path + List2Item + "/route.openshift.io/routes.yaml")
		MyRoutes := ROUTES{}
		yaml.Unmarshal(File, &MyRoutes)

		Items := MyRoutes.Items
		if len(Items) > 0 {
			for i := 0; i < len(Items); i++ {
				if Items[i].Metadata.Name == List4Item {

					name := Items[i].Metadata.Name

					host := Items[i].Spec.Host

					services := Items[i].Spec.To.Name

					port := Items[i].Spec.Port.TargetPort

					insecureEdgeTerminationPolicy := Items[i].Spec.TLS.InsecureEdgeTerminationPolicy
					termination := Items[i].Spec.TLS.Termination
					term := ""
					if termination == "" && insecureEdgeTerminationPolicy == "" {
						term = "None"
					} else {
						term = fmt.Sprintf("%v", termination) + "/" + fmt.Sprintf("%v", insecureEdgeTerminationPolicy)
					}
					wildcard := Items[i].Spec.WildcardPolicy
					CreationTime := Items[i].Metadata.CreationTimestamp
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

					Output = append(Output, name+"|"+host+"|"+""+"|"+services+"|"+port+"|"+term+"|"+wildcard+"|"+age+"|"+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput

		}
	} else if List1Item == "Projects" && List3Item == "Routes" && List6Item == "YAML" {
		File, _ = os.ReadFile(Namespaces_Path + List2Item + "/route.openshift.io/routes.yaml")
		MyRoutes := ROUTES{}
		yaml.Unmarshal(File, &MyRoutes)
		for i := range MyRoutes.Items {
			if MyRoutes.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MyRoutes.Items[i])
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
		File, _ = ioutil.ReadFile(Namespaces_Path + List2Item + "/image.openshift.io/imagestreams.yaml")
		MyImageStreams := IMAGESTREAMS{}
		yaml.Unmarshal(File, &MyImageStreams)
		if len(MyImageStreams.Items) > 0 {
			for i := 0; i < len(MyImageStreams.Items); i++ {
				if MyImageStreams.Items[i].Metadata.Name == List4Item {

					name := MyImageStreams.Items[i].Metadata.Name

					tagsS := ""
					if MyImageStreams.Items[i].Spec.Tags != nil {
						for x := 0; x < len(MyImageStreams.Items[i].Spec.Tags); x++ {
							tag_name := MyImageStreams.Items[i].Spec.Tags[x].Name
							tagsS = tagsS + tag_name
							if x != len(MyImageStreams.Items[i].Spec.Tags)-1 {
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

					Output = append(Output, List2Item+"|"+name+"|"+tagsS+"|"+age+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		}
	} else if List1Item == "Projects" && List3Item == "Image Stream" && List6Item == "YAML" {
		File, _ = os.ReadFile(Namespaces_Path + List2Item + "/image.openshift.io/imagestreams.yaml")
		MyImageStreams := IMAGESTREAMS{}
		yaml.Unmarshal(File, &MyImageStreams)

		for i := range MyImageStreams.Items {
			if MyImageStreams.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MyImageStreams.Items[i])
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
		namespaces, _ := ioutil.ReadDir(Namespaces_Path)
		if len(namespaces) > 0 {
			Output = []string{"NAMESPACE" + "|" + "NAME" + "|" + "STATUS" + "|" + "VOLUME" + "|" + "CAPACITY" + "|" + "ACCESS MODES" + "|" + "STORAGECLASS" + "|" + "AGE" + "\n"}
			for projectIndex := 0; projectIndex < len(namespaces); projectIndex++ {
				if _, err := os.Stat(Namespaces_Path + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml"); err == nil {
					File, _ = ioutil.ReadFile(Namespaces_Path + namespaces[projectIndex].Name() + "/core/persistentvolumeclaims.yaml")
					MyPVCS := PVCS{}
					yaml.Unmarshal(File, &MyPVCS)

					if len(MyPVCS.Items) > 0 {
						for i := 0; i < len(MyPVCS.Items); i++ {
							if MyPVCS.Items[i].Metadata.Name == List4Item {
								name := MyPVCS.Items[i].Metadata.Name

								status := MyPVCS.Items[i].Status.Phase

								volume := MyPVCS.Items[i].Spec.VolumeName

								capacity := MyPVCS.Items[i].Status.Capacity.Storage

								access := MyPVCS.Items[i].Status.AccessModes
								accessS := fmt.Sprintf("%v", access)
								accessS = strings.Replace(accessS, "[", "", -1)
								accessS = strings.Replace(accessS, "]", "", -1)

								storageClass := MyPVCS.Items[i].Spec.StorageClassName

								CreationTime := MyPVCS.Items[i].Metadata.CreationTimestamp
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
								Output = append(Output, namespaces[projectIndex].Name()+"|"+name+"|"+status+"|"+volume+"|"+capacity+"|"+accessS+"|"+storageClass+"|"+age+"\n")
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
		File, _ = os.ReadFile(Namespaces_Path + List2Item + "/core/persistentvolumeclaims.yaml")
		MyPVCS := PVCS{}
		yaml.Unmarshal(File, &MyPVCS)

		for i := range MyPVCS.Items {
			if MyPVCS.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MyPVCS.Items[i])
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
		File, _ = ioutil.ReadFile(Namespaces_Path + List2Item + "/core/configmaps.yaml")
		MyCMs := CONFIGMAPS{}
		yaml.Unmarshal(File, &MyCMs)

		if len(MyCMs.Items) > 0 {
			for i := 0; i < len(MyCMs.Items); i++ {
				if MyCMs.Items[i].Metadata.Name == List4Item {

					name := MyCMs.Items[i].Metadata.Name

					dataS := ""
					if MyCMs.Items[i].Data != nil {
						data := MyCMs.Items[i].Data
						dataN := len(data)
						dataS = fmt.Sprint(dataN)
					} else {
						dataS = "0"
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
					Output = append(Output, name+"|"+dataS+"|"+age+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput

		}
	} else if List1Item == "Projects" && List3Item == "ConfigMap" && List6Item == "YAML" {
		File, _ = os.ReadFile(Namespaces_Path + List2Item + "/core/configmaps.yaml")
		MyCMs := CONFIGMAPS{}
		yaml.Unmarshal(File, &MyCMs)
		for i := range MyCMs.Items {
			if MyCMs.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MyCMs.Items[i])
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
		File, _ = ioutil.ReadFile(Namespaces_Path + List2Item + "/core/secrets.yaml")
		MySecrets := SECRETS{}
		yaml.Unmarshal(File, &MySecrets)

		if len(MySecrets.Items) > 0 {
			for i := 0; i < len(MySecrets.Items); i++ {
				if MySecrets.Items[i].Metadata.Name == List4Item {
					name := MySecrets.Items[i].Metadata.Name

					dataS := ""
					if MySecrets.Items[i].Data != nil {
						data := MySecrets.Items[i].Data
						dataN := len(data)
						dataS = fmt.Sprint(dataN)
					} else {
						dataS = "0"
					}

					type_key := MySecrets.Items[i].Type

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

					Output = append(Output, name+"|"+type_key+"|"+dataS+"|"+age+"\n")
				}
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput

		}
	} else if List1Item == "Projects" && List3Item == "Secrets" && List6Item == "YAML" {
		File, _ = os.ReadFile(Namespaces_Path + List2Item + "/core/secrets.yaml")
		MySecrets := SECRETS{}
		yaml.Unmarshal(File, &MySecrets)
		for i := range MySecrets.Items {
			if MySecrets.Items[i].Metadata.Name == List4Item {
				yaml, _ := yaml.Marshal(MySecrets.Items[i])
				TextView.SetText(string(yaml))
			}
		}
		TextView.ScrollToBeginning()
	} else if List1Item == "Projects" && List3Item == "Subscriptions" && List6Item == "Info" {
		//TBA
	} else if List1Item == "Projects" && List3Item == "Supscriptions" && List6Item == "YAML" {
		// yfile, _ := os.ReadFile(Namespaces_Path + List2Item + "/apps/deployments.yaml")
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
		// yfile, _ := os.ReadFile(Namespaces_Path + List2Item + "/apps/deployments.yaml")
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
		File, _ = ioutil.ReadFile(Nodes_Path + List2Item + ".yaml")
		MyNode := NODE{}
		yaml.Unmarshal(File, &MyNode)
		if List6Item == "Metadata" {
			Metadta, _ := yaml.Marshal(MyNode.Metadata)
			MetadtaS := Colors.Orange + "Metadta:\n" + Colors.White + string(Metadta)
			TextView.Clear()
			TextView.SetWrap(true)
			TextView.SetText(MetadtaS)
			TextView.ScrollToBeginning()
		} else if List6Item == "Spec" {
			Spec, _ := yaml.Marshal(MyNode.Spec)
			SpecS := Colors.Orange + "Spec:\n" + Colors.White + string(Spec)
			TextView.Clear()
			TextView.SetText(SpecS)
			TextView.ScrollToBeginning()
		} else if List6Item == "Status" {
			Status, _ := yaml.Marshal(MyNode.Status)
			StatusS := Colors.Orange + "Status:\n" + Colors.White + string(Status)
			TextView.Clear()
			TextView.SetText(StatusS)
			TextView.ScrollToBeginning()
		} else if List6Item == "HW Specs" {
			HWSpec := ""
			Addresses, _ := yaml.Marshal(MyNode.Status.Addresses)
			Allocatable, _ := yaml.Marshal(MyNode.Status.Allocatable)
			Capacity, _ := yaml.Marshal(MyNode.Status.Capacity)
			HWSpec += Colors.Orange + "Addresses:\n" + Colors.White + string(Addresses) + Colors.Orange + "Allocatable:\n" + Colors.White + string(Allocatable) + Colors.Orange + "Capacity:\n" + Colors.White + string(Capacity)
			TextView.Clear()
			TextView.SetText(HWSpec)
			TextView.ScrollToBeginning()
		} else if List6Item == "Images" {
			Images, _ := yaml.Marshal(MyNode.Status.Images)
			ImagesS := Colors.Orange + "Images:\n" + Colors.White + string(Images)
			TextView.Clear()
			TextView.SetText(ImagesS)
			TextView.ScrollToBeginning()
		} else if List6Item == "nodeInfo" {
			NodeInfo, _ := yaml.Marshal(MyNode.Status.NodeInfo)
			NodeInfoS := Colors.Orange + "NodeInfo:\n" + Colors.White + string(NodeInfo)
			TextView.Clear()
			TextView.SetText(NodeInfoS)
			TextView.ScrollToBeginning()
		}
	} else if List1Item == "ETCD" && List2Item == "Endpoint Health" && List6Item == "JSON" {
		File, _ = ioutil.ReadFile(ETCD_Path + "endpoint_health.json")
		MyETCD_EP_H := ETCD_EP_H{}
		json.Unmarshal(File, &MyETCD_EP_H)
		File, _ = json.MarshalIndent(&MyETCD_EP_H, "", "\t")
		TextView.SetText(string(File))
	} else if List1Item == "ETCD" && List2Item == "Endpoint Status" && List6Item == "JSON" {
		File, _ = ioutil.ReadFile(ETCD_Path + "endpoint_status.json")
		MyETCD_EP_S := ETCD_EP_S{}
		json.Unmarshal(File, &MyETCD_EP_S)
		File, _ = json.MarshalIndent(&MyETCD_EP_S, "", "\t")
		TextView.SetText(string(File))
	} else if List1Item == "ETCD" && List2Item == "Member List" && List6Item == "JSON" {
		File, _ = ioutil.ReadFile(ETCD_Path + "member_list.json")
		MyETCD_M_L := ETCD_M_L{}
		json.Unmarshal(File, &MyETCD_M_L)
		File, _ = json.MarshalIndent(&MyETCD_M_L, "", "\t")
		TextView.SetText(string(File))
	}
}
