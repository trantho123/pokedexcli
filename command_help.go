package main

import "fmt"

func callBackHelp() {
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Here are your available command")
	availableCommands := getCommand()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s : %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
}
