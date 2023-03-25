package functions

import (
	"encoding/json"
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
var MGTimeStamp string

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
var All_Operators_Path = ""
var Operator_Path = ""
var InstalledOperators_Path = ""
var InstallPlans_Path = ""
var MCP_Path = ""
var MC_Path = ""
var PV_Path = ""
var CSR_Path = ""
var ETCD_Path = ""

// Resources Paths
func SetResourcesPath() {
	Version_Path = MG_Path + "cluster-scoped-resources/config.openshift.io/clusterversions/version.yaml"
	Configurations_Path = MG_Path + "cluster-scoped-resources/config.openshift.io/"
	Namespaces_Path = MG_Path + "namespaces/"
	Nodes_Path = MG_Path + "cluster-scoped-resources/core/nodes/"
	All_Operators_Path = MG_Path + "cluster-scoped-resources/config.openshift.io/clusteroperators.yaml"
	Operator_Path = MG_Path + "cluster-scoped-resources/config.openshift.io/clusteroperators/"
	InstalledOperators_Path = MG_Path + "/cluster-scoped-resources/operators.coreos.com/operators/"
	InstallPlans_Path = "/operators.coreos.com/installplans/"
	MCP_Path = MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigpools/"
	MC_Path = MG_Path + "cluster-scoped-resources/machineconfiguration.openshift.io/machineconfigs/"
	PV_Path = MG_Path + "cluster-scoped-resources/core/persistentvolumes/"
	CSR_Path = MG_Path + "cluster-scoped-resources/certificates.k8s.io/certificatesigningrequests/"
	ETCD_Path = MG_Path + "etcd_info/"
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

type CONFIGMAP struct {
	APIVersion string            `yaml:"apiVersion"`
	Data       map[string]string `yaml:"data"`
	Kind       string            `yaml:"kind"`
	Metadata   struct {
		Annotations       map[string]string `yaml:"annotations"`
		CreationTimestamp time.Time         `yaml:"creationTimestamp"`
		Name              string            `yaml:"name"`
		Namespace         string            `yaml:"namespace"`
		ResourceVersion   string            `yaml:"resourceVersion"`
		UID               string            `yaml:"uid"`
	} `yaml:"metadata"`
}

type CONFIGMAPS struct {
	APIVersion string      `yaml:"apiVersion"`
	Items      []CONFIGMAP `yaml:"items"`
	Kind       string      `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
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

type CSV struct {
	Kind     string `yaml:"kind"`
	Metadata struct {
		Annotations struct {
			AlmExamples                                 string    `yaml:"alm-examples"`
			Capabilities                                string    `yaml:"capabilities"`
			Categories                                  string    `yaml:"categories"`
			Certified                                   string    `yaml:"certified"`
			CreatedAt                                   time.Time `yaml:"createdAt"`
			Description                                 string    `yaml:"description"`
			OlmOperatorGroup                            string    `yaml:"olm.operatorGroup"`
			OlmOperatorNamespace                        string    `yaml:"olm.operatorNamespace"`
			OlmSkipRange                                string    `yaml:"olm.skipRange"`
			OlmTargetNamespaces                         string    `yaml:"olm.targetNamespaces"`
			OperatorframeworkIoInitializationResource   string    `yaml:"operatorframework.io/initialization-resource"`
			OperatorframeworkIoProperties               string    `yaml:"operatorframework.io/properties"`
			OperatorframeworkIoSuggestedNamespace       string    `yaml:"operatorframework.io/suggested-namespace"`
			OperatorsOpenshiftIoInfrastructureFeatures  string    `yaml:"operators.openshift.io/infrastructure-features"`
			OperatorsOpenshiftIoValidSubscription       string    `yaml:"operators.openshift.io/valid-subscription"`
			OperatorsOperatorframeworkIoInternalObjects string    `yaml:"operators.operatorframework.io/internal-objects"`
			Support                                     string    `yaml:"support"`
		} `yaml:"annotations"`
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Generation        int       `yaml:"generation"`
		Labels            struct {
			OperatorframeworkIoArchAmd64                           string `yaml:"operatorframework.io/arch.amd64"`
			OperatorframeworkIoArchArm64                           string `yaml:"operatorframework.io/arch.arm64"`
			OperatorframeworkIoArchPpc64Le                         string `yaml:"operatorframework.io/arch.ppc64le"`
			OperatorframeworkIoArchS390X                           string `yaml:"operatorframework.io/arch.s390x"`
			OperatorframeworkIoOsLinux                             string `yaml:"operatorframework.io/os.linux"`
			OperatorsCoreosComMulticlusterEngineMulticlusterEngine string `yaml:"operators.coreos.com/multicluster-engine.multicluster-engine"`
		} `yaml:"labels"`
		Name            string `yaml:"name"`
		Namespace       string `yaml:"namespace"`
		ResourceVersion string `yaml:"resourceVersion"`
		UID             string `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Apiservicedefinitions struct {
		} `yaml:"apiservicedefinitions"`
		Cleanup struct {
			Enabled bool `yaml:"enabled"`
		} `yaml:"cleanup"`
		Customresourcedefinitions struct {
			Owned []struct {
				Description     string `yaml:"description"`
				DisplayName     string `yaml:"displayName"`
				Kind            string `yaml:"kind"`
				Name            string `yaml:"name"`
				SpecDescriptors []struct {
					Description  string   `yaml:"description"`
					DisplayName  string   `yaml:"displayName"`
					Path         string   `yaml:"path"`
					XDescriptors []string `yaml:"x-descriptors"`
				} `yaml:"specDescriptors"`
				Version string `yaml:"version"`
			} `yaml:"owned"`
		} `yaml:"customresourcedefinitions"`
		Description string `yaml:"description"`
		DisplayName string `yaml:"displayName"`
		Icon        []struct {
			Base64Data string `yaml:"base64data"`
			Mediatype  string `yaml:"mediatype"`
		} `yaml:"icon"`
		Install struct {
			Spec struct {
				ClusterPermissions []struct {
					Rules []struct {
						APIGroups []string `yaml:"apiGroups"`
						Resources []string `yaml:"resources"`
						Verbs     []string `yaml:"verbs"`
					} `yaml:"rules"`
					ServiceAccountName string `yaml:"serviceAccountName"`
				} `yaml:"clusterPermissions"`
				Deployments []struct {
					Label struct {
						ControlPlane string `yaml:"control-plane"`
					} `yaml:"label"`
					Name string `yaml:"name"`
					Spec struct {
						Replicas int `yaml:"replicas"`
						Selector struct {
							MatchLabels struct {
								ControlPlane string `yaml:"control-plane"`
							} `yaml:"matchLabels"`
						} `yaml:"selector"`
						Strategy struct {
						} `yaml:"strategy"`
						Template struct {
							Metadata struct {
								CreationTimestamp interface{} `yaml:"creationTimestamp"`
								Labels            struct {
									ControlPlane string `yaml:"control-plane"`
								} `yaml:"labels"`
							} `yaml:"metadata"`
							Spec struct {
								Containers []struct {
									Args    []string `yaml:"args"`
									Command []string `yaml:"command"`
									Env     []struct {
										Name      string `yaml:"name"`
										ValueFrom struct {
											FieldRef struct {
												FieldPath string `yaml:"fieldPath"`
											} `yaml:"fieldRef"`
										} `yaml:"valueFrom,omitempty"`
										Value string `yaml:"value,omitempty"`
									} `yaml:"env"`
									Image         string `yaml:"image"`
									LivenessProbe struct {
										HTTPGet struct {
											Path string `yaml:"path"`
											Port int    `yaml:"port"`
										} `yaml:"httpGet"`
										InitialDelaySeconds int `yaml:"initialDelaySeconds"`
										PeriodSeconds       int `yaml:"periodSeconds"`
									} `yaml:"livenessProbe"`
									Name           string `yaml:"name"`
									ReadinessProbe struct {
										HTTPGet struct {
											Path string `yaml:"path"`
											Port int    `yaml:"port"`
										} `yaml:"httpGet"`
										InitialDelaySeconds int `yaml:"initialDelaySeconds"`
										PeriodSeconds       int `yaml:"periodSeconds"`
									} `yaml:"readinessProbe"`
									Resources struct {
										Limits struct {
											CPU    string `yaml:"cpu"`
											Memory string `yaml:"memory"`
										} `yaml:"limits"`
										Requests struct {
											CPU    string `yaml:"cpu"`
											Memory string `yaml:"memory"`
										} `yaml:"requests"`
									} `yaml:"resources"`
									SecurityContext struct {
										AllowPrivilegeEscalation bool `yaml:"allowPrivilegeEscalation"`
									} `yaml:"securityContext"`
									VolumeMounts []struct {
										MountPath string `yaml:"mountPath"`
										Name      string `yaml:"name"`
										ReadOnly  bool   `yaml:"readOnly"`
									} `yaml:"volumeMounts"`
								} `yaml:"containers"`
								SecurityContext struct {
									RunAsNonRoot bool `yaml:"runAsNonRoot"`
								} `yaml:"securityContext"`
								ServiceAccountName            string `yaml:"serviceAccountName"`
								TerminationGracePeriodSeconds int    `yaml:"terminationGracePeriodSeconds"`
								Volumes                       []struct {
									Name   string `yaml:"name"`
									Secret struct {
										DefaultMode int    `yaml:"defaultMode"`
										SecretName  string `yaml:"secretName"`
									} `yaml:"secret"`
								} `yaml:"volumes"`
							} `yaml:"spec"`
						} `yaml:"template"`
					} `yaml:"spec"`
				} `yaml:"deployments"`
				Permissions []struct {
					Rules []struct {
						APIGroups []string `yaml:"apiGroups"`
						Resources []string `yaml:"resources"`
						Verbs     []string `yaml:"verbs"`
					} `yaml:"rules"`
					ServiceAccountName string `yaml:"serviceAccountName"`
				} `yaml:"permissions"`
			} `yaml:"spec"`
			Strategy string `yaml:"strategy"`
		} `yaml:"install"`
		InstallModes []struct {
			Supported bool   `yaml:"supported"`
			Type      string `yaml:"type"`
		} `yaml:"installModes"`
		Keywords []string `yaml:"keywords"`
		Labels   struct {
			App string `yaml:"app"`
		} `yaml:"labels"`
		Links []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		} `yaml:"links"`
		Maintainers []struct {
			Email string `yaml:"email"`
			Name  string `yaml:"name"`
		} `yaml:"maintainers"`
		Maturity string `yaml:"maturity"`
		Provider struct {
			Name string `yaml:"name"`
		} `yaml:"provider"`
		RelatedImages []struct {
			Image string `yaml:"image"`
			Name  string `yaml:"name"`
		} `yaml:"relatedImages"`
		Replaces string `yaml:"replaces"`
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Version string `yaml:"version"`
	} `yaml:"spec"`
	Status struct {
		Cleanup struct {
		} `yaml:"cleanup"`
		Conditions []struct {
			LastTransitionTime time.Time `yaml:"lastTransitionTime"`
			LastUpdateTime     time.Time `yaml:"lastUpdateTime"`
			Message            string    `yaml:"message"`
			Phase              string    `yaml:"phase"`
			Reason             string    `yaml:"reason"`
		} `yaml:"conditions"`
		LastTransitionTime time.Time `yaml:"lastTransitionTime"`
		LastUpdateTime     time.Time `yaml:"lastUpdateTime"`
		Message            string    `yaml:"message"`
		Phase              string    `yaml:"phase"`
		Reason             string    `yaml:"reason"`
		RequirementStatus  []struct {
			Group      string `yaml:"group"`
			Kind       string `yaml:"kind"`
			Message    string `yaml:"message"`
			Name       string `yaml:"name"`
			Status     string `yaml:"status"`
			UUID       string `yaml:"uuid,omitempty"`
			Version    string `yaml:"version"`
			Dependents []struct {
				Group   string `yaml:"group"`
				Kind    string `yaml:"kind"`
				Message string `yaml:"message"`
				Status  string `yaml:"status"`
				Version string `yaml:"version"`
			} `yaml:"dependents,omitempty"`
		} `yaml:"requirementStatus"`
	} `yaml:"status"`
}

