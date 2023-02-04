package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server *serverConfig
	DB     *databaseConfig
}

type serverConfig struct {
	Port            int    `env:"PORT,default=8080"`
	AllowCorsOrigin string `env:"ALLOW_CORS_ORIGIN,default=*"`
}

type databaseConfig struct {
	DSN              string `env:"MYSQL_DSN,default=root:password@tcp(localhost:33061)/example?charset=utf8&parseTime=true"`
	MaxOpenConns     int    `env:"MAX_OPEN_CONNS,default=100"`
	MaxIdleConns     int    `env:"MAX_IDLE_CONNS,default=100"`
	ConnsMaxLifetime int    `env:"CONNS_MAX_LIFETIME,default=100"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg *Config) Address() string {
	return fmt.Sprintf(":%d", cfg.Server.Port)
}
