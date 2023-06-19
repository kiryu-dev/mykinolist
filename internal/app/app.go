package app

import (
	"log"
	"net/http"

	"github.com/kiryu-dev/mykinolist/internal/controller"
	"github.com/kiryu-dev/mykinolist/internal/infrastructure/repository"
	"github.com/kiryu-dev/mykinolist/internal/service"
)

func Run(config *Config) {
	db, err := repository.NewPostgresDB((*repository.Config)(config.DB))
	if err != nil {
		log.Fatal(err.Error())
	}
	repo := repository.New(db)
	service := service.New(repo)
	controller := controller.New(service)
	if err := http.ListenAndServe(config.ListeningPort, controller); err != nil {
		log.Fatal(err.Error())
	}
}
