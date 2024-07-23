package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig(conf interface{}, configPaths ...string) {
	v := viper.New()

	// 默认配置文件名称和类型
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// 添加配置文件路径
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	// 读取环境变量
	v.AutomaticEnv()

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 解析配置到结构体
	if err := v.Unmarshal(&conf); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
