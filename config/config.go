package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database DatabaseConfig
	Server   ServersConfig
	Logger   LoggerConfig
}
type ServersConfig struct {
	HTTP HTTPConfig `mapstructure:"http"`
}

type HTTPConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	SSLMode         string `mapstructure:"sslmode"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerConfig struct {
	Mode     string `mapstructure:"mode"`
	Level    string `mapstructure:"level"`
	FilePath string `mapstructure:"filepath"`
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] No .env file found. Using system environment variables.")
	}

	// Mapping
	_ = viper.BindEnv("server.http.host", "TODO_HTTP_HOST")
	_ = viper.BindEnv("server.http.port", "TODO_HTTP_PORT")
	_ = viper.BindEnv("database.host", "TODO_DATABASE_HOST")
	_ = viper.BindEnv("database.port", "TODO_DATABASE_PORT")
	_ = viper.BindEnv("database.user", "TODO_DATABASE_USER")
	_ = viper.BindEnv("database.password", "TODO_DATABASE_PASSWORD")
	_ = viper.BindEnv("database.dbname", "TODO_DATABASE_DBNAME")
	_ = viper.BindEnv("database.sslmode", "TODO_DATABASE_SSLMODE")
	_ = viper.BindEnv("database.maxidleconns", "TODO_DATABASE_MAXIDLECONNS")
	_ = viper.BindEnv("database.maxopenconns", "TODO_DATABASE_MAXOPENCONNS")
	_ = viper.BindEnv("database.connmaxlifetime", "TODO_DATABASE_CONNMAXLIFETIME")

	// Cfg file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	// Read cfg file
	if err := viper.MergeInConfig(); err != nil {

		log.Println("[DEBUG] Config file not found, fallback to .env")
	}

	// Decode
	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("[ERROR] Failed to load config: %v", err)
	}

	return cfg
}
