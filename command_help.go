package main

import "fmt"

func commandHelp(config *Config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
  return nil
}
