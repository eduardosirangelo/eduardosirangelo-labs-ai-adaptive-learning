package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port             string        `mapstructure:"port"`
	DatabaseURL      string        `mapstructure:"database_url"`
	KafkaBrokers     []string      `mapstructure:"kafka_brokers"`
	ReadTimeout      time.Duration `mapstructure:"read_timeout"`
	RequestTimeout   time.Duration `mapstructure:"request_timeout"`
	CORSOrigins      []string      `mapstructure:"cors_origins"`
	RateLimit        int           `mapstructure:"rate_limit"`
	RateWindow       time.Duration `mapstructure:"rate_window"`
	CompressionLevel int           `mapstructure:"compression_level"`
}

func Load(serviceName string) (*Config, error) {
	v := viper.New()

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// monta paths para services/<serviceName> e services/<serviceName>/config
	svcDir := filepath.Join(wd, "services", serviceName)
	v.AddConfigPath(svcDir)
	v.AddConfigPath(filepath.Join(svcDir, "config"))
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	// Read config file (ignore if not found)
	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// 2) Environment variables override
	v.AutomaticEnv()
	// Map ENV_VAR to nested keys: replace "." with "_"
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// Bind expected env vars
	for _, key := range []string{
		"PORT",
		"DATABASE_URL",
		"KAFKA_BROKERS",
		"READ_TIMEOUT",
		"REQUEST_TIMEOUT",
		"CORS_ORIGINS",
		"RATE_LIMIT",
		"RATE_WINDOW",
		"COMPRESSION_LEVEL",
	} {
		_ = v.BindEnv(strings.ToLower(key))
	}

	// 3) Defaults
	v.SetDefault("port", "8070")
	v.SetDefault("read_timeout", 5*time.Second)
	v.SetDefault("request_timeout", 10*time.Second)
	v.SetDefault("cors_origins", []string{"*"})
	v.SetDefault("rate_limit", 1000)
	v.SetDefault("rate_window", 1*time.Minute)
	v.SetDefault("compression_level", 5)

	// 4) Unmarshal into struct
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &cfg, nil
}