type DEPLOYMENT struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations struct {
			DeploymentKubernetesIoRevision string `yaml:"deployment.kubernetes.io/revision"`
		} `yaml:"annotations"`
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Generation        int       `yaml:"generation"`
		Labels            struct {
			BackplaneconfigName string `yaml:"backplaneconfig.name"`
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
		UID             string `yaml:"uid"`
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
				Affinity struct {
					NodeAffinity struct {
						RequiredDuringSchedulingIgnoredDuringExecution struct {
							NodeSelectorTerms []struct {
								MatchExpressions []struct {
									Key      string   `yaml:"key"`
									Operator string   `yaml:"operator"`
									Values   []string `yaml:"values"`
								} `yaml:"matchExpressions"`
							} `yaml:"nodeSelectorTerms"`
						} `yaml:"requiredDuringSchedulingIgnoredDuringExecution"`
					} `yaml:"nodeAffinity"`
					PodAntiAffinity struct {
						PreferredDuringSchedulingIgnoredDuringExecution []struct {
							PodAffinityTerm struct {
								LabelSelector struct {
									MatchExpressions []struct {
										Key      string   `yaml:"key"`
										Operator string   `yaml:"operator"`
										Values   []string `yaml:"values"`
									} `yaml:"matchExpressions"`
								} `yaml:"labelSelector"`
								TopologyKey string `yaml:"topologyKey"`
							} `yaml:"podAffinityTerm"`
							Weight int `yaml:"weight"`
						} `yaml:"preferredDuringSchedulingIgnoredDuringExecution"`
					} `yaml:"podAntiAffinity"`
				} `yaml:"affinity"`
				Containers []struct {
					Command []string `yaml:"command"`
					Env     []struct {
						Name      string `yaml:"name"`
						ValueFrom struct {
							FieldRef struct {
								APIVersion string `yaml:"apiVersion"`
								FieldPath  string `yaml:"fieldPath"`
							} `yaml:"fieldRef"`
						} `yaml:"valueFrom"`
					} `yaml:"env"`
					Image           string `yaml:"image"`
					ImagePullPolicy string `yaml:"imagePullPolicy"`
					Name            string `yaml:"name"`
					Resources       struct {
						Limits struct {
							CPU    string `yaml:"cpu"`
							Memory string `yaml:"memory"`
						} `yaml:"limits"`
						Requests struct {
							CPU    string `yaml:"cpu"`
							Memory string `yaml:"memory"`
						} `yaml:"requests"`
					} `yaml:"resources"`
					SecurityContext struct {
						AllowPrivilegeEscalation bool `yaml:"allowPrivilegeEscalation"`
						Capabilities             struct {
							Drop []string `yaml:"drop"`
						} `yaml:"capabilities"`
						Privileged             bool `yaml:"privileged"`
						ReadOnlyRootFilesystem bool `yaml:"readOnlyRootFilesystem"`
					} `yaml:"securityContext"`
					TerminationMessagePath   string `yaml:"terminationMessagePath"`
					TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
				} `yaml:"containers"`
				DNSPolicy       string `yaml:"dnsPolicy"`
				RestartPolicy   string `yaml:"restartPolicy"`
				SchedulerName   string `yaml:"schedulerName"`
				SecurityContext struct {
					RunAsNonRoot bool `yaml:"runAsNonRoot"`
				} `yaml:"securityContext"`
				ServiceAccount                string `yaml:"serviceAccount"`
				ServiceAccountName            string `yaml:"serviceAccountName"`
				TerminationGracePeriodSeconds int    `yaml:"terminationGracePeriodSeconds"`
				Tolerations                   []struct {
					Effect   string `yaml:"effect"`
					Key      string `yaml:"key"`
					Operator string `yaml:"operator"`
				} `yaml:"tolerations"`
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
}

