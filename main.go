package main

import (
	"fmt"
	"log"
	"os"
	"practice-go/src/api"
	"practice-go/src/database"
	"practice-go/src/model"
)

func main() {
	// Connect to the database.
	var db, err = database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	var routes = api.GenerateRoutes(db)

	port := os.Getenv("PORT")
	routes.Run(":" + port)
}

func defunct() {
	fmt.Println("Hello, world.")
	system := model.MakeSystem("Dystopia Rising", "The system for DR 3.0.")
	org := model.MakeOrganization("DRGA", "The org for DRGA")
	game := model.MakeGame("Dystopia Rising Georgia", "The game DRGA.", org, system)
	//event := model.MakeEvent("Dystopia Rising Georgia", "The game DRGA.", org, system)

	fmt.Println("The game's name is " + game.Name)
}
