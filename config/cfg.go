package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
)

const cfgFilePath = ".env"

type (
	Config struct {
		App      app
		RDB      rdb
		MDB      mongo
		Keycloak *keycloak
		Redis    redis
		Vault    vault
		ES       elasticseach
		S3       s3
		SES      *ses
		Discord  discord
	}

	app struct {
		Name        string `env:"APP_NAME"     env-required:"true"`
		Version     string `env:"APP_VERSION"  env-required:"true"`
		Env         string `env:"APP_ENV"      env-required:"true"`
		Port        string `env:"APP_PORT"     env-required:"true"`
		LogLevel    string `env:"LOG_LEVEL"    env-required:"true"`
		ServiceName string `env:"SERVICE_NAME" env-required:"true"`
		SchemasPath string `env:"SCHEMAS_PATH" env-required:"true"`
	}

	discord struct {
		WebhookUrl string `env:"DISCORD_WEBHOOK_URL" env-required:"true"`
	}
)

func NewConfig() *Config {
	cfg := &Config{}
	root := projectRoot()
	configFilePath := root + cfgFilePath

	err := loadCfg(configFilePath, cfg)
	if err != nil {
		panic(err)
	}

	cfg.RDB.MigrationsPath = root + cfg.RDB.MigrationsPath
	cfg.Vault.SecretIDFile = root + cfg.Vault.SecretIDFile
	return cfg
}

func loadCfg(cfgFilePath string, cfg *Config) error {
	envFileExists := checkFileExists(cfgFilePath)
	if envFileExists {
		err := cleanenv.ReadConfig(cfgFilePath, cfg)
		if err != nil {
			return fmt.Errorf("config error: %w", err)
		}
	} else {
		err := cleanenv.ReadEnv(cfg)
		if err != nil {
			if _, statErr := os.Stat(cfgFilePath); statErr != nil {
				return fmt.Errorf("missing environment variable: %w", err)
			}
			return err
		}
	}
	return nil
}

func checkFileExists(fileName string) bool {
	exist := false
	if _, err := os.Stat(fileName); err == nil {
		exist = true
	}
	return exist
}

func projectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	cwd := filepath.Dir(b)
	return cwd + "/../"
}
