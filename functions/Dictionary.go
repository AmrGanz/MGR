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
