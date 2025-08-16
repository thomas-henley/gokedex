package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const pokeurl = "https://pokeapi.co/api/v2/"

type LocationResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func commandMap(config *Config) error {
	locations, _ := getNextLocations(config)
	for _, location := range locations {
		fmt.Printf("%s\n", location)
	}
	return nil
}

func commandMapb(config *Config) error {
	locations, _ := getPrevLocations(config)
	for _, location := range locations {
		fmt.Printf("%s\n", location)
	}
	return nil
}

func getNextLocations(config *Config) ([]string, error) {
	if config.Next == "" {
		config.Next = pokeurl + "location-area/"
	}
	fmt.Println(config.Next)

	return getLocations(config, config.Next)
}

func getPrevLocations(config *Config) ([]string, error) {
    if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil, fmt.Errorf("you're on the first page")
	}

	return getLocations(config, config.Previous)
}

func getLocations(config *Config, url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	fmt.Println("check this")
	fmt.Println(res)

	data, err := io.ReadAll(res.Body)

	var response LocationResp
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json")
	}

	locations := []string{}
	for _, item := range response.Results {
		locations = append(locations, item.Name)
	}

	if response.Next != nil {
		config.Next = *response.Next
	}
	if response.Previous != nil {
		config.Previous = *response.Previous
	}

	return locations, nil
}

