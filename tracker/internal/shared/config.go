package shared

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load(envFile string) {
	if envFile == "" {
		envFile = ".env"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}
	c.DBUser = os.Getenv("DB_USER")
	c.DBPassword = os.Getenv("DB_PASSWORD")
	c.DBName = os.Getenv("DB_NAME")
	c.DBHost = os.Getenv("DB_HOST")
	c.DBPort = os.Getenv("DB_PORT")
}
