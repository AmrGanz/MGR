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

func GetNodesInfo(NodeName string, Flag string) {
	// Cleaning TextView and TextViewData
	TextView.Clear()
	TextViewData = ""
	if NodeName == "All Nodes" {
		Files, Err = ioutil.ReadDir(BasePath + "cluster-scoped-resources/core/nodes/")
		if Err != nil {
			TextView.SetText(Err.Error())
		} else {
			if Flag == "Summary" {
				Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + Colors.White + "\n"}
				for i := range Files {
					ReadNodeYaml(Files[i].Name(), "Summary")
				}
			} else if Flag == "Details" {
				Output = []string{Colors.Yellow + "NAME" + "|" + Colors.Yellow + "STATUS" + Colors.Yellow + "|" + "ROLES" + "|" + "AGE" + "|" + "VERSION" + "|" + "INTERNAL-IP" + "|" + "EXTERNAL-IP" + "|" + "OS-IMAGE" + "|" + "KERNEL-VERSION" + "|" + "CONTAINER-RUNTIME" + Colors.White + "\n"}
				for i := range Files {
					ReadNodeYaml(Files[i].Name(), "Details")
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
	now := time.Now().UTC()
	yfile, _ := ioutil.ReadFile(BasePath + "cluster-scoped-resources/core/nodes/" + NodeFileName)
	yaml.Unmarshal(yfile, &MyNode)

	name := MyNode.Metadata.Name

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
	for key, _ := range Labels {
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
