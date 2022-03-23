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
var TableGrid = tview.NewGrid()
var Table = tview.NewTable()

// var SearchModeGrid = tview.NewGrid()
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
var ProvidedDirPath = ""
var MG_Path string

var MyNode_Public = NODE{}
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
var Version_Path = ""
var Configurations_Path = ""
var Namespaces_Path = ""
var Nodes_Path = ""
var Operators_Path = ""
var InstalledOperators_Path = ""
var InstallPlans_Path = ""
var MCP_Path = ""
var MC_Path = ""
var PV_Path = ""
var CSR_Path = ""

// Resources Paths
func SetResourcesPath() {
	Version_Path = MG_Path + "cluster-scoped-resources/config.openshift.io/clusterversions/version.yaml"
	Configurations_Path = MG_Path + "cluster-scoped-resources/config.openshift.io/"
	Namespaces_Path = MG_Path + "namespaces/"
	Nodes_Path = MG_Path + "cluster-scoped-resources/core/nodes/"
	Operators_Path = MG_Path + "cluster-scoped-resources/config.openshift.io/clusteroperators.yaml"
	InstalledOperators_Path = MG_Path + "/cluster-scoped-resources/operators.coreos.com/operators/"
	InstallPlans_Path = "/operators.coreos.com/installplans/"
	MCP_Path = MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigpools/"
	MC_Path = MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigs/"
	PV_Path = MG_Path + "cluster-scoped-resources/core/persistentvolumes/"
	CSR_Path = MG_Path + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/"
}

type NODE struct {
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
		PodCIDR       string   `yaml:"podCIDR"`
		PodCIDRs      []string `yaml:"podCIDRs"`
		ProviderID    string   `yaml:"providerID"`
		Unschedulable bool     `yaml:"unschedulable"`
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

// type NODE struct {
// 	APIVersion string `yaml:"apiVersion"`
// 	Kind       string `yaml:"kind"`
// 	Metadata   struct {
// 		Annotations       map[string]string `yaml:"annotations"`
// 		CreationTimestamp time.Time         `yaml:"creationTimestamp"`
// 		Labels            map[string]string `yaml:"labels"`
// 		Name              string            `yaml:"name"`
// 		ResourceVersion   string            `yaml:"resourceVersion"`
// 		SelfLink          string            `yaml:"selfLink"`
// 		UID               string            `yaml:"uid"`
// 	} `yaml:"metadata"`
// 	Spec struct {
// 		ProviderID    string `yaml:"providerID"`
// 		Unschedulable bool   `yaml:"unschedulable"`
// 	} `yaml:"spec"`
// 	Status struct {
// 		Addresses []struct {
// 			Address string `yaml:"address"`
// 			Type    string `yaml:"type"`
// 		} `yaml:"addresses"`
// 		Allocatable struct {
// 			CPU              string `yaml:"cpu"`
// 			EphemeralStorage string `yaml:"ephemeral-storage"`
// 			Hugepages2Mi     string `yaml:"hugepages-2Mi"`
// 			Memory           string `yaml:"memory"`
// 			Pods             string `yaml:"pods"`
// 		} `yaml:"allocatable"`
// 		Capacity struct {
// 			CPU              string `yaml:"cpu"`
// 			EphemeralStorage string `yaml:"ephemeral-storage"`
// 			Hugepages2Mi     string `yaml:"hugepages-2Mi"`
// 			Memory           string `yaml:"memory"`
// 			Pods             string `yaml:"pods"`
// 		} `yaml:"capacity"`
// 		Conditions []struct {
// 			LastHeartbeatTime  time.Time `yaml:"lastHeartbeatTime"`
// 			LastTransitionTime time.Time `yaml:"lastTransitionTime"`
// 			Message            string    `yaml:"message"`
// 			Reason             string    `yaml:"reason"`
// 			Status             string    `yaml:"status"`
// 			Type               string    `yaml:"type"`
// 		} `yaml:"conditions"`
// 		DaemonEndpoints struct {
// 			KubeletEndpoint struct {
// 				Port int `yaml:"Port"`
// 			} `yaml:"kubeletEndpoint"`
// 		} `yaml:"daemonEndpoints"`
// 		Images []struct {
// 			Names     []string `yaml:"names"`
// 			SizeBytes int      `yaml:"sizeBytes"`
// 		} `yaml:"images"`
// 		NodeInfo struct {
// 			Architecture            string `yaml:"architecture"`
// 			BootID                  string `yaml:"bootID"`
// 			ContainerRuntimeVersion string `yaml:"containerRuntimeVersion"`
// 			KernelVersion           string `yaml:"kernelVersion"`
// 			KubeProxyVersion        string `yaml:"kubeProxyVersion"`
// 			KubeletVersion          string `yaml:"kubeletVersion"`
// 			MachineID               string `yaml:"machineID"`
// 			OperatingSystem         string `yaml:"operatingSystem"`
// 			OsImage                 string `yaml:"osImage"`
// 			SystemUUID              string `yaml:"systemUUID"`
// 		} `yaml:"nodeInfo"`
// 		VolumesAttached []struct {
// 			DevicePath string `yaml:"devicePath"`
// 			Name       string `yaml:"name"`
// 		} `yaml:"volumesAttached"`
// 		VolumesInUse []string `yaml:"volumesInUse"`
// 	} `yaml:"status"`
// }

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

type MCP struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Generation        int       `yaml:"generation"`
		Labels            struct {
			MachineconfigurationOpenshiftIoMcoBuiltIn                 string `yaml:"machineconfiguration.openshift.io/mco-built-in"`
			OperatorMachineconfigurationOpenshiftIoRequiredForUpgrade string `yaml:"operator.machineconfiguration.openshift.io/required-for-upgrade"`
			PoolsOperatorMachineconfigurationOpenshiftIoMaster        string `yaml:"pools.operator.machineconfiguration.openshift.io/master"`
		} `yaml:"labels"`
		Name            string `yaml:"name"`
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
		UID             string `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Configuration struct {
			Name   string `yaml:"name"`
			Source []struct {
				APIVersion string `yaml:"apiVersion"`
				Kind       string `yaml:"kind"`
				Name       string `yaml:"name"`
			} `yaml:"source"`
		} `yaml:"configuration"`
		MachineConfigSelector struct {
			MatchLabels struct {
				MachineconfigurationOpenshiftIoRole string `yaml:"machineconfiguration.openshift.io/role"`
			} `yaml:"matchLabels"`
		} `yaml:"machineConfigSelector"`
		NodeSelector struct {
			MatchLabels struct {
				NodeRoleKubernetesIoMaster string `yaml:"node-role.kubernetes.io/master"`
			} `yaml:"matchLabels"`
		} `yaml:"nodeSelector"`
		Paused bool `yaml:"paused"`
	} `yaml:"spec"`
	Status struct {
		Conditions []struct {
			LastTransitionTime time.Time `yaml:"lastTransitionTime"`
			Message            string    `yaml:"message"`
			Reason             string    `yaml:"reason"`
			Status             string    `yaml:"status"`
			Type               string    `yaml:"type"`
		} `yaml:"conditions"`
		Configuration struct {
			Name   string `yaml:"name"`
			Source []struct {
				APIVersion string `yaml:"apiVersion"`
				Kind       string `yaml:"kind"`
				Name       string `yaml:"name"`
			} `yaml:"source"`
		} `yaml:"configuration"`
		DegradedMachineCount    int `yaml:"degradedMachineCount"`
		MachineCount            int `yaml:"machineCount"`
		ObservedGeneration      int `yaml:"observedGeneration"`
		ReadyMachineCount       int `yaml:"readyMachineCount"`
		UnavailableMachineCount int `yaml:"unavailableMachineCount"`
		UpdatedMachineCount     int `yaml:"updatedMachineCount"`
	} `yaml:"status"`
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
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Generation        int       `yaml:"generation"`
		Name              string    `yaml:"name"`
		ResourceVersion   string    `yaml:"resourceVersion"`
		SelfLink          string    `yaml:"selfLink"`
		UID               string    `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
	} `yaml:"spec"`
	Status struct {
		Components struct {
			LabelSelector struct {
				MatchExpressions []struct {
					Key      string `yaml:"key"`
					Operator string `yaml:"operator"`
				} `yaml:"matchExpressions"`
			} `yaml:"labelSelector"`
			Refs []struct {
				APIVersion string `yaml:"apiVersion"`
				Conditions []struct {
					LastTransitionTime time.Time `yaml:"lastTransitionTime"`
					LastUpdateTime     time.Time `yaml:"lastUpdateTime"`
					Message            string    `yaml:"message"`
					Reason             string    `yaml:"reason"`
					Status             string    `yaml:"status"`
					Type               string    `yaml:"type"`
				} `yaml:"conditions,omitempty"`
				Kind      string `yaml:"kind"`
				Name      string `yaml:"name"`
				Namespace string `yaml:"namespace,omitempty"`
			} `yaml:"refs"`
		} `yaml:"components"`
	} `yaml:"status"`
}

