package bootstrap

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func OnInitialize(configFile *string, envPrefix string, loadDirs []string, defaultConfigName string) func() {
	return func() {
		fmt.Println("configFile=", *configFile)
		if configFile != nil && len(*configFile) > 0 {
			viper.SetConfigFile(*configFile)
		} else {
			for _, dir := range loadDirs {
				viper.AddConfigPath(dir)
			}

			viper.SetConfigType("yaml")
			viper.SetConfigName(defaultConfigName)
		}

		viper.AutomaticEnv()
		viper.SetEnvPrefix(envPrefix)
		replacer := strings.NewReplacer(".", "_", "-", "_")
		viper.SetEnvKeyReplacer(replacer)

		_ = viper.ReadInConfig()
	}
}
