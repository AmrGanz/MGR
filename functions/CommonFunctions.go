package functions

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/ryanuber/columnize"
	"gopkg.in/yaml.v2"
)

func GetNodesInfo(NodeName string, Flag string) {
	// Cleaning TextView and TextViewData
	TextView.Clear()
	TextViewData = ""
	if NodeName == "All Nodes" {
		Files, Err = ioutil.ReadDir(MG_Path + "cluster-scoped-resources/core/nodes/")
		if Err != nil {
			TextView.SetText(Err.Error())
		} else {
			if Flag == "Summary" {
				Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + Colors.White}
				for i := range Files {
					ReadNodeYaml(Files[i].Name(), "Summary")
				}
			} else if Flag == "Details" {
				Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "|" + "INTERNAL-IP" + "|" + "EXTERNAL-IP" + "|" + "OS-IMAGE" + "|" + "KERNEL-VERSION" + "|" + "CONTAINER-RUNTIME" + Colors.White + "\n"}
				for i := range Files {
					ReadNodeYaml(Files[i].Name(), "Details")
				}
			} else if Flag == "Labels" {
				Output = []string{Colors.Yellow + "NAME" + "|" + "LABELS" + Colors.White}
				for i := range Files {
					ReadNodeYaml(Files[i].Name(), "Labels")
				}
			}
		}
	} else {
		if Flag == "Summary" {
			Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + Colors.White + "\n"}
			ReadNodeYaml(List2Item+".yaml", "Summary")
		} else if Flag == "Details" {
			Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "|" + "INTERNAL-IP" + "|" + "EXTERNAL-IP" + "|" + "OS-IMAGE" + "|" + "KERNEL-VERSION" + "|" + "CONTAINER-RUNTIME" + Colors.White + "\n"}
			ReadNodeYaml(List2Item+".yaml", "Details")
		}
	}
	FormatedOutput := columnize.SimpleFormat(Output)
	TextView.SetText(FormatedOutput)
	TextView.ScrollToBeginning()
	TextViewData = FormatedOutput
}

func ReadNodeYaml(NodeFileName string, Flag string) {
	var MyNode = NODE{}
	now := time.Now().UTC()
	File, _ := ioutil.ReadFile(MG_Path + "cluster-scoped-resources/core/nodes/" + NodeFileName)

	yaml.Unmarshal(File, &MyNode)
	MyNode_Public = MyNode
	name := MyNode.Metadata.Name
	Labels := ""
	if Flag == "Labels" {
		i := len(MyNode.Metadata.Labels)
		for key, value := range MyNode.Metadata.Labels {
			if i > 1 {
				Labels += key + "=" + value + ","
				i--
			} else {
				Labels += key + value
			}

		}
		Output = append(Output, Colors.White+name+"|"+Labels+Colors.White)

	} else {
		conditions := MyNode.Status.Conditions
		statusS := ""
		for i := 0; i < len(conditions); i++ {
			if MyNode.Status.Conditions[i].Type == "Ready" {
				if MyNode.Status.Conditions[i].Status == "True" {
					if MyNode.Spec.Unschedulable {
						statusS = Colors.Orange + "Ready,SchedulingDisabled" + Colors.White
					} else {
						statusS = Colors.Green + "Ready" + Colors.White
					}

				} else {
					if MyNode.Spec.Unschedulable {
						statusS = Colors.Red + "NotReady,SchedulingDisabled" + Colors.White
					} else {
						statusS = Colors.Red + "NotReady" + Colors.White
					}
				}
			}
		}
		roles := ""
		Labels := MyNode.Metadata.Labels
		for key := range Labels {
			if strings.Contains(key, "node-role.kubernetes.io") {
				roles += strings.Split(key, "/")[1] + " "
			}
		}

		CreationTime := MyNode.Metadata.CreationTimestamp
		diff := now.Sub(CreationTime).Seconds()
		diffInt := int(diff)
		seconds := strconv.Itoa((diffInt % 60))
		minutes := strconv.Itoa((diffInt / 60) % 60)
		hours := strconv.Itoa((diffInt / 360) % 24)
		days := strconv.Itoa((diffInt / 86400))
		age := ""
		if days != "0" {
			age = days + "d" + hours + "h"
		} else if days == "0" && hours != "" {
			age = hours + "h" + minutes + "m"
		} else if hours == "0" {
			age = minutes + "m" + seconds + "s"
		}

		versionS := ""
		version := MyNode.Status.NodeInfo.KubeletVersion
		if Flag == "Summary" {
			versionS = fmt.Sprintf("%v", version)
			Output = append(Output, Colors.White+name+"|"+statusS+"|"+roles+"|"+age+"|"+versionS+Colors.White+"\n")
		} else if Flag == "Details" {
			versionS = fmt.Sprintf("%v", version)

			Addresses := MyNode.Status.Addresses
			internalIP := ""
			externalIP := ""
			for i := range Addresses {
				if Addresses[i].Type == "InternalIP" {
					internalIP = Addresses[i].Address
				} else if Addresses[i].Type == "ExternalIP" {
					externalIP = Addresses[i].Address

				}
			}

			osImage := MyNode.Status.NodeInfo.OsImage

			kernelVersion := MyNode.Status.NodeInfo.KernelVersion

			contRuntime := MyNode.Status.NodeInfo.ContainerRuntimeVersion

			Output = append(Output, Colors.White+name+"|"+statusS+"|"+roles+"|"+age+"|"+versionS+"|"+internalIP+"|"+externalIP+"|"+osImage+"|"+kernelVersion+"|"+contRuntime+Colors.White+"\n")
		}
	}

}

