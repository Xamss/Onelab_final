package main

import (
	"xamss.onelab.final/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

}
