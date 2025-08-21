package config

import (
	"fmt"
	"os/user"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

var GlobalConfig ILCConfig

func GetConfig() *ILCConfig {
	return &GlobalConfig
}

func setDefaultsAndBindEnvs(cfg any, path, envPrefix string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic in setDefaultsAndBindEnvs: %v\n", r)
		}
	}()

	cfgType := reflect.TypeOf(cfg)
	cfgValue := reflect.ValueOf(cfg)

	for i := 0; i < cfgType.NumField(); i++ {
		field := cfgType.Field(i)
		tag := field.Tag.Get("mapstructure")

		if field.Type.Kind() == reflect.Struct {
			var newPath string
			if path == "" {
				newPath = tag
			} else {
				newPath = path + "." + tag
			}
			setDefaultsAndBindEnvs(cfgValue.Field(i).Interface(), newPath, envPrefix)
		} else {
			var cfgPath string
			defaultValue := field.Tag.Get("default")
			if path == "" {
				cfgPath = tag
			} else {
				cfgPath = path + "." + tag
			}
			envPath := strings.ToUpper(strings.ReplaceAll(cfgPath, ".", "__"))
			err := viper.BindEnv(cfgPath, envPrefix+"_"+envPath)
			if err != nil {
				fmt.Printf("BindEnv error: %v\n", err)
			}
			viper.SetDefault(cfgPath, defaultValue)

		}
	}
}

// TODO: add test to be sure all the configs has default value or not if this is expected

func LoadConfig() (err error) {
	envPrefix := "ILC_"
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	setDefaultsAndBindEnvs(GlobalConfig, "", envPrefix)
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
