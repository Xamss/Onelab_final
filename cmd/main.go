package main

import (
	"xamss.onelab.final/internal/config"
	"xamss.onelab.final/internal/registration/app"
)

// @title           Hospital Application
// @version         0.0.1
// @description     Hospital application where patients and doctors can decide on appointment time

// @contact.name   Bakhityar
// @contact.email  211022@astanait.edu.kz

// @host      localhost:8081
// @BasePath  /

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
