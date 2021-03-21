package apimaster

import (
	"github.com/spf13/viper"
	"time"
)

const (
	EnvironmentKey         = "API_MASTER_ENV"
	EnvironmentFilePathKey = "API_MASTER_ENV_FILE_PATH"

	EnvironmentFilePrefix      = "env"
	EnvironmentFileType        = "yaml"
	DefaultEnvironmentFilePath = "."
)

type Environment struct {
}

func GetEnvironment() *Environment {
	return &Environment{}
}

func (environment *Environment) Get(key string) interface{} {
	return viper.Get(key)
}

func (environment *Environment) GetString(key string) string {
	return viper.GetString(key)
}

func (environment *Environment) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (environment *Environment) GetInt(key string) int {
	return viper.GetInt(key)
}

func (environment *Environment) GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func (environment *Environment) GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func (environment *Environment) GetUint(key string) uint {
	return viper.GetUint(key)
}

func (environment *Environment) GetUint32(key string) uint32 {
	return viper.GetUint32(key)
}

func (environment *Environment) GetUint64(key string) uint64 {
	return viper.GetUint64(key)
}

func (environment *Environment) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func (environment *Environment) GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func (environment *Environment) GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func (environment *Environment) GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func (environment *Environment) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (environment *Environment) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}