type DEPLOYMENTS struct {
	APIVersion string       `yaml:"apiVersion"`
	Items      []DEPLOYMENT `yaml:"items"`
	Kind       string       `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
	} `yaml:"metadata"`
}

type DEPLOYMENTCONFIG struct {
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
}

type DEPLOYMENTCONFIGS struct {
	APIVersion string             `yaml:"apiVersion"`
	Items      []DEPLOYMENTCONFIG `yaml:"items"`
	Kind       string             `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
	} `yaml:"metadata"`
}

type DAEMONSET struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations struct {
			DeprecatedDaemonsetTemplateGeneration string `yaml:"deprecated.daemonset.template.generation"`
		} `yaml:"annotations"`
		CreationTimestamp time.Time         `yaml:"creationTimestamp"`
		Generation        int               `yaml:"generation"`
		Labels            map[string]string `yaml:"labels"`
		Name              string            `yaml:"name"`
		Namespace         string            `yaml:"namespace"`
		ResourceVersion   string            `yaml:"resourceVersion"`
		UID               string            `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		RevisionHistoryLimit int `yaml:"revisionHistoryLimit"`
		Selector             struct {
			MatchLabels struct {
				AppKubernetesIoComponent string `yaml:"app.kubernetes.io/component"`
				AppKubernetesIoName      string `yaml:"app.kubernetes.io/name"`
				AppKubernetesIoPartOf    string `yaml:"app.kubernetes.io/part-of"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template struct {
			Metadata struct {
				Annotations struct {
					KubectlKubernetesIoDefaultContainer string `yaml:"kubectl.kubernetes.io/default-container"`
					TargetWorkloadOpenshiftIoManagement string `yaml:"target.workload.openshift.io/management"`
				} `yaml:"annotations"`
				CreationTimestamp interface{} `yaml:"creationTimestamp"`
				Labels            struct {
					AppKubernetesIoComponent string `yaml:"app.kubernetes.io/component"`
					AppKubernetesIoManagedBy string `yaml:"app.kubernetes.io/managed-by"`
					AppKubernetesIoName      string `yaml:"app.kubernetes.io/name"`
					AppKubernetesIoPartOf    string `yaml:"app.kubernetes.io/part-of"`
					AppKubernetesIoVersion   string `yaml:"app.kubernetes.io/version"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				AutomountServiceAccountToken bool `yaml:"automountServiceAccountToken"`
				Containers                   []struct {
					Args            []string `yaml:"args"`
					Image           string   `yaml:"image"`
					ImagePullPolicy string   `yaml:"imagePullPolicy"`
					Name            string   `yaml:"name"`
					Resources       struct {
						Requests struct {
							CPU    string `yaml:"cpu"`
							Memory string `yaml:"memory"`
						} `yaml:"requests"`
					} `yaml:"resources"`
					TerminationMessagePath   string `yaml:"terminationMessagePath"`
					TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
					VolumeMounts             []struct {
						MountPath        string `yaml:"mountPath"`
						MountPropagation string `yaml:"mountPropagation,omitempty"`
						Name             string `yaml:"name"`
						ReadOnly         bool   `yaml:"readOnly"`
					} `yaml:"volumeMounts"`
					WorkingDir string `yaml:"workingDir,omitempty"`
					Env        []struct {
						Name      string `yaml:"name"`
						ValueFrom struct {
							FieldRef struct {
								APIVersion string `yaml:"apiVersion"`
								FieldPath  string `yaml:"fieldPath"`
							} `yaml:"fieldRef"`
						} `yaml:"valueFrom"`
					} `yaml:"env,omitempty"`
					Ports []struct {
						ContainerPort int    `yaml:"containerPort"`
						HostPort      int    `yaml:"hostPort"`
						Name          string `yaml:"name"`
						Protocol      string `yaml:"protocol"`
					} `yaml:"ports,omitempty"`
					SecurityContext struct {
						AllowPrivilegeEscalation bool `yaml:"allowPrivilegeEscalation"`
						Capabilities             struct {
							Drop []string `yaml:"drop"`
						} `yaml:"capabilities"`
						ReadOnlyRootFilesystem bool `yaml:"readOnlyRootFilesystem"`
						RunAsGroup             int  `yaml:"runAsGroup"`
						RunAsNonRoot           bool `yaml:"runAsNonRoot"`
						RunAsUser              int  `yaml:"runAsUser"`
					} `yaml:"securityContext,omitempty"`
				} `yaml:"containers"`
				DNSPolicy      string `yaml:"dnsPolicy"`
				HostNetwork    bool   `yaml:"hostNetwork"`
				HostPID        bool   `yaml:"hostPID"`
				InitContainers []struct {
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
					SecurityContext struct {
						Privileged bool `yaml:"privileged"`
						RunAsUser  int  `yaml:"runAsUser"`
					} `yaml:"securityContext"`
					TerminationMessagePath   string `yaml:"terminationMessagePath"`
					TerminationMessagePolicy string `yaml:"terminationMessagePolicy"`
					VolumeMounts             []struct {
						MountPath string `yaml:"mountPath"`
						Name      string `yaml:"name"`
						ReadOnly  bool   `yaml:"readOnly,omitempty"`
					} `yaml:"volumeMounts"`
					WorkingDir string `yaml:"workingDir"`
				} `yaml:"initContainers"`
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
					HostPath struct {
						Path string `yaml:"path"`
						Type string `yaml:"type"`
					} `yaml:"hostPath,omitempty"`
					Name     string `yaml:"name"`
					EmptyDir struct {
					} `yaml:"emptyDir,omitempty"`
					Secret struct {
						DefaultMode int    `yaml:"defaultMode"`
						SecretName  string `yaml:"secretName"`
					} `yaml:"secret,omitempty"`
					ConfigMap struct {
						DefaultMode int    `yaml:"defaultMode"`
						Name        string `yaml:"name"`
					} `yaml:"configMap,omitempty"`
				} `yaml:"volumes"`
			} `yaml:"spec"`
		} `yaml:"template"`
		UpdateStrategy struct {
			RollingUpdate struct {
				MaxSurge       int    `yaml:"maxSurge"`
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
}

type DAEMONSETS struct {
	APIVersion string      `yaml:"apiVersion"`
	Items      []DAEMONSET `yaml:"items"`
	Kind       string      `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
	} `yaml:"metadata"`
}

type EVENTS struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []struct {
		APIVersion     string      `yaml:"apiVersion"`
		Count          int         `yaml:"count"`
		EventTime      interface{} `yaml:"eventTime"`
		FirstTimestamp time.Time   `yaml:"firstTimestamp"`
		InvolvedObject struct {
			APIVersion      string `yaml:"apiVersion"`
			FieldPath       string `yaml:"fieldPath"`
			Kind            string `yaml:"kind"`
			Name            string `yaml:"name"`
			Namespace       string `yaml:"namespace"`
			ResourceVersion string `yaml:"resourceVersion"`
			UID             string `yaml:"uid"`
		} `yaml:"involvedObject"`
		Kind          string    `yaml:"kind"`
		LastTimestamp time.Time `yaml:"lastTimestamp"`
		Message       string    `yaml:"message"`
		Metadata      struct {
			CreationTimestamp time.Time `yaml:"creationTimestamp"`
			Name              string    `yaml:"name"`
			Namespace         string    `yaml:"namespace"`
			ResourceVersion   string    `yaml:"resourceVersion"`
			SelfLink          string    `yaml:"selfLink"`
			UID               string    `yaml:"uid"`
		} `yaml:"metadata"`
		Reason             string `yaml:"reason"`
		ReportingComponent string `yaml:"reportingComponent"`
		ReportingInstance  string `yaml:"reportingInstance"`
		Source             struct {
			Component string `yaml:"component"`
			Host      string `yaml:"host"`
		} `yaml:"source"`
		Type string `yaml:"type"`
	} `yaml:"items"`
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

type IMAGESTREAM struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations struct {
			OpenshiftIoDisplayName                string    `yaml:"openshift.io/display-name"`
			OpenshiftIoImageDockerRepositoryCheck time.Time `yaml:"openshift.io/image.dockerRepositoryCheck"`
			SamplesOperatorOpenshiftIoVersion     string    `yaml:"samples.operator.openshift.io/version"`
		} `yaml:"annotations"`
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Generation        int       `yaml:"generation"`
		Labels            struct {
			SamplesOperatorOpenshiftIoManaged string `yaml:"samples.operator.openshift.io/managed"`
		} `yaml:"labels"`
		Name            string `yaml:"name"`
		Namespace       string `yaml:"namespace"`
		ResourceVersion string `yaml:"resourceVersion"`
		UID             string `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		LookupPolicy struct {
			Local bool `yaml:"local"`
		} `yaml:"lookupPolicy"`
		Tags []struct {
			Annotations struct {
				Description                    string `yaml:"description"`
				IconClass                      string `yaml:"iconClass"`
				OpenshiftIoDisplayName         string `yaml:"openshift.io/display-name"`
				OpenshiftIoProviderDisplayName string `yaml:"openshift.io/provider-display-name"`
				SampleRepo                     string `yaml:"sampleRepo"`
				Supports                       string `yaml:"supports"`
				Tags                           string `yaml:"tags"`
				Version                        string `yaml:"version"`
			} `yaml:"annotations,omitempty"`
			From struct {
				Kind string `yaml:"kind"`
				Name string `yaml:"name"`
			} `yaml:"from"`
			Generation   int `yaml:"generation"`
			ImportPolicy struct {
			} `yaml:"importPolicy"`
			Name            string `yaml:"name"`
			ReferencePolicy struct {
				Type string `yaml:"type"`
			} `yaml:"referencePolicy"`
		} `yaml:"tags"`
	} `yaml:"spec"`
	Status struct {
		DockerImageRepository string `yaml:"dockerImageRepository"`
		Tags                  []struct {
			Items []struct {
				Created              time.Time `yaml:"created"`
				DockerImageReference string    `yaml:"dockerImageReference"`
				Generation           int       `yaml:"generation"`
				Image                string    `yaml:"image"`
			} `yaml:"items"`
			Tag string `yaml:"tag"`
		} `yaml:"tags"`
	} `yaml:"status"`
}

type IMAGESTREAMS struct {
	APIVersion string        `yaml:"apiVersion"`
	Items      []IMAGESTREAM `yaml:"items"`
	Kind       string        `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
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

type NODES struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []NODE `yaml:"items"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
}

type CLUSTEROPERATOR struct {
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
		phase string `yaml:"phase"`
	} `yaml:"status"`
}

type CLUSTEROPERATORS struct {
	APIVersion string            `yaml:"apiVersion"`
	Items      []CLUSTEROPERATOR `yaml:"items"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
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
	Items      []POD  `yaml:"items"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
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

type PVC struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Annotations       map[string]string `yaml:"annotations"`
		CreationTimestamp time.Time         `yaml:"creationTimestamp"`
		Finalizers        []string          `yaml:"finalizers"`
		Name              string            `yaml:"name"`
		Namespace         string            `yaml:"namespace"`
		ResourceVersion   string            `yaml:"resourceVersion"`
		UID               string            `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		AccessModes []string `yaml:"accessModes"`
		Resources   struct {
			Requests struct {
				Storage string `yaml:"storage"`
			} `yaml:"requests"`
		} `yaml:"resources"`
		StorageClassName string `yaml:"storageClassName"`
		VolumeMode       string `yaml:"volumeMode"`
		VolumeName       string `yaml:"volumeName"`
	} `yaml:"spec"`
	Status struct {
		AccessModes []string `yaml:"accessModes"`
		Capacity    struct {
			Storage string `yaml:"storage"`
		} `yaml:"capacity"`
		Phase string `yaml:"phase"`
	} `yaml:"status"`
}

