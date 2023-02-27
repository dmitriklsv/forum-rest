package main

import (
	"fmt"
	"log"

	"forum/internal/app"
	"forum/internal/controller"
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
	repos := service.NewRepos(db)
	log.Println("| repositories are ready to work!")

	log.Println("| constructing services...")
	services := controller.NewServices(repos)
	log.Println("| services are ready to work!")

	log.Println("| constructing handlers...")
	handlers := app.NewHandlers(services)
	log.Println("| handlers are ready to work!")

	if err := app.Run(handlers); err != nil {
		fmt.Println(err)
		return
	}
}
