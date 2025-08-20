// package config
package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
	Cloud    CloudConfig    `mapstructure:"cloud"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"` // "disable" local, "require" AWS (RDS)
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"expiration"` // seconds
}

type LogConfig struct {
	Level  string `mapstructure:"level"`  // debug,info,warn,error
	Format string `mapstructure:"format"` // text,json
}

type CloudConfig struct {
	Environment string `mapstructure:"environment"` // local, aws-lambda, aws-ecs
	Region      string `mapstructure:"region"`
}

var (
	once     sync.Once
	instance *Config
)

func GetConfig(config struct {
	envPrefix  string
	configFile string
}) *Config {
	if config.configFile == "" {
		config.configFile = "./configs/app.yaml"
	}
	once.Do(func() {
		instance = load(config)
	})
	return instance
}

func load(config struct {
	envPrefix  string
	configFile string
}) *Config {
	env := detectEnvironment(config.envPrefix)
	log.Printf("Detected environment: %s", env)

	// Base Viper setup: env > file (only read file locally)
	viper.SetEnvPrefix(config.envPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if env == "local" {
		setupLocalFileConfig(config.configFile)
	}

	setDefaults(env)
	tryReadConfigFile() // best-effort; fine if missing

	// Inject cloud awareness
	if region := os.Getenv("AWS_REGION"); region != "" {
		viper.Set("cloud.region", region)
	}
	viper.Set("cloud.environment", env)

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("config unmarshal error: %v", err)
	}
	return &cfg
}

func detectEnvironment(envPrefix string) string {
	// Local
	env := os.Getenv(envPrefix + "_ENV")

	if env != "" {
		return env
	}
	return "local"
}

func setupLocalFileConfig(configFile string) {
	viper.SetConfigFile(configFile)
}

func tryReadConfigFile() {
	if err := viper.ReadInConfig(); err != nil {
		// Totally fine in AWS where we expect env-only config
		log.Printf("No config file loaded (ok in AWS): %v", err)
	}
}

func setDefaults(env string) {
	// Server
	if env == "local" {
		viper.SetDefault("server.host", "localhost")
		viper.SetDefault("server.port", "8080")
	} else {
		viper.SetDefault("server.host", "0.0.0.0")
		viper.SetDefault("server.port", "8080")
	}

	// Database
	if env == "local" {
		viper.SetDefault("database.host", "localhost")
		viper.SetDefault("database.port", "5432")
		viper.SetDefault("database.user", "postgres")
		viper.SetDefault("database.password", "postgres")
		viper.SetDefault("database.dbname", "appdb")
		viper.SetDefault("database.sslmode", "disable")
	} else {
		// In AWS, prefer env overrides (e.g., Secrets Manager/SSM -> env)
		// Keep safe defaults:
		viper.SetDefault("database.sslmode", "require")
	}

	// Logs: JSON in AWS, pretty text locally
	if env == "local" {
		viper.SetDefault("log.level", "debug")
	} else {
		viper.SetDefault("log.level", "info")
	}
}

// Convenience helpers
func (c *Config) IsProduction() bool { return c.Cloud.Environment != "local" }
func (c *Config) IsAWS() bool {
	return c.Cloud.Environment == "aws"
}
