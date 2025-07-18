package config

type redis struct {
	Host           string `env:"REDIS_HOST"             env-required:"true"`
	Port           string `env:"REDIS_PORT"             env-required:"true"`
	ClientName     string `env:"REDIS_CLIENT_NAME"      env-required:"true"`
	MaxRetries     int    `env:"REDIS_MAX_RETRIES"      env-default:"3"`
	PoolSize       int    `env:"REDIS_POOL_SIZE"        env-default:"10"`
	MaxIdleConns   int    `env:"REDIS_MAX_IDLE_CONNS"   env-default:"5"`
	MaxActiveConns int    `env:"REIDS_MAX_ACTIVE_CONNS" env-default:"10"`
	MaxIdleTime    int    `env:"REIDS_MAX_IDLE_TIME"    env-default:"30"`
	MaxLifeTime    int    `env:"REDIS_MAX_LIFE_TIME"    env-default:"10"`
}

type RedisCredentials struct {
	Username string
	Password string
}
