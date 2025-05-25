package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"reflect"
	"strings"
)

const (
	EnvPrefix = "IIMS"
)

type Config struct {
	Database DatabaseConfig
	Server   struct {
		Host           string `yaml:"host" mapstructure:"host"`
		GrpcPort       int    `yaml:"grpc_port" mapstructure:"grpc_port"`
		RequestTimeout int    `yaml:"request_timeout" mapstructure:"request_timeout"`
		InsertDuration int    `yaml:"insert_duration" mapstructure:"insert_duration"`
		PathToData     string `yaml:"path_to_data" mapstructure:"path_to_data"`
	} `yaml:"server" mapstructure:"server"`
}

type DatabaseConfig struct {
	HealthcheckTimeout int    `yaml:"healthcheck_timeout" mapstructure:"healthcheck_timeout"`
	Uri                string `yaml:"uri" mapstructure:"uri"`
	Database           string `yaml:"database" mapstructure:"database"`
	MigrationsPath     string `yaml:"migrations_path" mapstructure:"migrations_path"`
	*options.ClientOptions
}

func Get(logger zerolog.Logger) *Config {
	v := viper.New()
	v.SetEnvPrefix(EnvPrefix)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AddConfigPath("./config/")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to read config")
	}

	for _, key := range v.AllKeys() {
		val := v.Get(key)
		if val == nil {
			continue
		}

		if reflect.TypeOf(val).Kind() == reflect.String {
			v.Set(key, os.ExpandEnv(val.(string)))
		}
	}

	var cfg *Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to unmarshal config")
	}

	return cfg
}
