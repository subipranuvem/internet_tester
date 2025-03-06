package model

import "time"

type AppConfig struct {
	DBConnStr                     string        `env:"DB_CONN_STR" envDefault:"host=postgres user=user password=password dbname=pingdb port=5432 sslmode=disable"`
	TargetURL                     string        `env:"TARGET_URL" envDefault:"https://google.com.br"`
	TimeoutInSecs                 time.Duration `env:"TIMEOUT_IN_SECS" envDefault:"5s"`
	IntervalBetweenRequestsInSecs time.Duration `env:"INTERVAL_BETWEEN_REQUESTS_IN_SECS" envDefault:"10s"`
}
