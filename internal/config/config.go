package config

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

var AppConfig *Config

type Config struct {
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
}

type Database struct {
	DBFile       string `yaml:"db_file"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func LoadConfig(configPath string) error {
	if configPath == "" {
		env := os.Getenv("APP_ENV")
		if env != "" {
			configPath = fmt.Sprintf("config.%s.yaml", env)
		} else {
			configPath = "config.yaml"
		}
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", configPath, err)
	}

	// 解析YAML配置
	AppConfig = &Config{}
	if err := yaml.Unmarshal(data, AppConfig); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	log.Printf("Config loaded from %s", configPath)
	return nil
}
