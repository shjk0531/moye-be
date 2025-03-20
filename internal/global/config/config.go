package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort string `mapstructure:"server_port"`

	// PostgreSQL 설정
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`

	// MongoDB 설정
	MongoURI string `mapstructure:"mongo_uri"`
	MongoDB  string `mapstructure:"mongo_db"`
}

var Config AppConfig

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("설정 파일 읽기 실패: %v", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("설정값 매핑 실패: %v", err)
	}
}
