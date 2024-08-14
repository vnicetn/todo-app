package main

import (
	"log"

	"github.com/vnicetn/todo-app"
	"github.com/vnicetn/todo-app/pkg/handler"
	"github.com/vnicetn/todo-app/pkg/repository"
	"github.com/vnicetn/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
