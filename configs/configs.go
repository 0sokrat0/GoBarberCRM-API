package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"`
	Port        int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SslMode  string `mapstructure:"sslmode"`
}

var AppConfigInstance *Config

// LoadConfig загружает конфигурацию из файла и переменных окружения
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config") // Имя файла конфигурации (без расширения)
	viper.SetConfigType("yaml")   // Тип файла
	viper.AddConfigPath(path)     // Путь к файлу конфигурации
	viper.AutomaticEnv()          // Поддержка переменных окружения

	// Читаем файл конфигурации
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Маппим содержимое файла в структуру Config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	log.Println("Configuration loaded successfully.")
	AppConfigInstance = &config
	return &config, nil
}
