package config

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := GetConfig(struct {
		envPrefix  string
		configFile string
	}{
		envPrefix:  "MYAPP",
		configFile: "../../config.yml",
	})
	if cfg == nil {
		t.Fatal("Expected config to be loaded")
	}

	expected := &Config{
		Cloud: CloudConfig{
			Environment: "local",
		},
		Server: ServerConfig{
			Host: "127.0.0.1",
			Port: "1337",
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "123456",
			DBName:   "go_tracker",
			SSLMode:  "disable",
		},
		Log: LogConfig{
			Level: "debug",
		},
	}

	if reflect.DeepEqual(cfg, expected) {
		t.Log("Config loaded successfully")
	} else {
		t.Errorf("Expected %+v, got %+v", expected, cfg)
	}
}