type OPERATORS struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Annotations struct {
				ExcludeReleaseOpenshiftIoInternalOpenshiftHosted     string `yaml:"exclude.release.openshift.io/internal-openshift-hosted"`
				IncludeReleaseOpenshiftIoSelfManagedHighAvailability string `yaml:"include.release.openshift.io/self-managed-high-availability"`
			} `yaml:"annotations"`
			CreationTimestamp time.Time `yaml:"creationTimestamp"`
			Generation        int       `yaml:"generation"`
			Name              string    `yaml:"name"`
			ResourceVersion   string    `yaml:"resourceVersion"`
			SelfLink          string    `yaml:"selfLink"`
			UID               string    `yaml:"uid"`
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

type POD struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations struct {
			KubernetesIoConfigHash   string    `yaml:"kubernetes.io/config.hash"`
			KubernetesIoConfigMirror string    `yaml:"kubernetes.io/config.mirror"`
			KubernetesIoConfigSeen   time.Time `yaml:"kubernetes.io/config.seen"`
			KubernetesIoConfigSource string    `yaml:"kubernetes.io/config.source"`
		} `yaml:"annotations"`
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Labels            struct {
			App      string `yaml:"app"`
			Etcd     string `yaml:"etcd"`
			K8SApp   string `yaml:"k8s-app"`
			Revision string `yaml:"revision"`
		} `yaml:"labels"`
		Name            string `yaml:"name"`
		Namespace       string `yaml:"namespace"`
		OwnerReferences []struct {
			APIVersion string `yaml:"apiVersion"`
			Controller bool   `yaml:"controller"`
			Kind       string `yaml:"kind"`
			Name       string `yaml:"name"`
			UID        string `yaml:"uid"`
		} `yaml:"ownerReferences"`
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
		UID             string `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Containers []struct {
			Command []string `yaml:"command"`
			Env     []struct {
				Name  string `yaml:"name"`
				Value string `yaml:"value"`
			} `yaml:"env"`
			Image           string `yaml:"image"`
			ImagePullPolicy string `yaml:"imagePullPolicy"`
			Name            string `yaml:"name"`
			Resources       struct {
				Requests struct {
					CPU    string `yaml:"cpu"`
					Memory string `yaml:"memory"`
				} `yaml:"requests"`
			} `yaml:"resources"`
			TerminationMessagePath   string `yaml:"terminationMessagePath"`
			TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
			VolumeMounts             []struct {
				MountPath string `yaml:"mountPath"`
				Name      string `yaml:"name"`
			} `yaml:"volumeMounts"`
			ReadinessProbe struct {
				FailureThreshold    int `yaml:"failureThreshold"`
				InitialDelaySeconds int `yaml:"initialDelaySeconds"`
				PeriodSeconds       int `yaml:"periodSeconds"`
				SuccessThreshold    int `yaml:"successThreshold"`
				TCPSocket           struct {
					Port int `yaml:"port"`
				} `yaml:"tcpSocket"`
				TimeoutSeconds int `yaml:"timeoutSeconds"`
			} `yaml:"readinessProbe,omitempty"`
			SecurityContext struct {
				Privileged bool `yaml:"privileged"`
			} `yaml:"securityContext,omitempty"`
		} `yaml:"containers"`
		DNSPolicy          string `yaml:"dnsPolicy"`
		EnableServiceLinks bool   `yaml:"enableServiceLinks"`
		HostNetwork        bool   `yaml:"hostNetwork"`
		InitContainers     []struct {
			Command []string `yaml:"command"`
			Env     []struct {
				Name      string `yaml:"name"`
				Value     string `yaml:"value,omitempty"`
				ValueFrom struct {
					FieldRef struct {
						APIVersion string `yaml:"apiVersion"`
						FieldPath  string `yaml:"fieldPath"`
					} `yaml:"fieldRef"`
				} `yaml:"valueFrom,omitempty"`
			} `yaml:"env,omitempty"`
			Image           string `yaml:"image"`
			ImagePullPolicy string `yaml:"imagePullPolicy"`
			Name            string `yaml:"name"`
			Resources       struct {
				Requests struct {
					CPU    string `yaml:"cpu"`
					Memory string `yaml:"memory"`
				} `yaml:"requests"`
			} `yaml:"resources"`
			SecurityContext struct {
				Privileged bool `yaml:"privileged"`
			} `yaml:"securityContext"`
			TerminationMessagePath   string `yaml:"terminationMessagePath"`
			TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
			VolumeMounts             []struct {
				MountPath string `yaml:"mountPath"`
				Name      string `yaml:"name"`
			} `yaml:"volumeMounts,omitempty"`
		} `yaml:"initContainers"`
		NodeName          string `yaml:"nodeName"`
		PreemptionPolicy  string `yaml:"preemptionPolicy"`
		Priority          int    `yaml:"priority"`
		PriorityClassName string `yaml:"priorityClassName"`
		RestartPolicy     string `yaml:"restartPolicy"`
		SchedulerName     string `yaml:"schedulerName"`
		SecurityContext   struct {
		} `yaml:"securityContext"`
		TerminationGracePeriodSeconds int `yaml:"terminationGracePeriodSeconds"`
		Tolerations                   []struct {
			Operator string `yaml:"operator"`
		} `yaml:"tolerations"`
		Volumes []struct {
			HostPath struct {
				Path string `yaml:"path"`
				Type string `yaml:"type"`
			} `yaml:"hostPath"`
			Name string `yaml:"name"`
		} `yaml:"volumes"`
	} `yaml:"spec"`
	Status struct {
		Conditions []struct {
			LastProbeTime      interface{} `yaml:"lastProbeTime"`
			LastTransitionTime time.Time   `yaml:"lastTransitionTime"`
			Status             string      `yaml:"status"`
			Type               string      `yaml:"type"`
		} `yaml:"conditions"`
		ContainerStatuses []struct {
			ContainerID string `yaml:"containerID"`
			Image       string `yaml:"image"`
			ImageID     string `yaml:"imageID"`
			LastState   struct {
			} `yaml:"lastState"`
			Name         string `yaml:"name"`
			Ready        bool   `yaml:"ready"`
			RestartCount int    `yaml:"restartCount"`
			Started      bool   `yaml:"started"`
			State        struct {
				Running struct {
					StartedAt time.Time `yaml:"startedAt"`
				} `yaml:"running"`
			} `yaml:"state"`
		} `yaml:"containerStatuses"`
		HostIP                string `yaml:"hostIP"`
		InitContainerStatuses []struct {
			ContainerID string `yaml:"containerID"`
			Image       string `yaml:"image"`
			ImageID     string `yaml:"imageID"`
			LastState   struct {
			} `yaml:"lastState"`
			Name         string `yaml:"name"`
			Ready        bool   `yaml:"ready"`
			RestartCount int    `yaml:"restartCount"`
			State        struct {
				Terminated struct {
					ContainerID string    `yaml:"containerID"`
					ExitCode    int       `yaml:"exitCode"`
					FinishedAt  time.Time `yaml:"finishedAt"`
					Reason      string    `yaml:"reason"`
					StartedAt   time.Time `yaml:"startedAt"`
				} `yaml:"terminated"`
			} `yaml:"state"`
		} `yaml:"initContainerStatuses"`
		Phase  string `yaml:"phase"`
		PodIP  string `yaml:"podIP"`
		PodIPs []struct {
			IP string `yaml:"ip"`
		} `yaml:"podIPs"`
		QosClass  string    `yaml:"qosClass"`
		StartTime time.Time `yaml:"startTime"`
	} `yaml:"status"`
}

type PODS struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Annotations struct {
				KubernetesIoConfigHash   string    `yaml:"kubernetes.io/config.hash"`
				KubernetesIoConfigMirror string    `yaml:"kubernetes.io/config.mirror"`
				KubernetesIoConfigSeen   time.Time `yaml:"kubernetes.io/config.seen"`
				KubernetesIoConfigSource string    `yaml:"kubernetes.io/config.source"`
			} `yaml:"annotations"`
			CreationTimestamp time.Time `yaml:"creationTimestamp"`
			Labels            struct {
				App      string `yaml:"app"`
				Etcd     string `yaml:"etcd"`
				K8SApp   string `yaml:"k8s-app"`
				Revision string `yaml:"revision"`
			} `yaml:"labels"`
			Name            string `yaml:"name"`
			Namespace       string `yaml:"namespace"`
			OwnerReferences []struct {
				APIVersion string `yaml:"apiVersion"`
				Controller bool   `yaml:"controller"`
				Kind       string `yaml:"kind"`
				Name       string `yaml:"name"`
				UID        string `yaml:"uid"`
			} `yaml:"ownerReferences"`
			ResourceVersion string `yaml:"resourceVersion"`
			SelfLink        string `yaml:"selfLink"`
			UID             string `yaml:"uid"`
		} `yaml:"metadata"`
		Spec struct {
			Containers []struct {
				Command []string `yaml:"command"`
				Env     []struct {
					Name  string `yaml:"name"`
					Value string `yaml:"value"`
				} `yaml:"env"`
				Image           string `yaml:"image"`
				ImagePullPolicy string `yaml:"imagePullPolicy"`
				Name            string `yaml:"name"`
				Resources       struct {
					Requests struct {
						CPU    string `yaml:"cpu"`
						Memory string `yaml:"memory"`
					} `yaml:"requests"`
				} `yaml:"resources"`
				TerminationMessagePath   string `yaml:"terminationMessagePath"`
				TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
				VolumeMounts             []struct {
					MountPath string `yaml:"mountPath"`
					Name      string `yaml:"name"`
				} `yaml:"volumeMounts"`
				ReadinessProbe struct {
					FailureThreshold    int `yaml:"failureThreshold"`
					InitialDelaySeconds int `yaml:"initialDelaySeconds"`
					PeriodSeconds       int `yaml:"periodSeconds"`
					SuccessThreshold    int `yaml:"successThreshold"`
					TCPSocket           struct {
						Port int `yaml:"port"`
					} `yaml:"tcpSocket"`
					TimeoutSeconds int `yaml:"timeoutSeconds"`
				} `yaml:"readinessProbe,omitempty"`
				SecurityContext struct {
					Privileged bool `yaml:"privileged"`
				} `yaml:"securityContext,omitempty"`
			} `yaml:"containers"`
			DNSPolicy          string `yaml:"dnsPolicy"`
			EnableServiceLinks bool   `yaml:"enableServiceLinks"`
			HostNetwork        bool   `yaml:"hostNetwork"`
			InitContainers     []struct {
				Command []string `yaml:"command"`
				Env     []struct {
					Name      string `yaml:"name"`
					Value     string `yaml:"value,omitempty"`
					ValueFrom struct {
						FieldRef struct {
							APIVersion string `yaml:"apiVersion"`
							FieldPath  string `yaml:"fieldPath"`
						} `yaml:"fieldRef"`
					} `yaml:"valueFrom,omitempty"`
				} `yaml:"env,omitempty"`
				Image           string `yaml:"image"`
				ImagePullPolicy string `yaml:"imagePullPolicy"`
				Name            string `yaml:"name"`
				Resources       struct {
					Requests struct {
						CPU    string `yaml:"cpu"`
						Memory string `yaml:"memory"`
					} `yaml:"requests"`
				} `yaml:"resources"`
				SecurityContext struct {
					Privileged bool `yaml:"privileged"`
				} `yaml:"securityContext"`
				TerminationMessagePath   string `yaml:"terminationMessagePath"`
				TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
				VolumeMounts             []struct {
					MountPath string `yaml:"mountPath"`
					Name      string `yaml:"name"`
				} `yaml:"volumeMounts,omitempty"`
			} `yaml:"initContainers"`
			NodeName          string `yaml:"nodeName"`
			PreemptionPolicy  string `yaml:"preemptionPolicy"`
			Priority          int    `yaml:"priority"`
			PriorityClassName string `yaml:"priorityClassName"`
			RestartPolicy     string `yaml:"restartPolicy"`
			SchedulerName     string `yaml:"schedulerName"`
			SecurityContext   struct {
			} `yaml:"securityContext"`
			TerminationGracePeriodSeconds int `yaml:"terminationGracePeriodSeconds"`
			Tolerations                   []struct {
				Operator string `yaml:"operator"`
			} `yaml:"tolerations"`
			Volumes []struct {
				HostPath struct {
					Path string `yaml:"path"`
					Type string `yaml:"type"`
				} `yaml:"hostPath"`
				Name string `yaml:"name"`
			} `yaml:"volumes"`
		} `yaml:"spec"`
		Status struct {
			Conditions []struct {
				LastProbeTime      interface{} `yaml:"lastProbeTime"`
				LastTransitionTime time.Time   `yaml:"lastTransitionTime"`
				Status             string      `yaml:"status"`
				Type               string      `yaml:"type"`
			} `yaml:"conditions"`
			ContainerStatuses []struct {
				ContainerID string `yaml:"containerID"`
				Image       string `yaml:"image"`
				ImageID     string `yaml:"imageID"`
				LastState   struct {
				} `yaml:"lastState"`
				Name         string `yaml:"name"`
				Ready        bool   `yaml:"ready"`
				RestartCount int    `yaml:"restartCount"`
				Started      bool   `yaml:"started"`
				State        struct {
					Running struct {
						StartedAt time.Time `yaml:"startedAt"`
					} `yaml:"running"`
				} `yaml:"state"`
			} `yaml:"containerStatuses"`
			HostIP                string `yaml:"hostIP"`
			InitContainerStatuses []struct {
				ContainerID string `yaml:"containerID"`
				Image       string `yaml:"image"`
				ImageID     string `yaml:"imageID"`
				LastState   struct {
				} `yaml:"lastState"`
				Name         string `yaml:"name"`
				Ready        bool   `yaml:"ready"`
				RestartCount int    `yaml:"restartCount"`
				State        struct {
					Terminated struct {
						ContainerID string    `yaml:"containerID"`
						ExitCode    int       `yaml:"exitCode"`
						FinishedAt  time.Time `yaml:"finishedAt"`
						Reason      string    `yaml:"reason"`
						StartedAt   time.Time `yaml:"startedAt"`
					} `yaml:"terminated"`
				} `yaml:"state"`
			} `yaml:"initContainerStatuses"`
			Phase  string `yaml:"phase"`
			PodIP  string `yaml:"podIP"`
			PodIPs []struct {
				IP string `yaml:"ip"`
			} `yaml:"podIPs"`
			QosClass  string    `yaml:"qosClass"`
			StartTime time.Time `yaml:"startTime"`
		} `yaml:"status"`
	} `yaml:"items"`
}

type CLUSTERVERSION struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Generation        int       `yaml:"generation"`
		Name              string    `yaml:"name"`
		ResourceVersion   string    `yaml:"resourceVersion"`
		SelfLink          string    `yaml:"selfLink"`
		UID               string    `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Channel       string `yaml:"channel"`
		ClusterID     string `yaml:"clusterID"`
		DesiredUpdate struct {
			Force   bool   `yaml:"force"`
			Image   string `yaml:"image"`
			Version string `yaml:"version"`
		} `yaml:"desiredUpdate"`
		Upstream string `yaml:"upstream"`
	} `yaml:"spec"`
	Status struct {
		AvailableUpdates []struct {
			Channels []string `yaml:"channels"`
			Image    string   `yaml:"image"`
			URL      string   `yaml:"url"`
			Version  string   `yaml:"version"`
		} `yaml:"availableUpdates"`
		Conditions []struct {
			LastTransitionTime time.Time `yaml:"lastTransitionTime"`
			Message            string    `yaml:"message,omitempty"`
			Status             string    `yaml:"status"`
			Type               string    `yaml:"type"`
		} `yaml:"conditions"`
		Desired struct {
			Channels []string `yaml:"channels"`
			Image    string   `yaml:"image"`
			URL      string   `yaml:"url"`
			Version  string   `yaml:"version"`
		} `yaml:"desired"`
		History []struct {
			CompletionTime time.Time `yaml:"completionTime"`
			Image          string    `yaml:"image"`
			StartedTime    time.Time `yaml:"startedTime"`
			State          string    `yaml:"state"`
			Verified       bool      `yaml:"verified"`
			Version        string    `yaml:"version"`
		} `yaml:"history"`
		ObservedGeneration int    `yaml:"observedGeneration"`
		VersionHash        string `yaml:"versionHash"`
	} `yaml:"status"`
}

