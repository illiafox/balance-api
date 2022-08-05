package config

import "time"

type Redis struct {
	HashMap string `env:"REDIS_HASHMAP" env-default:"currency"`
	//
	Pass string `env:"REDIS_PASS" env-default:""`
	DB   int    `env:"REDIS_DB"   env-default:"0"`
	//
	IP   string `env:"REDIS_IP"   env-default:"0.0.0.0"`
	Port int    `env:"REDIS_PORT" env-default:"6379"`
	//
	Expire time.Duration `env:"REDIS_EXPIRE" env-default:"15m"`
}

type Postgres struct {
	User     string `env:"POSTGRES_USER"     env-required:""`
	Pass     string `env:"POSTGRES_PASSWORD" env-required:""`
	Database string `env:"POSTGRES_DB"       env-required:""`
	//
	IP   string `env:"POSTGRES_IP"   env-required:""`
	Port int    `env:"POSTGRES_PORT" env-default:"5432"`
}

// Host

type Host struct {
	Addr          string        `env:"HOST_ADDRESS"         env-default:"0.0.0.0"`
	Port          int           `env:"HOST_PORT"            env-default:"8080"`
	Key           string        `env:"HOST_KEY_PATH"`  // Path to TLS key
	Cert          string        `env:"HOST_CERT_PATH"` // Path to TLS certificate
	RequestTimout time.Duration `env:"HOST_REQUEST_TIMEOUT" env-default:"5s"`
}

//

type Config struct {
	Postgres Postgres
	Redis    Redis
	//
	Host Host
}
