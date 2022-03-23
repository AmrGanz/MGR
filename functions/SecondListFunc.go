package functions

import (
	"fmt"
	"io/ioutil"

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
	if List1Item == "Configurations" {
		file, err := ioutil.ReadFile(MG_Path + "cluster-scoped-resources/config.openshift.io/" + List2Item + ".yaml")
		if err != nil {
			fmt.Print(err)
		} else {
			TextView.SetText(string(file))
		}
	} else if List1Item == "OCP Version" {
		File, _ = ioutil.ReadFile(MG_Path + "cluster-scoped-resources/config.openshift.io/clusterversions/version.yaml")
		if List2Item == "YAML" {
			TextView.SetText(string(File))
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		} else if List2Item == "Cluster Update Details" {
			Output = []string{}
			//////////////////////
			// Get cluster version
			MyCV := CLUSTERVERSION{}
			yaml.Unmarshal(File, &MyCV)
			ClusterVersion := MyCV.Spec.DesiredUpdate.Version
			UpgradeChannel := MyCV.Spec.Channel
			ClusterID := MyCV.Spec.ClusterID
			Output = append(Output, "Cluster desired version: "+Colors.Green+ClusterVersion+Colors.White)
			Output = append(Output, "")
			Output = append(Output, "Upgarde channel: "+Colors.Green+UpgradeChannel+Colors.White)
			Output = append(Output, "")
			Output = append(Output, "ClusterID: "+Colors.Green+ClusterID+Colors.White)
			Output = append(Output, "")
			Conditions := MyCV.Status.Conditions
			for x := range Conditions {
				if MyCV.Status.Conditions[x].Type == "Available" && MyCV.Status.Conditions[x].Status == "True" {
					Output = append(Output, "Cluster update status: "+Colors.Green+"cluster is updated to "+ClusterVersion+Colors.White)
					Output = append(Output, "")
				} else if MyCV.Status.Conditions[x].Type == "Available" && MyCV.Status.Conditions[x].Status == "False" {
					Output = append(Output, "Cluster update status: "+Colors.Red+"cluster is not fully updated to "+ClusterVersion+Colors.White)
					Output = append(Output, "")
				}
			}
			//////////////////////
			// Get update path
			Output = append(Output, "Upgrade Path:")
			Output = append(Output, Colors.Yellow+"Version"+"|"+"Completion date/time"+Colors.White)
			for x := len(MyCV.Status.History) - 1; x > -1; x-- {
				Output = append(Output, Colors.White+MyCV.Status.History[x].Version+"|"+MyCV.Status.History[x].CompletionTime.String()+Colors.White)
			}
			FormatedOutput := columnize.SimpleFormat(Output)
			TextView.SetText(FormatedOutput)
			TextView.ScrollToBeginning()
			TextViewData = FormatedOutput
		}

	} else if List1Item == "Projects" {
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
				AddItem("Operators", "", 0, nil).
				AddItem("Install Plans", "", 0, nil)
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
				AddItem("Operators", "", 0, nil).
				AddItem("Install Plans", "", 0, nil)
		}
	} else if List1Item == "Nodes" {
		if List2Item == "Nodes Summary" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			Files, Err = ioutil.ReadDir(Nodes_Path)
			if Err != nil {
				TextView.SetText(Err.Error())
			} else {
				Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + Colors.White}
				for i := range Files {
					MyNode := NODE{}
					File, _ = ioutil.ReadFile(Nodes_Path + Files[i].Name())
					yaml.Unmarshal(File, &MyNode)
					name, status, roles, age, version, _, _, _, _, _, _ := GetNodeDetails(MyNode)
					Output = append(Output, Colors.White+name+"|"+status+"|"+roles+"|"+age+"|"+version+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}
		} else if List2Item == "Nodes Details" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			Files, Err = ioutil.ReadDir(Nodes_Path)
			if Err != nil {
				TextView.SetText("Couldn't read Nodes directory's contents")
			} else {
				Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "|" + "INTERNAL-IP" + "|" + "EXTERNAL-IP" + "|" + "OS-IMAGE" + "|" + "KERNEL-VERSION" + "|" + "CONTAINER-RUNTIME" + Colors.White + "\n"}
				for i := range Files {
					MyNode := NODE{}
					File, _ = ioutil.ReadFile(Nodes_Path + Files[i].Name())
					yaml.Unmarshal(File, &MyNode)
					name, status, roles, age, version, internalIP, externalIP, osImage, kernelVersion, contRuntime, _ := GetNodeDetails(MyNode)
					Output = append(Output, Colors.White+name+"|"+status+"|"+roles+"|"+age+"|"+version+"|"+internalIP+"|"+externalIP+"|"+osImage+"|"+kernelVersion+"|"+contRuntime+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}
		} else if List2Item == "Nodes Labels" {
			// Cleaning TextView and TextViewData
			TextView.Clear()
			TextViewData = ""
			Files, Err = ioutil.ReadDir(Nodes_Path)
			if Err != nil {
				TextView.SetText("Couldn't read Nodes directory's contents")
			} else {
				Output = []string{Colors.Yellow + "NAME" + "|" + "LABELS" + Colors.White}
				for i := range Files {
					MyNode := NODE{}
					File, _ = ioutil.ReadFile(Nodes_Path + Files[i].Name())
					yaml.Unmarshal(File, &MyNode)
					name, _, _, _, _, _, _, _, _, _, AllLabels := GetNodeDetails(MyNode)
					Output = append(Output, Colors.White+name+"|"+AllLabels+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}
		} else {
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
			List3.
				AddItem("Summary", "", 0, nil).
				AddItem("Details", "", 0, nil).
				AddItem("YAML", "", 0, nil)
		}
	} else if List1Item == "Operators" {
		List3.SetTitle("Info")
		List3.
			AddItem("YAML", "", 0, nil).
			AddItem("Summary", "", 0, nil)
	} else if List1Item == "Installed Operators" {
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
	} else if List1Item == "CSR" {
		if List2Item == "All Certificate Signing Requests" {
			GetCSRInfo()
		} else {
			File, _ = ioutil.ReadFile(MG_Path + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/" + List2Item + ".yaml")
			TextView.SetText(string(File))
			TextView.ScrollToBeginning()
			TextViewData = TextView.GetText(false)
		}
	}
}
