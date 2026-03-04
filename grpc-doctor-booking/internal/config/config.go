package config

import "github.com/spf13/viper"

type Config struct {
	GRPCPort    string
	GatewayPort string
}

func Load() (Config, error) {
	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("GATEWAY_PORT", "8080")
	viper.AutomaticEnv()

	return Config{
		GRPCPort:    viper.GetString("GRPC_PORT"),
		GatewayPort: viper.GetString("GATEWAY_PORT"),
	}, nil
}
