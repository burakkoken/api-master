package apimaster

import (
	"fmt"
	"os"
)

import "github.com/spf13/viper"

func init() {
	viper.SetConfigName(getEnvironmentFileName())
	viper.AddConfigPath(getEnvironmentFilePath())

	viper.AutomaticEnv()
	viper.SetConfigType(EnvironmentFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("error reading config file, %s", err))
	}
	fmt.Println(viper.Get("endpoints.httpbin"))
}

func getEnvironmentFileName() string {
	environment := os.Getenv(EnvironmentKey)

	if environment == "" {
		return EnvironmentFilePrefix
	}

	return EnvironmentFilePrefix + "." + environment
}

func getEnvironmentFilePath() string {
	environmentFilePath := os.Getenv(EnvironmentFilePathKey)

	if environmentFilePath == "" {
		return DefaultEnvironmentFilePath
	}

	return environmentFilePath
}
