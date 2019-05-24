package env

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var (
	once   sync.Once
	Config config //Хранит переменные окружения
)

type config struct {
	DBHost     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort     int    `envconfig:"DB_PORT" default:"3306"`
	DBDatabase string `envconfig:"DB_DATABASE" default:"default"`
	DBUser     string `envconfig:"DB_USERNAME" default:"default"`
	DBPassword string `envconfig:"DB_PASSWORD" default:""`
	DBDebug    bool   `envconfig:"DB_DEBUG" default:"false"`
}

func init() {
	once.Do(func() {
		err := envconfig.Process("", &Config)
		if err != nil {
			fmt.Printf("Failed to read env: %v\n", err)
		}
	})
}
