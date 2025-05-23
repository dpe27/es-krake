package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
)

const defaultCfgFilePath = "config/default.yml"

type (
	Config struct {
		App    `yaml:"application"`
		Log    `yaml:"logger"`
		Server `yaml:"server"`
		RDB    `yaml:"rdb"`
		// Redis  `yaml:"redis"`
		// ES     `yaml:"elastic_search"`
	}

	App struct {
		Name    string `yaml:"name" env:"APP_NAME" env-required:"true"`
		Version string `yaml:"version" env:"APP_VERSION" env-required:"true"`
		Env     string `yaml:"env" env:"APP_ENV" env-required:"true"`
	}

	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-required:"true"`
	}

	Log struct {
		Level string `yaml:"level" env:"LOG_LEVEL" env-required:"true"`
	}

	RDB struct {
		Driver   string `yaml:"driver" env:"DB_DRIVER" env-required:"true"`
		Host     string `yaml:"host" env:"DB_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"DB_PORT" env-required:"true"`
		Username string `yaml:"username" env:"DB_USERNAME" env-required:"true"`
		Password string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
		Name     string `yaml:"name" env:"DB_NAME" env-required:"true"`
		SSLMode  string `yaml:"ssl_mode" env:"DB_SSLMODE" env-default:"disable"`

		MaxOpenConns int `yaml:"max_open_conns" env:"DB_MAX_OPEN_CONNS" env-required:"true"`
		MaxIdleConns int `yaml:"max_idle_conns" env:"DB_MAX_IDLE_CONNS" env-required:"true"`
		MaxLifeTime  int `yaml:"max_life_time" env:"DB_MAX_LIFE_TIME" env-required:"true"`
		MaxIdleTime  int `yaml:"max_idle_time" env:"DB_MAX_IDLE_TIME" env-required:"true"`
		ConnTimeout  int `yaml:"conn_timeout" env:"DB_CONN_TIMEOUT" env-default:"10000"`
		ConnAttempts int `yaml:"conn_attempts" env:"DB_CONN_ATTEMPTS" env-default:"10"`

		MigrationsPath string `yaml:"migrations_path" env:"DB_MIGRATIONS_PATH" env-required:"true"`
	}

	Redis struct {
		Host       string `yaml:"host" env:"REDIS_HOST" env-required:"true"`
		Port       string `yaml:"port" env:"REDIS_PORT" env-required:"true"`
		ClientName string `yaml:"client-name" env:"REDIS_CLIENT_NAME" env-required:"true"`
		Username   string `yaml:"username" env:"REDIS_USERNAME" env-required:"true"`
		Password   string `yaml:"password" env:"REDIS_PASSWORD" env-required:"true"`

		MaxRetries     int ` yaml:"max_retries" env:"REDIS_MAX_RETRIES" env-default:"3"`
		PoolSize       int `yaml:"pool_size" env:"REDIS_POOL_SIZE" env-default:"10"`
		MaxIdleConns   int `yaml:"max_idle_conns" env:"REDIS_MAX_IDLE_CONNS" env-default:"0"`
		MaxActiveConns int `yaml:"max_active_conns" env:"REIDS_MAX_ACTIVE_CONNS" env-default:"10"`
		MaxIdleTime    int `yaml:"max_idle_time" env:"REIDS_MAX_IDLE_TIME" env-default:"30"`
		MaxLifeTime    int `yaml:"max_life_time" env:"REDIS_MAX_LIFE_TIME" env-default:"10"`
	}

	ES struct {
		Addresses []string `yaml:"addresses" env:"ES_ADDRESSES" env-required:"true"`
		Username  string   `yaml:"username" env:"ES_USERNAME" env-required:"true"`
		Password  string   `yaml:"password" env:"ES_PASSWORD" env-required:"true"`

		MaxRetries    int  `yaml:"max_retries" env:"ES_MAX_RETRIES" env-default:"3"`
		EnableMetrics bool `yaml:"enable_metrics" env:"ES_ENABLE_METRICS" env-default:"false"`
		Debug         bool `yaml:"debug" env:"ES_DEBUG" env-default:"true"`
	}

	Keycloak struct{}

	Kafka struct{}

	Discord struct {
		Webhook string `yaml:"webhook" env:"DISCORD_WEBHOOK" env-required:"true"`
	}
)

func NewConfig() *Config {
	cfg := &Config{}
	root := projectRoot()
	configFilePath := root + defaultCfgFilePath

	err := loadCfgFile(configFilePath, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func loadCfgFile(cfgFilePath string, cfg *Config) error {
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
