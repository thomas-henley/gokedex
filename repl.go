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
	callback    func(*Config) error
}

type Config struct {
	next   string
	prev   string
}

func startRepl() {
	config := Config{"https://pokeapi.co/api/v2/location-area", ""}

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
		"map": {
			name:				 "map",
			description: "Get a list of 20 locations from the API",
			callback:		 commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous list of 20 locations",
			callback:    commandMapb,
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
		command.callback(&config)
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

