package main

import (
	"fmt"
	"time"

	"github.com/thomas-henley/gokedex/internal/pokeapi"
)

func handleCommand(input string, config *config) {
	words := cleanInput(input)
	commands := getCommands()
	switch words[0] {
	case "exit":
		err := commands["exit"].callback(config)
		if err != nil {
			fmt.Printf("%v", err)
		}
	case "help":
		err := commands["help"].callback(config)
		if err != nil {
			fmt.Printf("%v", err)
		}
	case "map":
		err := commands["map"].callback(config)
		if err != nil {
			fmt.Printf("%v", err)
		}
	case "mapb":
		err := commands["mapb"].callback(config)
		if err != nil {
			fmt.Printf("%v", err)
		}
	default:
		fmt.Printf("'%s' is not a valid command\n\n", words[0])
		err := commands["help"].callback(config)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
