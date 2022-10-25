package config

import (
	"strconv"
	"time"
)

const (
	NAMESPACE             = "test"
	RABBIT_MQ_CONNECTION  = "amqp://guest:guest@localhost:5672/"
	POSTGRESQL_CONNECTION = "host=host.docker.internal user=postgres password=polo1374 dbname=PLAY_GROUND port=5431 sslmode=disable"
	JWT_SECRET            = "secret_is_secret"
	JWT_TOKEN_LIFETIME    = 24 * time.Hour
	PORT                  = 8080
	API_TOKEN             = "belbel"
	SYSTEM_TOKEN          = "khar"
)

const (
	API_HOME_PAGE_MESSAGE = "Network Adapter"
)

func PORT_STRING() string {
	return strconv.Itoa(PORT)
}
