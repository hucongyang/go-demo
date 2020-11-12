package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	RunMode string `json:"run_mode"`
	App struct{
		PageSize int `json:"page_size"`
		JwtSecret string `json:"jwt_secret"`
	}
	Server struct{
		HttpPort int `json:"http_port"`
		ReadTimeout int `json:"read_timeout"`
		WriteTimeout int `json:"write_timeout"`
	}
	Database struct{
		Driver string `json:"driver"`
		User string `json:"user"`
		Password string `json:"password"`
		Host string `json:"host"`
		Port int `json:"port"`
		Name string `json:"name"`
		TablePrefix string `json:"table_prefix"`
	}
	Redis struct{
		Addr string `json:"addr"`
		Password string `json:"password"`
		Db int `json:"db"`
	}
}

var conf *Configuration

func Config() *Configuration {
	if conf != nil {
		return conf
	}
	var err error

	viper.SetConfigName("configuration")
	viper.AddConfigPath("../go-demo")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	if err = viper.Unmarshal(&conf); err != nil {
		fmt.Println("config file error:", err)
		os.Exit(1)
	}
	fmt.Println("Configuration.conf", conf)
	return conf
}