package app

import (
	"log"
	"os"
	"os/signal"
	"xamss.onelab.final/internal/config"
	"xamss.onelab.final/internal/registration/handler"
	pgrepo "xamss.onelab.final/internal/registration/repository/pg"
	"xamss.onelab.final/internal/registration/service"
	"xamss.onelab.final/pkg/httpserver"
	"xamss.onelab.final/pkg/jwttoken"
	"xamss.onelab.final/pkg/postgres"
)

func Run(cfg *config.Config) error {
	db, err := postgres.New(
		postgres.WithHost(cfg.DB.Host),
		postgres.WithPort(cfg.DB.Port),
		postgres.WithDBName(cfg.DB.DBName),
		postgres.WithUsername(cfg.DB.Username),
		postgres.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	repo := pgrepo.New(db.Pool)

	log.Println("connection success")

	migration := pgrepo.NewMigrate(cfg)

	err = migration.MigrateToVersion(cfg.DB.MigrationVersion)
	if err != nil {
		return err
	}

	token := jwttoken.NewToken(cfg.Token.SecretKey)
	srvs := service.New(repo, token, cfg)
	hndlr := handler.NewHandler(srvs)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