type INSTALLPLAN struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		GenerateName      string    `yaml:"generateName"`
		Generation        int       `yaml:"generation"`
		Name              string    `yaml:"name"`
		Namespace         string    `yaml:"namespace"`
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
		Approval                   string   `yaml:"approval"`
		Approved                   bool     `yaml:"approved"`
		ClusterServiceVersionNames []string `yaml:"clusterServiceVersionNames"`
		Generation                 int      `yaml:"generation"`
	} `yaml:"spec"`
	Status struct {
		BundleLookups []struct {
			CatalogSourceRef struct {
				Name      string `yaml:"name"`
				Namespace string `yaml:"namespace"`
			} `yaml:"catalogSourceRef"`
			Identifier string `yaml:"identifier"`
			Path       string `yaml:"path"`
			Properties string `yaml:"properties"`
			Replaces   string `yaml:"replaces"`
		} `yaml:"bundleLookups"`
		CatalogSources []interface{} `yaml:"catalogSources"`
		Conditions     []struct {
			LastTransitionTime time.Time `yaml:"lastTransitionTime"`
			LastUpdateTime     time.Time `yaml:"lastUpdateTime"`
			Status             string    `yaml:"status"`
			Type               string    `yaml:"type"`
		} `yaml:"conditions"`
		Phase string `yaml:"phase"`
		Plan  []struct {
			Resolving string `yaml:"resolving"`
			Resource  struct {
				Group           string `yaml:"group"`
				Kind            string `yaml:"kind"`
				Manifest        string `yaml:"manifest"`
				Name            string `yaml:"name"`
				SourceName      string `yaml:"sourceName"`
				SourceNamespace string `yaml:"sourceNamespace"`
				Version         string `yaml:"version"`
			} `yaml:"resource"`
			Status string `yaml:"status"`
		} `yaml:"plan"`
	} `yaml:"status"`
}

