package main

import (
	"diploma/authorization/internal/config"
	"diploma/authorization/internal/delivery/app"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
