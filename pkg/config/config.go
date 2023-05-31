package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	DB *databaseConfig
}

type databaseConfig struct {
	AwsRegion       string `env:"AWS_REGION"`
	DynamoEndpoint  string `env:"DYNAMO_ENDPOINT"`
	DynamoTableName string `env:"TABLE_NAME"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
