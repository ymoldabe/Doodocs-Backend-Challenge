package main

import (
	"github.com/sirupsen/logrus"
	"github.com/ymoldabe/Doodocs-Backend-Challenge/configs"
	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/handler"
	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/server"
	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/service"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	cnf, err := configs.New()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	service := service.New()

	handler := handler.New(service)

	srv := new(server.Server)

	if err := srv.Run(cnf.Port, cnf.Host, handler.InitRouters()); err != nil {
		logrus.Fatalf("err of listening server, %s", err.Error())
	}

}
