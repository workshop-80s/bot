package lib

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	root = "./config"
)

func GetEnv() string {
	env := os.Getenv("ENV")

	if env != "" {
		return env
	}

	return "dev"
}

func GetConfigString(name string, key string) string {
	configLocation := "./config"
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configLocation)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return viper.GetString(key)
}

func GetAllConfigMap(name string) map[string]string {
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(root)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	config := map[string]string{}
	for k, v := range viper.AllSettings() {
		config[k] = v.(string)
	}

	return config
}

func GetConfigMap(name string, key string) map[string]string {
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(root)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return viper.GetStringMapString(key)
}

func GetEnvConfigMap(name string, params ...string) map[string]string {
	key := ""

	switch len(params) {
	case 0:
		key = GetEnv()
	case 1:
		key = params[0]
	default:
		key = params[1] + "." + params[0]
	}

	return GetConfigMap(name, key)
}
