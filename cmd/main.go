package main

import (
	"fmt"
	"log"

	"forum/internal/app"
	"forum/internal/controller"
	"forum/internal/repository"
	"forum/internal/service"
	"forum/pkg/sqlite3"
)

func main() {
	log.Println("| connecting database...")
	db, err := sqlite3.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("| database connected!")
	defer db.Collection.Close()

	log.Println("| constructing repositories...")
	repos := repository.NewRepos(db)
	log.Println("| repositories are ready to work!")

	log.Println("| constructing services...")
	services := service.NewServices(repos)
	log.Println("| services are ready to work!")

	log.Println("| constructing handlers...")
	handlers := controller.NewHandlers(services)
	log.Println("| handlers are ready to work!")

	if err := app.Run(handlers); err != nil {
		fmt.Println(err)
		return
	}
}
