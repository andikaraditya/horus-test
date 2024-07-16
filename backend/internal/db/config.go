package db

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// MaxConn, IddleConn can be set here.
	// Refer to https://pkg.go.dev/github.com/jackc/pgx/v4@v4.13.0/pgxpool#ParseConfig
	DSN string `envconfig:"dsn" required:"true"`
}

var cfg config

func init() {
	if err := envconfig.Process("", &cfg); err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("initialize postgres")
	fmt.Println("config: ", cfg)
}
