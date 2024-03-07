package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEcho :")
		scanner.Scan()
		text := scanner.Text()
		inputWord := cleanInput(text)
		if len(inputWord) == 0 {
			continue
		}
		commandName := inputWord[0]
		availabCommands := getCommand()
		command, ok := availabCommands[commandName]
		if !ok {
			fmt.Printf("invalid command")
			continue
		}
		command.callback()
	}
}

func cleanInput(str string) []string {
	lowerd := strings.ToLower(str)
	words := strings.Fields(lowerd)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help",
			callback:    callBackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turn off the commmand",
			callback:    callBackExit,
		},
	}
}
