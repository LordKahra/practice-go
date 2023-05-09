package main

import (
	"fmt"
	"practice-go/larp"
)

func main() {
	fmt.Println("Hello, world.")
	system := larp.MakeSystem("Dystopia Rising", "The system for DR 3.0.")
	org := larp.MakeOrganization("DRGA", "The org for DRGA")
	game := larp.MakeGame("Dystopia Rising Georgia", "The game DRGA.", org, system)

	fmt.Println("The game's name is " + game.Name)

}
