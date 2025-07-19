package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server ConfServer
}

type ConfServer struct {
	Port         string        `env:"SERVER_PORT,default=8080"`
	Host         string        `env:"SERVER_HOST,default=localhost"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,default=10s"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,default=10s"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,default=10s"`
	Debug        bool          `env:"SERVER_DEBUG,default=false"`
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