func GetAllMCPInfo(mcp_files []fs.FileInfo) {
	// Cleaning TextView and TextViewData
	TextView.Clear()
	TextViewData = ""
	Output := []string{Colors.Yellow + "NAME" + "|" + "CONFIG" + "|" + Colors.Yellow + "UPDATED" + Colors.Yellow + "|" + Colors.Yellow + "UPDATING" + Colors.Yellow + "|" + Colors.Yellow + "DEGRADED" + Colors.Yellow + "|" + "MACHINECOUNT" + "|" + "READYMACHINECOUNT" + "|" + "UPDATEDMACHINECOUNT" + "|" + "DEGRADEDMACHINECOUNT" + "|" + "AGE" + Colors.White}

	for _, mcp := range Files {
		File, _ = ioutil.ReadFile(MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigpools/" + mcp.Name())

		MyMCP := MCP{}
		yaml.Unmarshal(File, &MyMCP)

		name := MyMCP.Metadata.Name
		List2.AddItem(name, "", 0, nil)

		config := MyMCP.Status.Configuration.Name
		configS := fmt.Sprintf("%v", config)

		conditions := MyMCP.Status.Conditions
		updatedS := ""
		updatingS := ""
		degradedS := ""
		for i := range conditions {
			if conditions[i].Type == "Updated" {
				if conditions[i].Status == "True" {
					updatedS = Colors.Green + "True" + Colors.White
				} else {
					updatedS = Colors.Red + "False" + Colors.White
				}

			} else if conditions[i].Type == "Updating" {
				if conditions[i].Status == "True" {
					updatingS = Colors.Red + "True" + Colors.White
				} else {
					updatingS = Colors.Green + "False" + Colors.White
				}
			} else if conditions[i].Type == "Degraded" {
				if conditions[i].Status == "True" {
					degradedS = Colors.Red + "True" + Colors.White
				} else {
					degradedS = Colors.Green + "False" + Colors.White
				}
			}

		}
		machineCount := MyMCP.Status.MachineCount
		machineCountS := fmt.Sprint(machineCount)

		machineReady := MyMCP.Status.ReadyMachineCount
		machineReadyS := fmt.Sprint(machineReady)

		machineUpdated := MyMCP.Status.UpdatedMachineCount
		machineUpdatedS := fmt.Sprint(machineUpdated)

		machineDegraded := MyMCP.Status.DegradedMachineCount
		machineDegradedS := fmt.Sprint(machineDegraded)

		now := time.Now().UTC()
		CreationTime := MyMCP.Metadata.CreationTimestamp
		diff := now.Sub(CreationTime).Seconds()
		diffInt := int(diff)
		seconds := strconv.Itoa((diffInt % 60))
		minutes := strconv.Itoa((diffInt / 60) % 60)
		hours := strconv.Itoa((diffInt / 360) % 24)
		days := strconv.Itoa((diffInt / 86400))
		age := ""
		if days != "0" {
			age = days + "d" + hours + "h"
		} else if days == "0" && hours != "" {
			age = hours + "h" + minutes + "m"
		} else if hours == "0" {
			age = minutes + "m" + seconds + "s"
		}
		Output = append(Output, Colors.White+name+"|"+configS+"|"+updatedS+"|"+updatingS+"|"+degradedS+"|"+machineCountS+"|"+machineReadyS+"|"+machineUpdatedS+"|"+machineDegradedS+"|"+age+Colors.White)
	}
	FormatedOutput := columnize.SimpleFormat(Output)
	TextView.SetText(FormatedOutput)
	TextView.ScrollToBeginning()
	TextViewData = FormatedOutput
}

func GetMCPInfo(mcp_file []byte) {
	// Cleaning TextView and TextViewData
	TextView.Clear()
	TextViewData = ""
	Output := []string{Colors.Yellow + "NAME" + "|" + "CONFIG" + "|" + Colors.Yellow + "UPDATED" + Colors.Yellow + "|" + Colors.Yellow + "UPDATING" + Colors.Yellow + "|" + Colors.Yellow + "DEGRADED" + Colors.Yellow + "|" + "MACHINECOUNT" + "|" + "READYMACHINECOUNT" + "|" + "UPDATEDMACHINECOUNT" + "|" + "DEGRADEDMACHINECOUNT" + "|" + "AGE" + Colors.White}

	File, _ = ioutil.ReadFile(MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigpools/" + List2Item + ".yaml")

	MyMCP := MCP{}
	yaml.Unmarshal(File, &MyMCP)

	name := MyMCP.Metadata.Name

	config := MyMCP.Status.Configuration.Name
	configS := fmt.Sprintf("%v", config)

	conditions := MyMCP.Status.Conditions
	updatedS := ""
	updatingS := ""
	degradedS := ""
	for i := range conditions {
		if conditions[i].Type == "Updated" {
			if conditions[i].Status == "True" {
				updatedS = Colors.Green + "True" + Colors.White
			} else {
				updatedS = Colors.Red + "False" + Colors.White
			}

		} else if conditions[i].Type == "Updating" {
			if conditions[i].Status == "True" {
				updatingS = Colors.Red + "True" + Colors.White
			} else {
				updatingS = Colors.Green + "False" + Colors.White
			}
		} else if conditions[i].Type == "Degraded" {
			if conditions[i].Status == "True" {
				degradedS = Colors.Red + "True" + Colors.White
			} else {
				degradedS = Colors.Green + "False" + Colors.White
			}
		}

	}
	machineCount := MyMCP.Status.MachineCount
	machineCountS := fmt.Sprintf("%v", machineCount)

	machineReady := MyMCP.Status.ReadyMachineCount
	machineReadyS := fmt.Sprintf("%v", machineReady)

	machineUpdated := MyMCP.Status.UpdatedMachineCount
	machineUpdatedS := fmt.Sprintf("%v", machineUpdated)

	machineDegraded := MyMCP.Status.DegradedMachineCount
	machineDegradedS := fmt.Sprintf("%v", machineDegraded)

	now := time.Now().UTC()
	CreationTime := MyMCP.Metadata.CreationTimestamp
	diff := now.Sub(CreationTime).Seconds()
	diffInt := int(diff)
	seconds := strconv.Itoa((diffInt % 60))
	minutes := strconv.Itoa((diffInt / 60) % 60)
	hours := strconv.Itoa((diffInt / 360) % 24)
	days := strconv.Itoa((diffInt / 86400))
	age := ""
	if days != "0" {
		age = days + "d" + hours + "h"
	} else if days == "0" && hours != "" {
		age = hours + "h" + minutes + "m"
	} else if hours == "0" {
		age = minutes + "m" + seconds + "s"
	}
	Output = append(Output, Colors.White+name+"|"+configS+"|"+updatedS+"|"+updatingS+"|"+degradedS+"|"+machineCountS+"|"+machineReadyS+"|"+machineUpdatedS+"|"+machineDegradedS+"|"+age+Colors.White)
	FormatedOutput := columnize.SimpleFormat(Output)
	TextView.SetText(FormatedOutput)
	TextView.ScrollToBeginning()
	TextViewData = FormatedOutput
}

func GetCSRInfo() {
	now := time.Now().UTC()
	Files, _ = ioutil.ReadDir(MG_Path + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/")
	Output := []string{Colors.Yellow + "NAME" + "|" + "AGE" + "|" + "SIGNERNAME" + "|" + "REQUESTOR" + "|" + "CONDITION" + Colors.White}
	if len(Files) == 0 {
		TextView.SetText("No available data about CSR")
	} else {
		for _, File := range Files {
			var MyCSR = CSR{}
			yfile, _ := ioutil.ReadFile(MG_Path + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/" + File.Name())
			yaml.Unmarshal(yfile, &MyCSR)

			name := MyCSR.Metadata.Name

			CreationTime := MyCSR.Metadata.CreationTimestamp
			diff := now.Sub(CreationTime).Seconds()
			diffInt := int(diff)
			seconds := strconv.Itoa((diffInt % 60))
			minutes := strconv.Itoa((diffInt / 60) % 60)
			hours := strconv.Itoa((diffInt / 360) % 24)
			days := strconv.Itoa((diffInt / 86400))
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
}

func GetInstallPlan() {

}

func GetAge(CreationTime time.Time) string {
	now := time.Now().UTC()
	diff := now.Sub(CreationTime).Seconds()
	diffInt := int(diff)
	seconds := strconv.Itoa((diffInt % 60))
	minutes := strconv.Itoa((diffInt / 60) % 60)
	hours := strconv.Itoa((diffInt / 360) % 24)
	days := strconv.Itoa((diffInt / 86400))
	age := ""
	if days != "0" {
		age = days + "d" + hours + "h"
	} else if days == "0" && hours != "" {
		age = hours + "h" + minutes + "m"
	} else if hours == "0" {
		age = minutes + "m" + seconds + "s"
	}
	return age
}
