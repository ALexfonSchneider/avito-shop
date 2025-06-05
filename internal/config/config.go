package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"os"
)

type AppEnv string

const (
	Local AppEnv = "local"
	Test  AppEnv = "test"
)

type Config struct {
	Postgres PostgresConfig `koanf:"postgres"`
	Auth     AuthConfig     `koanf:"auth"`
	App      AppConfig      `koanf:"app"`
}

func loadConfig(path string) (*Config, error) {
	cfg := &Config{}

	k := koanf.New(".")

	if err := k.Load(
		file.Provider(path),
		yaml.Parser(),
	); err != nil {
		return nil, err
	}

	// overriding
	if err := k.Load(env.Provider("", ".", nil), nil); err != nil {
		fmt.Printf("error loading env vars: %v\n", err)
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, err
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	if err := v.Struct(cfg); err != nil {
		panic(err)
	}

	return cfg, nil
}

func getSelectedEnv() AppEnv {
	selectedEnv := Local

	if env := os.Getenv("APP_ENV"); env != "" {
		selectedEnv = AppEnv(env)
	}

	return selectedEnv
}

func getConfigFolder() string {
	selectedFolder := "config/"

	if path := os.Getenv("APP_CONFIG_PATH"); path != "" {
		selectedFolder = path
	}

	return selectedFolder
}

func MustConfig() *Config {
	selectedPath := fmt.Sprintf("%s/%s.yaml", getConfigFolder(), getSelectedEnv())

	cfg, err := loadConfig(selectedPath)
	if err != nil {
		panic(err)
	}

	return cfg
}

func MustLoadTestConfig() *Config {
	selectedPath := fmt.Sprintf("%s/%s.yaml", getConfigFolder(), Test)

	cfg, err := loadConfig(selectedPath)
	if err != nil {
		panic(err)
	}

	return cfg
}