type PV struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations struct {
			KubernetesIoCreatedby           string `yaml:"kubernetes.io/createdby"`
			PvKubernetesIoBoundByController string `yaml:"pv.kubernetes.io/bound-by-controller"`
			PvKubernetesIoProvisionedBy     string `yaml:"pv.kubernetes.io/provisioned-by"`
		} `yaml:"annotations"`
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Finalizers        []string  `yaml:"finalizers"`
		Name              string    `yaml:"name"`
		ResourceVersion   string    `yaml:"resourceVersion"`
		SelfLink          string    `yaml:"selfLink"`
		UID               string    `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		AccessModes []string `yaml:"accessModes"`
		Capacity    struct {
			Storage string `yaml:"storage"`
		} `yaml:"capacity"`
		ClaimRef struct {
			APIVersion      string `yaml:"apiVersion"`
			Kind            string `yaml:"kind"`
			Name            string `yaml:"name"`
			Namespace       string `yaml:"namespace"`
			ResourceVersion string `yaml:"resourceVersion"`
			UID             string `yaml:"uid"`
		} `yaml:"claimRef"`
		PersistentVolumeReclaimPolicy string `yaml:"persistentVolumeReclaimPolicy"`
		StorageClassName              string `yaml:"storageClassName"`
		VolumeMode                    string `yaml:"volumeMode"`
		VsphereVolume                 struct {
			FsType     string `yaml:"fsType"`
			VolumePath string `yaml:"volumePath"`
		} `yaml:"vsphereVolume"`
	} `yaml:"spec"`
	Status struct {
		Phase string `yaml:"phase"`
	} `yaml:"status"`
}

type DEPLOYMENT struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Annotations struct {
				ConfigOpenshiftIoInjectProxy                         string `yaml:"config.openshift.io/inject-proxy"`
				DeploymentKubernetesIoRevision                       string `yaml:"deployment.kubernetes.io/revision"`
				IncludeReleaseOpenshiftIoIbmCloudManaged             string `yaml:"include.release.openshift.io/ibm-cloud-managed"`
				IncludeReleaseOpenshiftIoSelfManagedHighAvailability string `yaml:"include.release.openshift.io/self-managed-high-availability"`
				IncludeReleaseOpenshiftIoSingleNodeDeveloper         string `yaml:"include.release.openshift.io/single-node-developer"`
			} `yaml:"annotations"`
			CreationTimestamp time.Time `yaml:"creationTimestamp"`
			Generation        int       `yaml:"generation"`
			Name              string    `yaml:"name"`
			Namespace         string    `yaml:"namespace"`
			ResourceVersion   string    `yaml:"resourceVersion"`
			SelfLink          string    `yaml:"selfLink"`
			UID               string    `yaml:"uid"`
		} `yaml:"metadata"`
		Spec struct {
			ProgressDeadlineSeconds int `yaml:"progressDeadlineSeconds"`
			Replicas                int `yaml:"replicas"`
			RevisionHistoryLimit    int `yaml:"revisionHistoryLimit"`
			Selector                struct {
				MatchLabels struct {
					Name string `yaml:"name"`
				} `yaml:"matchLabels"`
			} `yaml:"selector"`
			Strategy struct {
				RollingUpdate struct {
					MaxSurge       string `yaml:"maxSurge"`
					MaxUnavailable string `yaml:"maxUnavailable"`
				} `yaml:"rollingUpdate"`
				Type string `yaml:"type"`
			} `yaml:"strategy"`
			Template struct {
				Metadata struct {
					CreationTimestamp interface{} `yaml:"creationTimestamp"`
					Labels            struct {
						Name string `yaml:"name"`
					} `yaml:"labels"`
				} `yaml:"metadata"`
				Spec struct {
					Containers []struct {
						Args []string `yaml:"args"`
						Env  []struct {
							Name      string `yaml:"name"`
							Value     string `yaml:"value,omitempty"`
							ValueFrom struct {
								FieldRef struct {
									APIVersion string `yaml:"apiVersion"`
									FieldPath  string `yaml:"fieldPath"`
								} `yaml:"fieldRef"`
							} `yaml:"valueFrom,omitempty"`
						} `yaml:"env"`
						Image           string `yaml:"image"`
						ImagePullPolicy string `yaml:"imagePullPolicy"`
						Name            string `yaml:"name"`
						Ports           []struct {
							ContainerPort int    `yaml:"containerPort"`
							Name          string `yaml:"name"`
							Protocol      string `yaml:"protocol"`
						} `yaml:"ports"`
						Resources struct {
							Requests struct {
								CPU    string `yaml:"cpu"`
								Memory string `yaml:"memory"`
							} `yaml:"requests"`
						} `yaml:"resources"`
						TerminationMessagePath   string `yaml:"terminationMessagePath"`
						TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
						VolumeMounts             []struct {
							MountPath string `yaml:"mountPath"`
							Name      string `yaml:"name"`
							ReadOnly  bool   `yaml:"readOnly,omitempty"`
						} `yaml:"volumeMounts"`
					} `yaml:"containers"`
					DNSPolicy    string `yaml:"dnsPolicy"`
					NodeSelector struct {
						NodeRoleKubernetesIoMaster string `yaml:"node-role.kubernetes.io/master"`
					} `yaml:"nodeSelector"`
					PriorityClassName string `yaml:"priorityClassName"`
					RestartPolicy     string `yaml:"restartPolicy"`
					SchedulerName     string `yaml:"schedulerName"`
					SecurityContext   struct {
					} `yaml:"securityContext"`
					ServiceAccount                string `yaml:"serviceAccount"`
					ServiceAccountName            string `yaml:"serviceAccountName"`
					ShareProcessNamespace         bool   `yaml:"shareProcessNamespace"`
					TerminationGracePeriodSeconds int    `yaml:"terminationGracePeriodSeconds"`
					Tolerations                   []struct {
						Effect            string `yaml:"effect"`
						Key               string `yaml:"key"`
						Operator          string `yaml:"operator"`
						TolerationSeconds int    `yaml:"tolerationSeconds,omitempty"`
					} `yaml:"tolerations"`
					Volumes []struct {
						ConfigMap struct {
							DefaultMode int `yaml:"defaultMode"`
							Items       []struct {
								Key  string `yaml:"key"`
								Path string `yaml:"path"`
							} `yaml:"items"`
							Name     string `yaml:"name"`
							Optional bool   `yaml:"optional"`
						} `yaml:"configMap,omitempty"`
						Name   string `yaml:"name"`
						Secret struct {
							DefaultMode int    `yaml:"defaultMode"`
							SecretName  string `yaml:"secretName"`
						} `yaml:"secret,omitempty"`
						Projected struct {
							DefaultMode int `yaml:"defaultMode"`
							Sources     []struct {
								ServiceAccountToken struct {
									Audience          string `yaml:"audience"`
									ExpirationSeconds int    `yaml:"expirationSeconds"`
									Path              string `yaml:"path"`
								} `yaml:"serviceAccountToken"`
							} `yaml:"sources"`
						} `yaml:"projected,omitempty"`
					} `yaml:"volumes"`
				} `yaml:"spec"`
			} `yaml:"template"`
		} `yaml:"spec"`
		Status struct {
			AvailableReplicas int `yaml:"availableReplicas"`
			Conditions        []struct {
				LastTransitionTime time.Time `yaml:"lastTransitionTime"`
				LastUpdateTime     time.Time `yaml:"lastUpdateTime"`
				Message            string    `yaml:"message"`
				Reason             string    `yaml:"reason"`
				Status             string    `yaml:"status"`
				Type               string    `yaml:"type"`
			} `yaml:"conditions"`
			ObservedGeneration int `yaml:"observedGeneration"`
			ReadyReplicas      int `yaml:"readyReplicas"`
			Replicas           int `yaml:"replicas"`
			UpdatedReplicas    int `yaml:"updatedReplicas"`
		} `yaml:"status"`
	} `yaml:"items"`
}

type DEPLOYMENTCONFIG struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Annotations struct {
				OpenshiftIoGeneratedBy string `yaml:"openshift.io/generated-by"`
			} `yaml:"annotations"`
			CreationTimestamp time.Time `yaml:"creationTimestamp"`
			Generation        int       `yaml:"generation"`
			Labels            struct {
				App                      string `yaml:"app"`
				AppKubernetesIoComponent string `yaml:"app.kubernetes.io/component"`
				AppKubernetesIoInstance  string `yaml:"app.kubernetes.io/instance"`
			} `yaml:"labels"`
			Name            string `yaml:"name"`
			Namespace       string `yaml:"namespace"`
			ResourceVersion string `yaml:"resourceVersion"`
			UID             string `yaml:"uid"`
		} `yaml:"metadata"`
		Spec struct {
			Replicas             int `yaml:"replicas"`
			RevisionHistoryLimit int `yaml:"revisionHistoryLimit"`
			Selector             struct {
				Deploymentconfig string `yaml:"deploymentconfig"`
			} `yaml:"selector"`
			Strategy struct {
				ActiveDeadlineSeconds int `yaml:"activeDeadlineSeconds"`
				Resources             struct {
				} `yaml:"resources"`
				RollingParams struct {
					IntervalSeconds     int    `yaml:"intervalSeconds"`
					MaxSurge            string `yaml:"maxSurge"`
					MaxUnavailable      string `yaml:"maxUnavailable"`
					TimeoutSeconds      int    `yaml:"timeoutSeconds"`
					UpdatePeriodSeconds int    `yaml:"updatePeriodSeconds"`
				} `yaml:"rollingParams"`
				Type string `yaml:"type"`
			} `yaml:"strategy"`
			Template struct {
				Metadata struct {
					Annotations struct {
						OpenshiftIoGeneratedBy string `yaml:"openshift.io/generated-by"`
					} `yaml:"annotations"`
					CreationTimestamp interface{} `yaml:"creationTimestamp"`
					Labels            struct {
						Deploymentconfig string `yaml:"deploymentconfig"`
					} `yaml:"labels"`
				} `yaml:"metadata"`
				Spec struct {
					Containers []struct {
						Image           string `yaml:"image"`
						ImagePullPolicy string `yaml:"imagePullPolicy"`
						Name            string `yaml:"name"`
						Ports           []struct {
							ContainerPort int    `yaml:"containerPort"`
							Protocol      string `yaml:"protocol"`
						} `yaml:"ports"`
						Resources struct {
						} `yaml:"resources"`
						TerminationMessagePath   string `yaml:"terminationMessagePath"`
						TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
					} `yaml:"containers"`
					DNSPolicy       string `yaml:"dnsPolicy"`
					RestartPolicy   string `yaml:"restartPolicy"`
					SchedulerName   string `yaml:"schedulerName"`
					SecurityContext struct {
					} `yaml:"securityContext"`
					TerminationGracePeriodSeconds int `yaml:"terminationGracePeriodSeconds"`
				} `yaml:"spec"`
			} `yaml:"template"`
			Test     bool `yaml:"test"`
			Triggers []struct {
				Type              string `yaml:"type"`
				ImageChangeParams struct {
					Automatic      bool     `yaml:"automatic"`
					ContainerNames []string `yaml:"containerNames"`
					From           struct {
						Kind      string `yaml:"kind"`
						Name      string `yaml:"name"`
						Namespace string `yaml:"namespace"`
					} `yaml:"from"`
					LastTriggeredImage string `yaml:"lastTriggeredImage"`
				} `yaml:"imageChangeParams,omitempty"`
			} `yaml:"triggers"`
		} `yaml:"spec"`
		Status struct {
			AvailableReplicas int `yaml:"availableReplicas"`
			Conditions        []struct {
				LastTransitionTime time.Time `yaml:"lastTransitionTime"`
				LastUpdateTime     time.Time `yaml:"lastUpdateTime"`
				Message            string    `yaml:"message"`
				Reason             string    `yaml:"reason,omitempty"`
				Status             string    `yaml:"status"`
				Type               string    `yaml:"type"`
			} `yaml:"conditions"`
			Details struct {
				Causes []struct {
					Type string `yaml:"type"`
				} `yaml:"causes"`
				Message string `yaml:"message"`
			} `yaml:"details"`
			LatestVersion       int `yaml:"latestVersion"`
			ObservedGeneration  int `yaml:"observedGeneration"`
			ReadyReplicas       int `yaml:"readyReplicas"`
			Replicas            int `yaml:"replicas"`
			UnavailableReplicas int `yaml:"unavailableReplicas"`
			UpdatedReplicas     int `yaml:"updatedReplicas"`
		} `yaml:"status"`
	} `yaml:"items"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
}

