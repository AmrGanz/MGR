package functions

import (
	"io/fs"
	"time"

	"github.com/rivo/tview"
)

var Files []fs.FileInfo
var File []byte
var Err error
var Output []string
var App = tview.NewApplication()
var MainGrid = tview.NewGrid()
var SearchModeGrid = tview.NewGrid()
var CopyModeGrid = tview.NewGrid()
var MGDropDown = tview.NewDropDown()
var ClusterInfoButton = tview.NewButton("Cluster Info")
var ClusterStatusButton = tview.NewButton("Cluster Status")
var HelpButton = tview.NewButton("Help")
var FocusModeButton = tview.NewButton("Focus Window")
var ExitButton = tview.NewButton("Exit")
var SearchBox = tview.NewInputField()
var CopyModeButton = tview.NewButton("Copy Mode")
var SearchButton = tview.NewButton("Search")
var GoBackButton = tview.NewButton("Go Back")
var ActivePathBox = tview.NewTextView()
var List1 = tview.NewList()
var List2 = tview.NewList()
var List3 = tview.NewList()
var List4 = tview.NewList()
var List5 = tview.NewList()
var List6 = tview.NewList()
var TextView = tview.NewTextView()
var MessageWindow = tview.NewModal()
var List1Item string = ""
var List2Item string = ""
var List3Item string = ""
var List4Item string = ""
var List5Item string = ""
var List6Item string = ""
var TextViewData string = ""
var SearchResult []string = []string{""}
var height = 1
var width = 15
var rows []int = []int{height, 20, 10, 0, height}
var columns []int = []int{width, width, width, width, 0, width, width, width}
var ProvidedDirPath = "empty"
var BasePath = "empty"
var MyNode = Node{}
var YamlFile []byte

// Colors HEX codes:
var Colors struct {
	White  string
	Yellow string
	Red    string
	Green  string
	Blue   string
	Orange string
	Filler string
}

type Node struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations       map[string]string `yaml:"annotations"`
		CreationTimestamp time.Time         `yaml:"creationTimestamp"`
		Labels            map[string]string `yaml:"labels"`
		Name              string            `yaml:"name"`
		ResourceVersion   string            `yaml:"resourceVersion"`
		SelfLink          string            `yaml:"selfLink"`
		UID               string            `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		ProviderID    string `yaml:"providerID"`
		Unschedulable bool   `yaml:"unschedulable"`
	} `yaml:"spec"`
	Status struct {
		Addresses []struct {
			Address string `yaml:"address"`
			Type    string `yaml:"type"`
		} `yaml:"addresses"`
		Allocatable struct {
			CPU              string `yaml:"cpu"`
			EphemeralStorage string `yaml:"ephemeral-storage"`
			Hugepages2Mi     string `yaml:"hugepages-2Mi"`
			Memory           string `yaml:"memory"`
			Pods             string `yaml:"pods"`
		} `yaml:"allocatable"`
		Capacity struct {
			CPU              string `yaml:"cpu"`
			EphemeralStorage string `yaml:"ephemeral-storage"`
			Hugepages2Mi     string `yaml:"hugepages-2Mi"`
			Memory           string `yaml:"memory"`
			Pods             string `yaml:"pods"`
		} `yaml:"capacity"`
		Conditions []struct {
			LastHeartbeatTime  time.Time `yaml:"lastHeartbeatTime"`
			LastTransitionTime time.Time `yaml:"lastTransitionTime"`
			Message            string    `yaml:"message"`
			Reason             string    `yaml:"reason"`
			Status             string    `yaml:"status"`
			Type               string    `yaml:"type"`
		} `yaml:"conditions"`
		DaemonEndpoints struct {
			KubeletEndpoint struct {
				Port int `yaml:"Port"`
			} `yaml:"kubeletEndpoint"`
		} `yaml:"daemonEndpoints"`
		Images []struct {
			Names     []string `yaml:"names"`
			SizeBytes int      `yaml:"sizeBytes"`
		} `yaml:"images"`
		NodeInfo struct {
			Architecture            string `yaml:"architecture"`
			BootID                  string `yaml:"bootID"`
			ContainerRuntimeVersion string `yaml:"containerRuntimeVersion"`
			KernelVersion           string `yaml:"kernelVersion"`
			KubeProxyVersion        string `yaml:"kubeProxyVersion"`
			KubeletVersion          string `yaml:"kubeletVersion"`
			MachineID               string `yaml:"machineID"`
			OperatingSystem         string `yaml:"operatingSystem"`
			OsImage                 string `yaml:"osImage"`
			SystemUUID              string `yaml:"systemUUID"`
		} `yaml:"nodeInfo"`
		VolumesAttached []struct {
			DevicePath string `yaml:"devicePath"`
			Name       string `yaml:"name"`
		} `yaml:"volumesAttached"`
		VolumesInUse []string `yaml:"volumesInUse"`
	} `yaml:"status"`
}

type MC struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations       map[string]string `yaml:"annotations"`
		CreationTimestamp time.Time         `yaml:"creationTimestamp"`
		Generation        int               `yaml:"generation"`
		Labels            map[string]string `yaml:"labels"`
		Name              string            `yaml:"name"`
		OwnerReferences   []struct {
			APIVersion         string `yaml:"apiVersion"`
			BlockOwnerDeletion bool   `yaml:"blockOwnerDeletion"`
			Controller         bool   `yaml:"controller"`
			Kind               string `yaml:"kind"`
			Name               string `yaml:"name"`
			UID                string `yaml:"uid"`
		} `yaml:"ownerReferences"`
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
		UID             string `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Config struct {
			Ignition struct {
				Version string `yaml:"version"`
			} `yaml:"ignition"`
			Storage struct {
				Files []struct {
					Contents struct {
						Source string `yaml:"source"`
					} `yaml:"contents"`
					Mode      int    `yaml:"mode"`
					Overwrite bool   `yaml:"overwrite"`
					Path      string `yaml:"path"`
				} `yaml:"files"`
			} `yaml:"storage"`
			Systemd struct {
				Units []struct {
					Contents string `yaml:"contents"`
					Enabled  bool   `yaml:"enabled"`
					Name     string `yaml:"name"`
				} `yaml:"units"`
			} `yaml:"systemd"`
		} `yaml:"config"`
		Extensions      interface{} `yaml:"extensions"`
		Fips            bool        `yaml:"fips"`
		KernelArguments interface{} `yaml:"kernelArguments"`
		KernelType      string      `yaml:"kernelType"`
		OsImageURL      string      `yaml:"osImageURL"`
	} `yaml:"spec"`
}

