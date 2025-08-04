package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name   string
	Url    string
}

type ResourceList struct {
	Count int
	Next string
	Previous string
	Results []Location
}

func commandMapb(config *Config) error {
  if config.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(config.prev)
	if err != nil {
		return err
	}
  defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	var resourceList ResourceList
	if err = json.Unmarshal(data, &resourceList); err != nil {
  	return fmt.Errorf("error unmarshaling json data: %w", err)
	}

	for _, location := range resourceList.Results {
		fmt.Printf("%v\n", location.Name)
	}
	config.next = resourceList.Next
	config.prev = resourceList.Previous
	return nil
}

func commandMap(config *Config) error {

	res, err := http.Get(config.next)
	if err != nil {
		return err
	}
  defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	var resourceList ResourceList
	if err = json.Unmarshal(data, &resourceList); err != nil {
  	return fmt.Errorf("error unmarshaling json data: %w", err)
	}

	for _, location := range resourceList.Results {
		fmt.Printf("%v\n", location.Name)
	}
	config.next = resourceList.Next
	config.prev = resourceList.Previous
	return nil
}  