type DAEMONSET struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Annotations struct {
				DeprecatedDaemonsetTemplateGeneration string `yaml:"deprecated.daemonset.template.generation"`
			} `yaml:"annotations"`
			CreationTimestamp time.Time `yaml:"creationTimestamp"`
			Generation        int       `yaml:"generation"`
			Labels            struct {
				DNSOperatorOpenshiftIoOwningDNS string `yaml:"dns.operator.openshift.io/owning-dns"`
			} `yaml:"labels"`
			Name            string `yaml:"name"`
			Namespace       string `yaml:"namespace"`
			OwnerReferences []struct {
				APIVersion string `yaml:"apiVersion"`
				Controller bool   `yaml:"controller"`
				Kind       string `yaml:"kind"`
				Name       string `yaml:"name"`
				UID        string `yaml:"uid"`
			} `yaml:"ownerReferences"`
			ResourceVersion string `yaml:"resourceVersion"`
			SelfLink        string `yaml:"selfLink"`
			UID             string `yaml:"uid"`
		} `yaml:"metadata"`
		Spec struct {
			RevisionHistoryLimit int `yaml:"revisionHistoryLimit"`
			Selector             struct {
				MatchLabels struct {
					DNSOperatorOpenshiftIoDaemonsetDNS string `yaml:"dns.operator.openshift.io/daemonset-dns"`
				} `yaml:"matchLabels"`
			} `yaml:"selector"`
			Template struct {
				Metadata struct {
					CreationTimestamp interface{} `yaml:"creationTimestamp"`
					Labels            struct {
						DNSOperatorOpenshiftIoDaemonsetDNS string `yaml:"dns.operator.openshift.io/daemonset-dns"`
					} `yaml:"labels"`
				} `yaml:"metadata"`
				Spec struct {
					Containers []struct {
						Args            []string `yaml:"args,omitempty"`
						Command         []string `yaml:"command,omitempty"`
						Image           string   `yaml:"image"`
						ImagePullPolicy string   `yaml:"imagePullPolicy"`
						LivenessProbe   struct {
							FailureThreshold int `yaml:"failureThreshold"`
							HTTPGet          struct {
								Path   string `yaml:"path"`
								Port   int    `yaml:"port"`
								Scheme string `yaml:"scheme"`
							} `yaml:"httpGet"`
							InitialDelaySeconds int `yaml:"initialDelaySeconds"`
							PeriodSeconds       int `yaml:"periodSeconds"`
							SuccessThreshold    int `yaml:"successThreshold"`
							TimeoutSeconds      int `yaml:"timeoutSeconds"`
						} `yaml:"livenessProbe,omitempty"`
						Name  string `yaml:"name"`
						Ports []struct {
							ContainerPort int    `yaml:"containerPort"`
							Name          string `yaml:"name"`
							Protocol      string `yaml:"protocol"`
						} `yaml:"ports,omitempty"`
						ReadinessProbe struct {
							FailureThreshold int `yaml:"failureThreshold"`
							HTTPGet          struct {
								Path   string `yaml:"path"`
								Port   int    `yaml:"port"`
								Scheme string `yaml:"scheme"`
							} `yaml:"httpGet"`
							InitialDelaySeconds int `yaml:"initialDelaySeconds"`
							PeriodSeconds       int `yaml:"periodSeconds"`
							SuccessThreshold    int `yaml:"successThreshold"`
							TimeoutSeconds      int `yaml:"timeoutSeconds"`
						} `yaml:"readinessProbe,omitempty"`
						Resources struct {
							Limits struct {
								Memory string `yaml:"memory"`
							} `yaml:"limits"`
							Requests struct {
								CPU    string `yaml:"cpu"`
								Memory string `yaml:"memory"`
							} `yaml:"requests"`
						} `yaml:"resources,omitempty"`
						TerminationMessagePath   string `yaml:"terminationMessagePath"`
						TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
						VolumeMounts             []struct {
							MountPath string `yaml:"mountPath"`
							Name      string `yaml:"name"`
							ReadOnly  bool   `yaml:"readOnly"`
						} `yaml:"volumeMounts"`
						Env []struct {
							Name  string `yaml:"name"`
							Value string `yaml:"value"`
						} `yaml:"env,omitempty"`
						SecurityContext struct {
							Privileged bool `yaml:"privileged"`
						} `yaml:"securityContext,omitempty"`
					} `yaml:"containers"`
					DNSPolicy         string            `yaml:"dnsPolicy"`
					NodeSelector      map[string]string `yaml:"nodeSelector"`
					PriorityClassName string            `yaml:"priorityClassName"`
					RestartPolicy     string            `yaml:"restartPolicy"`
					SchedulerName     string            `yaml:"schedulerName"`
					SecurityContext   struct {
					} `yaml:"securityContext"`
					ServiceAccount                string `yaml:"serviceAccount"`
					ServiceAccountName            string `yaml:"serviceAccountName"`
					TerminationGracePeriodSeconds int    `yaml:"terminationGracePeriodSeconds"`
					Tolerations                   []struct {
						Operator string `yaml:"operator"`
					} `yaml:"tolerations"`
					Volumes []struct {
						ConfigMap struct {
							DefaultMode int `yaml:"defaultMode"`
							Items       []struct {
								Key  string `yaml:"key"`
								Path string `yaml:"path"`
							} `yaml:"items"`
							Name string `yaml:"name"`
						} `yaml:"configMap,omitempty"`
						Name     string `yaml:"name"`
						HostPath struct {
							Path string `yaml:"path"`
							Type string `yaml:"type"`
						} `yaml:"hostPath,omitempty"`
						Secret struct {
							DefaultMode int    `yaml:"defaultMode"`
							SecretName  string `yaml:"secretName"`
						} `yaml:"secret,omitempty"`
					} `yaml:"volumes"`
				} `yaml:"spec"`
			} `yaml:"template"`
			UpdateStrategy struct {
				RollingUpdate struct {
					MaxUnavailable string `yaml:"maxUnavailable"`
				} `yaml:"rollingUpdate"`
				Type string `yaml:"type"`
			} `yaml:"updateStrategy"`
		} `yaml:"spec"`
		Status struct {
			CurrentNumberScheduled int `yaml:"currentNumberScheduled"`
			DesiredNumberScheduled int `yaml:"desiredNumberScheduled"`
			NumberAvailable        int `yaml:"numberAvailable"`
			NumberMisscheduled     int `yaml:"numberMisscheduled"`
			NumberReady            int `yaml:"numberReady"`
			ObservedGeneration     int `yaml:"observedGeneration"`
			UpdatedNumberScheduled int `yaml:"updatedNumberScheduled"`
		} `yaml:"status"`
	} `yaml:"items"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
}

type SERVICES struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []struct {
		APIVersion string `yaml:"apiVersion"`
		Kind       string `yaml:"kind"`
		Metadata   struct {
			Annotations struct {
				ServiceAlphaOpenshiftIoServingCertSignedBy  string `yaml:"service.alpha.openshift.io/serving-cert-signed-by"`
				ServiceBetaOpenshiftIoServingCertSecretName string `yaml:"service.beta.openshift.io/serving-cert-secret-name"`
				ServiceBetaOpenshiftIoServingCertSignedBy   string `yaml:"service.beta.openshift.io/serving-cert-signed-by"`
			} `yaml:"annotations"`
			CreationTimestamp time.Time `yaml:"creationTimestamp"`
			Labels            struct {
				App string `yaml:"app"`
			} `yaml:"labels"`
			Name            string `yaml:"name"`
			Namespace       string `yaml:"namespace"`
			OwnerReferences []struct {
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
			ClusterIP  string   `yaml:"clusterIP"`
			ClusterIPs []string `yaml:"clusterIPs"`
			Ports      []struct {
				Name       string `yaml:"name"`
				Port       int    `yaml:"port"`
				Protocol   string `yaml:"protocol"`
				TargetPort int    `yaml:"targetPort"`
			} `yaml:"ports"`
			PublishNotReadyAddresses bool `yaml:"publishNotReadyAddresses"`
			Selector                 struct {
				App string `yaml:"app"`
			} `yaml:"selector"`
			SessionAffinity string `yaml:"sessionAffinity"`
			Type            string `yaml:"type"`
		} `yaml:"spec"`
		Status struct {
			LoadBalancer struct {
			} `yaml:"loadBalancer"`
		} `yaml:"status"`
	} `yaml:"items"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
}

