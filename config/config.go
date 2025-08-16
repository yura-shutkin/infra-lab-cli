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

	return bindings
}

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
