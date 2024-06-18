package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      string
	AppVer   string
	Env      string
	Http     HttpConfig
	Log      LogConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Toggle   ToggleConfig
}

type HttpConfig struct {
	Port         int
	WriteTimeout int
	ReadTimeout  int
}

type LogConfig struct {
	FileLocation    string
	FileTDRLocation string
	FileMaxSize     int
	FileMaxBackup   int
	FileMaxAge      int
	Stdout          bool
}

type DatabaseConfig struct {
	Host            string
	User            string
	Password        string
	DBName          string
	Port            string
	SSLMode         string
	MaxIdleConn     int
	ConnMaxLifetime int
	MaxOpenConn     int
}

type RedisConfig struct {
	Mode     string
	Address  string
	Port     int
	Password string
}

type ToggleConfig struct {
	AppName string
	URL     string
	Token   string
}

func (c *Config) LoadConfig(path string) {
	viper.AddConfigPath(".")
	viper.SetConfigName(path)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	err = viper.Unmarshal(c)
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
