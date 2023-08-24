package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

type Config interface {
	GetBool(key string) bool
	GetInt(key string) int
	GetString(key string) string
	GetStringSlice(key string) []string
	GetUInt64(key string) uint64
	GetStringMap(key string) map[string]interface{}
	Init()
}

type viperConfig struct{}

func (v *viperConfig) Init() {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error Reading Config file: %s", err)
	}
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetUInt64 ...
func (v *viperConfig) GetUInt64(key string) uint64 {
	return viper.GetUint64(key)
}

func (v *viperConfig) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()

	return v
}
