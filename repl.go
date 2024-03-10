package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(conf *config) error {
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
		err := command.callback(conf)
		if err != nil {
			return err
		}
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
	callback    func(conf *config) error
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
		"map": {
			name:        "Next",
			description: "Prints area the commmand",
			callback:    callBackMapNext,
		},
		"mapv": {
			name:        "Prev",
			description: "Prints area the commmand",
			callback:    callBackMapPrev,
		},
	}
}
