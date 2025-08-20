package config

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/spf13/viper"
)

var GlobalConfig ILCConfig

func bindEnvVars(bindings map[string]string) (err error) {
	for key, envVar := range bindings {
		err = viper.BindEnv(key, envVar)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateEnvBindings() map[string]string {
	bindings := make(map[string]string)

	// Top-level fields
	bindings["project_dir"] = "ILC__PROJECT_DIR"
	bindings["apps.config_path"] = "ILC__CONFIG_PATH"

	// App binaries
	apps := []string{
		"docker",
		"helm",
		"kind",
		"minikube",
		"podman",
		"skopeo",
		"webhook",
	}
	for _, app := range apps {
		key := fmt.Sprintf("apps.%s.binary", app)
		envVar := fmt.Sprintf("ILC__%s__BINARY", strings.ToUpper(app))
		bindings[key] = envVar
	}

	// docker
	bindings["apps.docker.config_path"] = "ILC__DOCKER__CONFIG_PATH"

	// webhook
	bindings["apps.webhook.webhooks_path"] = "ILC__WEBHOOK__WEBHOOKS_PATH"
	bindings["apps.webhook.listen_addr"] = "ILC__WEBHOOK__LISTEN_ADDR"
	bindings["apps.webhook.listen_port"] = "ILC__WEBHOOK__LISTEN_PORT"
	bindings["apps.webhook.secret"] = "ILC__WEBHOOK__SECRET"
	bindings["apps.webhook.prefix"] = "ILC__WEBHOOK__PREFIX"

	// minikube
	bindings["apps.minikube.cluster_name"] = "ILC__MINIKUBE__CLUSTER_NAME"
	bindings["apps.minikube.kube_version"] = "ILC__MINIKUBE__KUBE_VERSION"
	bindings["apps.minikube.nodes_count"] = "ILC__MINIKUBE__NODES_COUNT"
	bindings["apps.minikube.driver"] = "ILC__MINIKUBE__DRIVER"
	bindings["apps.minikube.cni"] = "ILC__MINIKUBE__CNI"
	bindings["apps.minikube.cidr"] = "ILC__MINIKUBE__CIDR"
	bindings["apps.minikube.cpus"] = "ILC__MINIKUBE__CPUS"
	bindings["apps.minikube.memory"] = "ILC__MINIKUBE__MEMORY"
	bindings["apps.minikube.disk_size"] = "ILC__MINIKUBE__DISK_SIZE"

	// kind
	bindings["apps.kind.cluster_name"] = "ILC__KIND__CLUSTER_NAME"
	bindings["apps.kind.config_path"] = "ILC__KIND__CONFIG_PATH"

	// podman
	bindings["apps.podman.machine_name"] = "ILC__PODMAN__MACHINE_NAME"
	bindings["apps.podman.cpus"] = "ILC__PODMAN__CPUS"
	bindings["apps.podman.memory"] = "ILC__PODMAN__MEMORY"
	bindings["apps.podman.disk_size"] = "ILC__PODMAN__DISK_SIZE"

	return bindings
}

// TODO: add test to be sure all the configs has default value or not if this is expected
// TODO: Consider to make new configs values and bind to vars easier if possible. Right now 3 places to modify and easy to forget or miss

func setDefaultValues() {
	viper.SetDefault("version", "1")

	apps := []string{
		"docker",
		"helm",
		"kind",
		"minikube",
		"podman",
		"skopeo",
		"webhook",
	}
	for _, app := range apps {
		viper.SetDefault(fmt.Sprintf("apps.%s.binary", app), app)
	}

	usr, _ := user.Current()
	viper.SetDefault("apps.docker.config_path", usr.HomeDir+"/.docker")

	viper.SetDefault("apps.webhook.webhooks_path", "./webhooks.yaml")
	viper.SetDefault("apps.webhook.listen_addr", "127.0.0.1")
	viper.SetDefault("apps.webhook.listen_port", 8080)
	viper.SetDefault("apps.webhook.secret", "")
	viper.SetDefault("apps.webhook.prefix", "hooks")

	viper.SetDefault("apps.kind.cluster_name", "kind-local")
	viper.SetDefault("apps.kind.config_path", "")

	viper.SetDefault("apps.minikube.cluster_name", "mini-local")
	viper.SetDefault("apps.minikube.kube_version", "v1.30.0")
	viper.SetDefault("apps.minikube.nodes_count", 1)
	viper.SetDefault("apps.minikube.driver", "podman")
	viper.SetDefault("apps.minikube.cni", "auto")
	viper.SetDefault("apps.minikube.cidr", "172.16.0.0/16")
	viper.SetDefault("apps.minikube.cpus", "2")
	viper.SetDefault("apps.minikube.memory", "2G")
	viper.SetDefault("apps.minikube.disk_size", "15G")

	viper.SetDefault("apps.podman.machine_name", "podman-machine-default")
	viper.SetDefault("apps.podman.cpus", "2")
	viper.SetDefault("apps.podman.memory", "2G")
	viper.SetDefault("apps.podman.disk_size", "10G")
}

func GetConfig() *ILCConfig {
	return &GlobalConfig
}

func LoadConfig() (err error) {
	viper.SetEnvPrefix("ILC_")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	setDefaultValues()
	envBindings := generateEnvBindings()

	err = bindEnvVars(envBindings)
	if err != nil {
		return err
	}
	usr, _ := user.Current()
	configPath := fmt.Sprintf("%s/.infra-lab.yaml", usr.HomeDir)

	viper.SetConfigFile(configPath)

	// TODO add error handling. It is possible to use `viper.SafeWriteConfigAs(configPath)` to try to write to file, but not to override it
	_ = viper.ReadInConfig()

	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	_ = LoadConfig()
}
