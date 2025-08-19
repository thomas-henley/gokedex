package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
        fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
