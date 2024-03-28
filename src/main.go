package main

import (
	"fmt"
	"log"
	"practice-go/api"
	"practice-go/database"
	"practice-go/model"
)

func main() {
	// Connect to the database.
	var db, err = database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	var routes = api.GenerateRoutes(db)

	routes.Run(":8080")
}

func defunct() {
	fmt.Println("Hello, world.")
	system := model.MakeSystem("Dystopia Rising", "The system for DR 3.0.")
	org := model.MakeOrganization("DRGA", "The org for DRGA")
	game := model.MakeGame("Dystopia Rising Georgia", "The game DRGA.", org, system)
	//event := model.MakeEvent("Dystopia Rising Georgia", "The game DRGA.", org, system)

	fmt.Println("The game's name is " + game.Name)
}
