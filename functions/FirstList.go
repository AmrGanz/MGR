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
	ActivePathBox.SetText(List1Item + " > ")
	if list_item_name == "Navigate" {
		// App.SetInputCapture(KeyboardKeys).
		// 	SetRoot(TableGrid, true).
		// 	EnableMouse(true)

		// App.SetInputCapture(KeyboardKeys).
		// SetRoot(MainGrid, true).
		// EnableMouse(true)
	} else if list_item_name == "Summary" {

		// Clear Output string array
		Output = nil
		//////////////////////
		// Get cluster version
		File, _ = ioutil.ReadFile(Version_Path)
		MyCV := CLUSTERVERSION{}
		yaml.Unmarshal(File, &MyCV)
		ClusterVersion := MyCV.Spec.DesiredUpdate.Version
		UpgradeChannel := MyCV.Spec.Channel
		ClusterID := MyCV.Spec.ClusterID
		if ClusterVersion != "" {
			Output = append(Output, "Cluster desired version: "+Colors.Yellow+ClusterVersion+Colors.White)
			Output = append(Output, "")
		}
		if UpgradeChannel != "" {
			Output = append(Output, "Upgrade channel: "+Colors.Yellow+UpgradeChannel+Colors.White)
			Output = append(Output, "")
		}
		if ClusterID != "" {
			Output = append(Output, "ClusterID: "+Colors.Yellow+ClusterID+Colors.White)
			Output = append(Output, "")
		}

		Conditions := MyCV.Status.Conditions
		for x := range Conditions {
			if MyCV.Status.Conditions[x].Type == "Available" && MyCV.Status.Conditions[x].Status == "True" {
				Output = append(Output, "Cluster update status: "+Colors.Yellow+"cluster is updated to "+ClusterVersion+Colors.White)
				Output = append(Output, "")
			} else if MyCV.Status.Conditions[x].Type == "Available" && MyCV.Status.Conditions[x].Status == "False" {
				Output = append(Output, "Cluster update status: "+Colors.Red+"cluster is not fully updated to "+ClusterVersion+Colors.White)
				Output = append(Output, "")
			}
		}

		//////////////////////
		// Get update path
		UpdatePath := ""
		for x := len(MyCV.Status.History) - 1; x > -1; x-- {
			if x > 0 {
				UpdatePath = UpdatePath + MyCV.Status.History[x].Version + " > "
			} else {
				UpdatePath = UpdatePath + MyCV.Status.History[x].Version
			}

		}
		if UpdatePath != "" {
			Output = append(Output, "Update Path: "+Colors.Yellow+UpdatePath+Colors.White)
			Output = append(Output, "")
		}

		//////////////////////
		// Get nodes status
		Files, _ = ioutil.ReadDir(Nodes_Path)
		NodeDownCount := 0
		for i := range Files {
			MyNode := NODE{}
			File, _ = ioutil.ReadFile(Nodes_Path + Files[i].Name())
			yaml.Unmarshal(File, &MyNode)
			conditions := MyNode.Status.Conditions
			for i := 0; i < len(conditions); i++ {
				if MyNode.Status.Conditions[i].Type == "Ready" {
					if MyNode.Status.Conditions[i].Status != "True" {
						NodeDownCount++
					}
				}
			}
		}
		if NodeDownCount > 0 {
			Output = append(Output, "Nodes status: "+Colors.Red+fmt.Sprint(NodeDownCount)+" cluster node(s) are NotReady"+Colors.White)
			Output = append(Output, "")
		} else {
			Output = append(Output, "Nodes status: "+Colors.Yellow+"All of the cluster nodes are Ready"+Colors.White)
			Output = append(Output, "")
		}

		//////////////////////
		// Get Operators status

		if _, err := os.Stat(All_Operators_Path); err == nil {
			File, _ = ioutil.ReadFile(All_Operators_Path)
			MyOperators := CLUSTEROPERATORS{}
			yaml.Unmarshal(File, &MyOperators)
			OperatorsDownCount := 0
			for i := range MyOperators.Items {
				Conditions := MyOperators.Items[i].Status.Conditions
				for x := range Conditions {
					if MyOperators.Items[i].Status.Conditions[x].Type == "Degraded" && MyOperators.Items[i].Status.Conditions[x].Status == "True" {
						OperatorsDownCount++
						break
					} else if MyOperators.Items[i].Status.Conditions[x].Type == "Progressing" && MyOperators.Items[i].Status.Conditions[x].Status == "True" {
						OperatorsDownCount++
						break
					} else if MyOperators.Items[i].Status.Conditions[x].Type == "Available" && MyOperators.Items[i].Status.Conditions[x].Status == "False" {
						OperatorsDownCount++
						break
					}
				}
			}
			if OperatorsDownCount > 0 {
				Output = append(Output, "Operators status: "+Colors.Red+fmt.Sprint(OperatorsDownCount)+" cluster operator(s) not fully Active"+Colors.White)
				Output = append(Output, "")
			} else {
				Output = append(Output, "Operators status: "+Colors.Yellow+"All of the cluster operators are Active"+Colors.White)
				Output = append(Output, "")
			}
		}

		//////////////////////
		// Get MCP status
		if _, err := os.Stat(MCP_Path); err == nil {
			Files, _ = ioutil.ReadDir(MCP_Path)
			MCPDownCount := 0
			for i := range Files {
				MyMCP := MCP{}
				File, _ = ioutil.ReadFile(MCP_Path + Files[i].Name())
				yaml.Unmarshal(File, &MyMCP)
				Conditions := MyMCP.Status.Conditions
				for x := range Conditions {
					if MyMCP.Status.Conditions[x].Type == "Updated" && MyMCP.Status.Conditions[x].Status == "False" {
						MCPDownCount++
						break
					} else if MyMCP.Status.Conditions[x].Type == "Updating" && MyMCP.Status.Conditions[x].Status == "True" {
						MCPDownCount++
						break
					} else if MyMCP.Status.Conditions[x].Type == "Degraded" && MyMCP.Status.Conditions[x].Status == "True" {
						MCPDownCount++
						break
					}
				}
			}
			if MCPDownCount > 0 {
				Output = append(Output, "Machine Config Pools Status: "+Colors.Red+fmt.Sprint(MCPDownCount)+" MCP(s) not fully Updated"+Colors.White)
				Output = append(Output, "")
			} else {
				Output = append(Output, "Machine Config Pools Status: "+Colors.Yellow+"All MCPs are Updated"+Colors.White)
				Output = append(Output, "")
			}
		}

		//////////////////////
		// Get Pending CSR count
		if _, err := os.Stat(CSR_Path); err == nil {
			Files, _ = ioutil.ReadDir(CSR_Path)
			PendingCSRCount := 0
			for i := range Files {
				var MyCSR = CSR{}
				File, _ = ioutil.ReadFile(CSR_Path + Files[i].Name())
				yaml.Unmarshal(File, &MyCSR)
				if MyCSR.Status.Certificate == "" {
					PendingCSRCount++
				}
			}
			if PendingCSRCount > 0 {
				Output = append(Output, "CSR status: "+Colors.Red+fmt.Sprint(PendingCSRCount)+" CSR(s) not approved yet"+Colors.White)
				Output = append(Output, "")
			} else {
				Output = append(Output, "CSR status: "+Colors.Green+"All CSRs are Approved"+Colors.White)
				Output = append(Output, "")
			}
		}

		//////////////////////
		// Get failed pods count

		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput

	} else if list_item_name == "OCP Version" {
		List2.SetTitle("Cluster Version Detail")
		List2.
			AddItem("YAML", "", 0, nil).
			AddItem("Cluster Update Details", "", 0, nil)
		File, _ = ioutil.ReadFile(Version_Path)
		TextView.SetText(string(File))
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput

	} else if list_item_name == "Configurations" {
		List2.SetTitle("Cluster Configurations")
		files, _ := ioutil.ReadDir(Configurations_Path)
		for i := range files {
			if !files[i].IsDir() {
				List2.AddItem(strings.Split(files[i].Name(), ".yaml")[0], "", 0, nil)
			}
		}

	} else if list_item_name == "Projects" {
		List2.SetTitle("Projects")
		namespaces, _ := ioutil.ReadDir(Namespaces_Path)
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
		files, _ := ioutil.ReadDir(Nodes_Path)
		List2.
			AddItem("Nodes Summary", "", 0, nil).
			AddItem("Nodes Details", "", 0, nil).
			AddItem("Nodes Labels", "", 0, nil)
		// Populate List2 with nodes names
		for _, node := range files {
			nodeName := strings.Split(node.Name(), ".yaml")
			List2.AddItem(nodeName[0], "", 0, nil)
		}
	} else if list_item_name == "Operators" {
		List2.SetTitle("Operators")
		// Get cluster version
		File, _ = ioutil.ReadFile(Version_Path)
		MyCV := CLUSTERVERSION{}
		yaml.Unmarshal(File, &MyCV)
		ClusterVersion := MyCV.Spec.DesiredUpdate.Version

		File, _ = ioutil.ReadFile(All_Operators_Path)
		Output := []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "VERSION" + Colors.Yellow + "|" + Colors.Yellow + "AVAILABLE" + Colors.Yellow + "|" + Colors.Yellow + "PROGRESSING" + Colors.Yellow + "|" + Colors.Yellow + "DEGRADED" + Colors.Yellow + "|" + "SINCE" + Colors.White}
		MyOperators := CLUSTEROPERATORS{}
		yaml.Unmarshal(File, &MyOperators)
		items := MyOperators.Items
		for i := range items {
			name := items[i].Metadata.Name
			List2.AddItem(name, "", 0, nil)

			versions := items[i].Status.Versions
			versionsS := ""
			for i := range versions {
				if versions[i].Name == "operator" {
					versionsS = versions[i].Version
					if versionsS == ClusterVersion {
						versionsS = Colors.White + versionsS + Colors.White
					} else {
						versionsS = Colors.Red + versionsS + Colors.White
					}
				}
			}
			status := items[i].Status.Conditions
			availableS := ""
			progressingS := ""
			degradedS := ""
			availableSince := ""
			for i := range status {
				if status[i].Type == "Available" {
					if status[i].Status == "True" {
						availableS = Colors.Green + "True" + Colors.White
					} else {
						availableS = Colors.Red + "False" + Colors.White
					}

					now := time.Now().UTC()
					statusTime := status[i].LastTransitionTime
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

				} else if status[i].Type == "Progressing" {
					if status[i].Status == "True" {
						progressingS = Colors.Red + "True" + Colors.White
					} else {
						progressingS = Colors.Green + "False" + Colors.White
					}
				} else if status[i].Type == "Degraded" {
					if status[i].Status == "True" {
						degradedS = Colors.Red + "True" + Colors.White
					} else {
						degradedS = Colors.Green + "False" + Colors.White
					}
				}

			}
			Output = append(Output, Colors.White+name+"|"+versionsS+"|"+availableS+"|"+progressingS+"|"+degradedS+"|"+availableSince+Colors.White)
		}
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput
	} else if list_item_name == "Installed Operators" {
		TextView.Clear()
		TextViewData = ""
		List2.SetTitle("Installed Operators")
		// Get installed operators file
		Output = []string{Colors.Yellow + "NAME" + "|" + "AGE" + Colors.White}
		Files, _ = ioutil.ReadDir(InstalledOperators_Path)
		for i := range Files {
			List2.AddItem(strings.Split(Files[i].Name(), ".yaml")[0], "", 0, nil)
			File, _ = ioutil.ReadFile(InstalledOperators_Path + Files[i].Name())
			MyOperator := CLUSTEROPERATOR{}
			yaml.Unmarshal(File, &MyOperator)
			name := MyOperator.Metadata.Name
			age := GetAge(MyOperator.Metadata.CreationTimestamp)
			Output = append(Output, Colors.White+name+"|"+age+Colors.White)
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
		Files, _ = ioutil.ReadDir(MCP_Path)
		GetAllMCPInfo(Files)

	} else if list_item_name == "MC" {
		List2.SetTitle("MC")
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		now := time.Now().UTC()
		Output := []string{Colors.Yellow + "NAME" + "|" + "GENERATEDBYCONTROLLER" + "|" + "IGNITIONVERSION" + "|" + "AGE" + Colors.White + "\n"}
		Files, _ = ioutil.ReadDir(MC_Path)
		for _, file := range Files {
			File, _ = ioutil.ReadFile(MC_Path + file.Name())

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
			Output = append(Output, Colors.White+name+"|"+generatedBy+"|"+ignitionVersion+"|"+age+Colors.White+"\n")
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
		Files, _ = ioutil.ReadDir(PV_Path)
		for _, f := range Files {
			File, _ = ioutil.ReadFile(PV_Path + f.Name())
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
		}
		FormatedOutput := columnize.SimpleFormat(Output)
		TextView.SetText(FormatedOutput)
		TextView.ScrollToBeginning()
		TextViewData = FormatedOutput

	} else if list_item_name == "CSR" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		Files, _ = ioutil.ReadDir(CSR_Path)
		List2.SetTitle("CSR")
		GetCSRInfo()
		List2.AddItem("All Certificate Signing Requests", "", 0, nil)
		for _, File := range Files {
			List2.AddItem(strings.Split(File.Name(), ".yaml")[0], "", 0, nil)
		}
	} else if list_item_name == Colors.Yellow+"Logging"+Colors.White {
		TextView.SetText("TBA")
	} else if list_item_name == "ETCD" {
		// Cleaning TextView and TextViewData
		TextView.Clear()
		TextViewData = ""
		List2.SetTitle("ETCD")

		List2.
			AddItem("Alarm list", "", 0, nil).
			AddItem("Endpoint Health", "", 0, nil).
			AddItem("Endpoint Status", "", 0, nil).
			AddItem("Member List", "", 0, nil)

	}
}
