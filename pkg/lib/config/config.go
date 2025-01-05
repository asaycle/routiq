package config

import (
	"log"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `config:"server"`
}

type ServerConfig struct {
	Port string `config:"port"`
}

func Load() (*Config, error) {
	// Set default values
	v := viper.New()
	v.SetEnvPrefix("MOTIFY")
	v.AutomaticEnv() // Automatically map environment variables

	setDefaults(v, defaultConfig, "")

	// Read the config file
	v.SetConfigName("config") // Name of config file (without extension)
	v.SetConfigType("yaml")   // Config file type
	v.AddConfigPath(".")      // Path to look for the config file in
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Config file not found, using environment variables: %v\n", err)
	}

	// Unmarshal into the AppConfig struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

var defaultConfig = &Config{
	Server: ServerConfig{
		Port: "50051",
	},
}

func setDefaults(v *viper.Viper, data interface{}, prefix string) {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i).Interface()
		fieldName := field.Tag.Get("config")

		// Build the Viper key with the prefix
		key := fieldName
		if prefix != "" {
			key = strings.Join([]string{prefix, fieldName}, ".")
		}

		// If the field is a struct, recursively set defaults
		if reflect.TypeOf(fieldValue).Kind() == reflect.Struct {
			setDefaults(v, fieldValue, key)
		} else {
			v.SetDefault(key, fieldValue)
		}
	}
}
