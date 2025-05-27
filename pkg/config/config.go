package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port             string        `mapstructure:"PORT"`
	DatabaseURL      string        `mapstructure:"DATABASE_URL"`
	KafkaBrokers     []string      `mapstructure:"KAFKA_BROKERS"`
	ReadTimeout      time.Duration `mapstructure:"READ_TIMEOUT"`
	RequestTimeout   time.Duration `mapstructure:"REQUEST_TIMEOUT"`
	CORSOrigins      []string      `mapstructure:"CORS_ORIGINS"`      // e.g., ["https://example.com", "https://another.com"]
	RateLimit        int           `mapstructure:"RATE_LIMIT"`        // max requests per IP
	RateWindow       time.Duration `mapstructure:"RATE_WINDOW"`       // e.g., 1 minute
	CompressionLevel int           `mapstructure:"COMPRESSION_LEVEL"` // gzip compression level (1-9)
}

func Load() (*Config, error) {
	v := viper.New()

	// 1) Configuration via YAML file//
	// Search in ./config.yaml and ./config.{ENV}.yaml
	v.SetConfigName("config") // base: config.yaml
	v.SetConfigType("yaml")
	v.AddConfigPath(".")                 // ./services/content-engine
	v.AddConfigPath("./internal/config") // now includes the folder where config.yaml is located

	if err := v.ReadInConfig(); err != nil {
		var notFoundErr viper.ConfigFileNotFoundError
		if errors.As(err, &notFoundErr) {
			// if no file is found, continue without error
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// 2) Environment variables
	// reads any VAR in UPPERCASE format
	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
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
		_ = v.BindEnv(key)
	}

	// 3) Default values (if not from the file or env)
	// 3) Defaults para HTTP
	v.SetDefault("PORT", "8070")
	v.SetDefault("READ_TIMEOUT", 5*time.Second)
	v.SetDefault("REQUEST_TIMEOUT", 10*time.Second)

	// HTTP-specific defaults:
	v.SetDefault("CORS_ORIGINS", []string{"*"})       // permite todos por padrão
	v.SetDefault("RATE_LIMIT", 1000)                  // 1000 req/IP
	v.SetDefault("RATE_WINDOW", 1*time.Minute)        // janela de 1 minuto
	v.SetDefault("COMPRESSION_LEVEL", 5)              // nível médio de gzip

	// 4) Unmarshal
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &cfg, nil
}
