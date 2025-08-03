package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"usercenter/model"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

var Cfg Config
var DB *gorm.DB

func InitConfig() {
	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
}

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		Cfg.Database.Host,
		Cfg.Database.Port,
		Cfg.Database.Username,
		Cfg.Database.Password,
		Cfg.Database.DBName,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("数据库自动迁移失败: %v", err)
	}
}
