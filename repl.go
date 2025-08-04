package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Print usage information",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Gokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		command, ok := commands[input[0]]
    if !ok {
			fmt.Println("Unknown command.")
			continue
		}
		command.callback()
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

