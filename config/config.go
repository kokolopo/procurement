package config

import "os"

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	JWT_KEY        string
	VT_SERVER_KEY  string
	VT_CLIENT_KEY  string
}

func InitConfiguration() Config {
	// logic dapatin env
	// file(.env, env.yaml, env.json,...), system env

	return Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8080"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "root"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "password"),
		DB_NAME:        GetOrDefault("DB_NAME", "procurementDB"),
		DB_PORT:        GetOrDefault("DB_PORT", "3306"),
		DB_HOST:        GetOrDefault("DB_HOST", "35.202.50.135"),
	}
}

func GetOrDefault(envKey, defaultValue string) string {
	// cek env
	if val, exist := os.LookupEnv(envKey); exist {
		return val
	}
	// kalau ada, return valuenya
	// kalau gaada, return defaultValuenya
	return defaultValue
}
