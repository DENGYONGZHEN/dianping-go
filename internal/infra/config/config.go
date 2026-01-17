package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// docker build 构建镜像
// docker compose 启动容器并注入 env
// linux 进程拥有环境变量
// go 读取环境变量
// viper 管理环境变量
// unmarshal 映射到配置结构体

type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	RedisAddress      string `mapstructure:"REDIS_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// 读取环境变量
	viper.AutomaticEnv()

	// 显式绑定
	keys := []string{
		"ENVIRONMENT",
		"DB_SOURCE",
		"HTTP_SERVER_ADDRESS",
		"REDIS_ADDRESS",
	}

	for _, key := range keys {
		_ = viper.BindEnv(key)
	}

	// 默认值
	viper.SetDefault("HTTP_SERVER_ADDRESS", ":8080")
	viper.SetDefault("ENVIRONMENT", "dev")

	// 如果有文件就读，没有就忽略
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("no config file found, using env variables only")
	}

	err = viper.Unmarshal(&config)
	return
}
