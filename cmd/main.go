package main

import (
	"xamss.onelab.final/internal/config"
	"xamss.onelab.final/internal/registration/app"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
