package config

type AppConfig struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
}
