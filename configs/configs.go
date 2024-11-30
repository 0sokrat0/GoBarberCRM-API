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

	BindEnvVariables() // Привязываем переменные окружения

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

// BindEnvVariables привязывает ключи конфигурации к переменным окружения
func BindEnvVariables() {
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.sslmode", "DB_SSLMODE")
	viper.BindEnv("app.port", "APP_PORT")
	viper.BindEnv("app.environment", "APP_ENVIRONMENT")
}

// GetDSN формирует строку подключения к базе данных
func GetDSN(config DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
		config.SslMode,
	)
}

// TestConfig проверяет загрузку конфигурации
func TestConfig() {
	config, err := LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("App Name: %s, Environment: %s, Port: %d",
		config.App.Name,
		config.App.Environment,
		config.App.Port,
	)

	log.Printf("Database Host: %s, User: %s, DB Name: %s",
		config.Database.Host,
		config.Database.User,
		config.Database.Name,
	)
}
