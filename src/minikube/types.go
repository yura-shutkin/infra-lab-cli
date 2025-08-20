package minikube

// TODO: ensure ExtraArgs are passed in commands

type Cluster struct {
	Name       string        `json:"Name"`
	Status     string        `json:"Status,omitempty"`
	Config     ClusterConfig `json:"Config,omitempty"`
	CNI        string        `json:"-"`
	CIDR       string        `json:"-"`
	NodesCount int           `json:"-"`
	ExtraArgs  string        `json:"-"`
}

type ClusterConfig struct {
	CPUs           int             `json:"cpus"`
	CPUsFlag       string          `json:"-"`
	Memory         int             `json:"memory"`
	MemoryFlag     string          `json:"-"`
	DiskSize       int             `json:"DiskSize"`
	DiskSizeFlag   string          `json:"-"`
	Driver         string          `json:"-"`
	KubeConfig     KubeConfig      `json:"KubernetesConfig"`
	Nodes          []Node          `json:"Nodes,omitempty"`
	RegistryMirror []string        `json:"RegistryMirror"`
	Addons         map[string]bool `json:"Addons,omitempty"`
}

type KubeConfig struct {
	KubeVersion  string         `json:"kubernetesVersion"`
	ExtraOptions []ExtraOptions `json:"ExtraOptions,omitempty"`
}

type Node struct {
	Name              string `json:"Name"`
	IP                string `json:"IP"`
	Port              int    `json:"Port"`
	KubernetesVersion string `json:"KubernetesVersion"`
	ContainerRuntime  string `json:"ContainerRuntime"`
	ControlPlane      bool   `json:"ControlPlane"`
	Worker            bool   `json:"Worker"`
}

type MkList struct {
	Invalid []Cluster `json:"-"`
	Valid   []Cluster `json:"valid"`
}

type ExtraOptions struct {
	Component string `json:"Component"`
	Key       string `json:"Key"`
	Value     string `json:"Value"`
}