type ROUTE struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations       map[string]string `yaml:"annotations"`
		CreationTimestamp time.Time         `yaml:"creationTimestamp"`
		Name              string            `yaml:"name"`
		Namespace         string            `yaml:"namespace"`
		ResourceVersion   string            `yaml:"resourceVersion"`
		UID               string            `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Host string `yaml:"host"`
		Port struct {
			TargetPort string `yaml:"targetPort"`
		} `yaml:"port"`
		TLS struct {
			InsecureEdgeTerminationPolicy string `yaml:"insecureEdgeTerminationPolicy"`
			Termination                   string `yaml:"termination"`
		} `yaml:"tls"`
		To struct {
			Kind   string `yaml:"kind"`
			Name   string `yaml:"name"`
			Weight int    `yaml:"weight"`
		} `yaml:"to"`
		WildcardPolicy string `yaml:"wildcardPolicy"`
	} `yaml:"spec"`
	Status struct {
		Ingress []struct {
			Conditions []struct {
				LastTransitionTime time.Time `yaml:"lastTransitionTime"`
				Status             string    `yaml:"status"`
				Type               string    `yaml:"type"`
			} `yaml:"conditions"`
			Host                    string `yaml:"host"`
			RouterCanonicalHostname string `yaml:"routerCanonicalHostname"`
			RouterName              string `yaml:"routerName"`
			WildcardPolicy          string `yaml:"wildcardPolicy"`
		} `yaml:"ingress"`
	} `yaml:"status"`
}

type ROUTES struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []ROUTE
	Kind       string `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
}
