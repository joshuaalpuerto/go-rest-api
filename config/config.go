// Main purpose is to load the configuration and start the app
package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Version string `env:"VERSION,default=v1"`
	Server  ConfServer
	DB      DBConf
}

type ConfServer struct {
	Port         string        `env:"SERVER_PORT,default=8080"`
	Host         string        `env:"SERVER_HOST,default=localhost"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,default=10s"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,default=10s"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,default=10s"`
	Debug        bool          `env:"SERVER_DEBUG,default=false"`
}

type DBConf struct {
	Host     string `env:"DB_HOST,default=localhost"`
	Port     string `env:"DB_PORT,default=5432"`
	User     string `env:"DB_USER,default=postgres"`
	Password string `env:"DB_PASSWORD,default=postgres"`
	Database string `env:"DB_NAME,default=postgres"`
}

func New() Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return c
}