type PVCS struct {
	APIVersion string `yaml:"apiVersion"`
	Items      []PVC  `yaml:"items"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
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
		Path string `yaml:"path"`
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
	APIVersion string  `yaml:"apiVersion"`
	Items      []ROUTE `yaml:"items"`
	Kind       string  `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
}

type SERVICE struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Labels            struct {
			Component string `yaml:"component"`
			Provider  string `yaml:"provider"`
		} `yaml:"labels"`
		Name            string `yaml:"name"`
		Namespace       string `yaml:"namespace"`
		ResourceVersion string `yaml:"resourceVersion"`
		UID             string `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterIP             string   `yaml:"clusterIP"`
		ClusterIPs            []string `yaml:"clusterIPs"`
		InternalTrafficPolicy string   `yaml:"internalTrafficPolicy"`
		IPFamilies            []string `yaml:"ipFamilies"`
		IPFamilyPolicy        string   `yaml:"ipFamilyPolicy"`
		Ports                 []struct {
			Name       string `yaml:"name"`
			Port       int    `yaml:"port"`
			Protocol   string `yaml:"protocol"`
			TargetPort int    `yaml:"targetPort"`
		} `yaml:"ports"`
		ExternalName    string `yaml:"externalName"`
		SessionAffinity string `yaml:"sessionAffinity"`
		Type            string `yaml:"type"`
	} `yaml:"spec"`
	Status struct {
		LoadBalancer struct {
		} `yaml:"loadBalancer"`
	} `yaml:"status"`
}

type SERVICES struct {
	APIVersion string    `yaml:"apiVersion"`
	Items      []SERVICE `yaml:"items"`
	Kind       string    `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
	} `yaml:"metadata"`
}