type CSR struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		GenerateName      string    `yaml:"generateName"`
		Name              string    `yaml:"name"`
		ResourceVersion   string    `yaml:"resourceVersion"`
		SelfLink          string    `yaml:"selfLink"`
		UID               string    `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Groups     []string `yaml:"groups"`
		Request    string   `yaml:"request"`
		SignerName string   `yaml:"signerName"`
		Usages     []string `yaml:"usages"`
		Username   string   `yaml:"username"`
	} `yaml:"spec"`
	Status struct {
		Certificate string `yaml:"certificate"`
		Conditions  []struct {
			LastTransitionTime time.Time `yaml:"lastTransitionTime"`
			LastUpdateTime     time.Time `yaml:"lastUpdateTime"`
			Message            string    `yaml:"message"`
			Reason             string    `yaml:"reason"`
			Status             string    `yaml:"status"`
			Type               string    `yaml:"type"`
		} `yaml:"conditions"`
	} `yaml:"status"`
}

type OPERATOR struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Annotations       map[string]string `yaml:"annotations"`
			CreationTimestamp time.Time         `yaml:"creationTimestamp"`
			Generation        int               `yaml:"generation"`
			Name              string            `yaml:"name"`
			ResourceVersion   string            `yaml:"resourceVersion"`
			SelfLink          string            `yaml:"selfLink"`
			UID               string            `yaml:"uid"`
		} `yaml:"metadata"`
		Spec struct {
		} `yaml:"spec"`
		Status struct {
			Conditions []struct {
				LastTransitionTime time.Time `yaml:"lastTransitionTime"`
				Message            string    `yaml:"message"`
				Reason             string    `yaml:"reason"`
				Status             string    `yaml:"status"`
				Type               string    `yaml:"type"`
			} `yaml:"conditions"`
			Extension      interface{} `yaml:"extension"`
			RelatedObjects []struct {
				Group     string `yaml:"group"`
				Name      string `yaml:"name"`
				Resource  string `yaml:"resource"`
				Namespace string `yaml:"namespace,omitempty"`
			} `yaml:"relatedObjects"`
			Versions []struct {
				Name    string `yaml:"name"`
				Version string `yaml:"version"`
			} `yaml:"versions"`
		} `yaml:"status"`
	} `yaml:"items"`
}
