package main

import "fmt"

func callBackHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Here are your available command")
	availableCommands := getCommand()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s : %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
