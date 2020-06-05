package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/configor"
)

// Config is a config :)
type Config struct {
	PgURL                  string `env:"SIGNAL_PG_URL"`
	PgProto                string `env:"SIGNAL_PG_PROTO"`
	PgAddr                 string `env:"SIGNAL_PG_ADDR"`
	PgDb                   string `env:"SIGNAL_PG_DB"`
	PgUser                 string `env:"SIGNAL_PG_USER"`
	PgPassword             string `env:"SIGNAL_PG_PASSWORD"`
	CallRoomEndpoint       string `env:"SIGNAL_CALL_ROOM_ENDPOINT"`
	SchedulerLocation      string `env:"SIGNAL_SCHEDULER_LOCATION"`
	SchedulerTimeZone      string `env:"SIGNAL_SCHEDULER_TIMEZONE"`
	SchedulerMaxRetryCount int    `env:"SIGNAL_SCHEDULER_MAX_RETRY_COUNT"`
	SchedulerRetryPeriod   string `env:"SIGNAL_SCHEDULER_RETRY_PERIOD"`
	CallEndpoint           string `env:"SIGNAL_CALL_ENDPOINT"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment
func Get() *Config {
	once.Do(func() {
		envType := os.Getenv("SIGNAL_ENV")
		if envType == "" {
			envType = "dev"
		}
		if err := configor.New(&configor.Config{Environment: envType}).Load(&config, "config.json"); err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
