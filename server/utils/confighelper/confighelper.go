package confighelper

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

//InitViper function to initialize viper
func InitViper(configName string) {
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()
	// Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file %s ", err))
	}
}

//GetConfig method to get configs from config file
func GetConfig(keyName string) string {
	keyValue := viper.GetString(keyName)
	return keyValue
}

//GetConfigAsDuration method to get configs from config file
func GetConfigAsDuration(keyName string) time.Duration {
	keyValue := viper.GetDuration(keyName)
	return keyValue
}

//GetConfigAsInt64 method to get configs from config file
func GetConfigAsInt64(keyName string) int64 {
	keyValue := viper.GetInt64(keyName)
	return keyValue
}

//GetConfigAsInt method to get configs from config file
func GetConfigAsInt(keyName string) int {
	keyValue := viper.GetInt(keyName)
	return keyValue
}

//GetConfigAsBool method to get configs from config file
func GetConfigAsBool(keyName string) bool {
	keyValue := viper.GetBool(keyName)
	return keyValue
}
