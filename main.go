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

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	return commands
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}

func handleCommand(input string) {
	words := cleanInput(input)
	commands := getCommands()
	switch words[0] {
	case "exit":
		err := commands["exit"].callback()
		if err != nil {
			fmt.Printf("%v", err)
		}
	case "help":
		err := commands["help"].callback()
		if err != nil {
			fmt.Printf("%v", err)
		}
	default:
		fmt.Printf("'%s' is not a valid command\n\n", words[0])
		err := commands["help"].callback()
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		handleCommand(input)
	}
}

func main() {
	startRepl()
}