type SECRET struct {
	APIVersion string            `yaml:"apiVersion"`
	Data       map[string]string `yaml:"data"`
	Kind       string            `yaml:"kind"`
	Metadata   struct {
		Annotations       map[string]string `yaml:"annotations"`
		CreationTimestamp time.Time         `yaml:"creationTimestamp"`
		Name              string            `yaml:"name"`
		Namespace         string            `yaml:"namespace"`
		OwnerReferences   []struct {
			APIVersion         string `yaml:"apiVersion"`
			BlockOwnerDeletion bool   `yaml:"blockOwnerDeletion"`
			Controller         bool   `yaml:"controller"`
			Kind               string `yaml:"kind"`
			Name               string `yaml:"name"`
			UID                string `yaml:"uid"`
		} `yaml:"ownerReferences"`
		ResourceVersion string `yaml:"resourceVersion"`
		UID             string `yaml:"uid"`
	} `yaml:"metadata"`
	Type string `yaml:"type"`
}

type SECRETS struct {
	APIVersion string   `yaml:"apiVersion"`
	Items      []SECRET `yaml:"items"`
	Kind       string   `yaml:"kind"`
	Metadata   struct {
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
	} `yaml:"metadata"`
}

type SUBSCRIPTION struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		CreationTimestamp time.Time `yaml:"creationTimestamp"`
		Generation        int       `yaml:"generation"`
		Labels            struct {
			OperatorsCoreosComClusterLoggingOpenshiftLogging string `yaml:"operators.coreos.com/cluster-logging.openshift-logging"`
		} `yaml:"labels"`
		Name            string `yaml:"name"`
		Namespace       string `yaml:"namespace"`
		ResourceVersion string `yaml:"resourceVersion"`
		SelfLink        string `yaml:"selfLink"`
		UID             string `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Channel             string `yaml:"channel"`
		InstallPlanApproval string `yaml:"installPlanApproval"`
		Name                string `yaml:"name"`
		Source              string `yaml:"source"`
		SourceNamespace     string `yaml:"sourceNamespace"`
		StartingCSV         string `yaml:"startingCSV"`
	} `yaml:"spec"`
	Status struct {
		CatalogHealth []struct {
			CatalogSourceRef struct {
				APIVersion      string `yaml:"apiVersion"`
				Kind            string `yaml:"kind"`
				Name            string `yaml:"name"`
				Namespace       string `yaml:"namespace"`
				ResourceVersion string `yaml:"resourceVersion"`
				UID             string `yaml:"uid"`
			} `yaml:"catalogSourceRef"`
			Healthy     bool      `yaml:"healthy"`
			LastUpdated time.Time `yaml:"lastUpdated"`
		} `yaml:"catalogHealth"`
		Conditions []struct {
			LastTransitionTime time.Time `yaml:"lastTransitionTime"`
			Message            string    `yaml:"message"`
			Reason             string    `yaml:"reason"`
			Status             string    `yaml:"status"`
			Type               string    `yaml:"type"`
		} `yaml:"conditions"`
		CurrentCSV            string `yaml:"currentCSV"`
		InstallPlanGeneration int    `yaml:"installPlanGeneration"`
		InstallPlanRef        struct {
			APIVersion      string `yaml:"apiVersion"`
			Kind            string `yaml:"kind"`
			Name            string `yaml:"name"`
			Namespace       string `yaml:"namespace"`
			ResourceVersion string `yaml:"resourceVersion"`
			UID             string `yaml:"uid"`
		} `yaml:"installPlanRef"`
		InstalledCSV string `yaml:"installedCSV"`
		Installplan  struct {
			APIVersion string `yaml:"apiVersion"`
			Kind       string `yaml:"kind"`
			Name       string `yaml:"name"`
			UUID       string `yaml:"uuid"`
		} `yaml:"installplan"`
		LastUpdated time.Time `yaml:"lastUpdated"`
		State       string    `yaml:"state"`
	} `yaml:"status"`
}

//////////////////

type ETCD_EP_H []struct {
	Endpoint string `json:"endpoint"`
	Health   bool   `json:"health"`
	Took     string `json:"took"`
}

type ETCD_EP_S []struct {
	Endpoint string `json:"Endpoint"`
	Status   struct {
		Header struct {
			ClusterID json.Number `json:"cluster_id"`
			MemberID  json.Number `json:"member_id"`
			Revision  int         `json:"revision"`
			RaftTerm  int         `json:"raft_term"`
		} `json:"header"`
		Version          string      `json:"version"`
		DbSize           int         `json:"dbSize"`
		Leader           json.Number `json:"leader"`
		RaftIndex        int         `json:"raftIndex"`
		RaftTerm         int         `json:"raftTerm"`
		RaftAppliedIndex int         `json:"raftAppliedIndex"`
		DbSizeInUse      int         `json:"dbSizeInUse"`
	} `json:"Status"`
}

type ETCD_M_L struct {
	Header struct {
		ClusterID json.Number `json:"cluster_id"`
		MemberID  json.Number `json:"member_id"`
		RaftTerm  int         `json:"raft_term"`
	} `json:"header"`
	Members []struct {
		ID         json.Number `json:"ID"`
		Name       string      `json:"name"`
		PeerURLs   []string    `json:"peerURLs"`
		ClientURLs []string    `json:"clientURLs"`
	} `json:"members"`
}
