package config

type ILCConfig struct {
	Version     string `mapstructure:"version" yaml:"version"`
	Apps        apps   `mapstructure:"apps" yaml:"apps"`
	ProjectsDir string `mapstructure:"projects_dir" yaml:"projects_dir"`
}

type apps struct {
	Docker   dockerConfig `mapstructure:"docker" yaml:"docker"`
	Helm     appConfig    `mapstructure:"helm" yaml:"helm"`
	Kind     Kind         `mapstructure:"kind" yaml:"kind"`
	Minikube appConfig    `mapstructure:"minikube" yaml:"minikube"`
	Podman   appConfig    `mapstructure:"podman" yaml:"podman"`
	Skopeo   appConfig    `mapstructure:"skopeo" yaml:"skopeo"`
	Webhook  Webhook      `mapstructure:"webhook" yaml:"webhook"`
}

type appConfig struct {
	Binary string `mapstructure:"binary" yaml:"binary"`
}

type dockerConfig struct {
	Binary     string `mapstructure:"binary" yaml:"binary"`
	ConfigPath string `mapstructure:"config_path" yaml:"config_path"`
}

type Webhook struct {
	Binary       string `mapstructure:"binary" yaml:"binary"`
	Secret       string `mapstructure:"secret" yaml:"secret"`
	ListenAddr   string `mapstructure:"listen_addr" yaml:"listen_addr"`
	ListenPort   int    `mapstructure:"listen_port" yaml:"listen_port"`
	WebhooksPath string `mapstructure:"webhooks_path" yaml:"webhooks_path"`
	Prefix       string `mapstructure:"prefix" yaml:"prefix"`
}

type Kind struct {
	Binary      string `mapstructure:"binary" yaml:"binary"`
	ClusterName string `mapstructure:"cluster_name" yaml:"cluster_name"`
	ConfigPath  string `mapstructure:"config_path" yaml:"config_path"`
}
