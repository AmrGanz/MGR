package functions

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
	if List1Item == "Configurations" {
		file, err := ioutil.ReadFile(BasePath + "cluster-scoped-resources/config.openshift.io/" + List2Item + ".yaml")
		if err != nil {
			fmt.Print(err)
		} else {
			TextView.SetText(string(file))
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
			GetNodesInfo("All Nodes", "Summary")
		} else if List2Item == "Nodes Details" {
			GetNodesInfo("All Nodes", "Details")
		} else {
			GetNodesInfo(List2Item, "Summary")
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
			now := time.Now().UTC()
			Files, _ = ioutil.ReadDir(BasePath + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/")
			Output := []string{Colors.Yellow + "NAME" + "|" + "AGE" + "|" + "SIGNERNAME" + "|" + "REQUESTOR" + "|" + "CONDITION" + Colors.White}
			if len(Files) == 0 {
				TextView.SetText("No available data about CSR")
			} else {
				for _, File := range Files {
					var MyCSR = CSR{}
					yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/" + File.Name())
					yaml.Unmarshal(yfile, &MyCSR)

					name := MyCSR.Metadata.Name

					// Not accurate yet!!!
					CreationTime := MyCSR.Metadata.CreationTimestamp
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

					signature := MyCSR.Spec.SignerName
					requester := MyCSR.Spec.Username
					condition := ""

					if MyCSR.Status.Certificate == "" {
						condition = Colors.Red + "Pending" + Colors.White
					} else {
						condition = Colors.Green + "Approved,Issued" + Colors.White
					}
					Output = append(Output, Colors.White+name+"|"+age+"|"+signature+"|"+requester+"|"+condition+"|"+Colors.White)
				}
				FormatedOutput := columnize.SimpleFormat(Output)
				TextView.SetText(FormatedOutput)
				TextView.ScrollToBeginning()
				TextViewData = FormatedOutput
			}

		} else {
			File, _ = ioutil.ReadFile(BasePath + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/" + List2Item + ".yaml")
			TextView.SetText(string(File))
			TextView.ScrollToBeginning()
			TextViewData = TextView.GetText(false)
		}
	}
}
