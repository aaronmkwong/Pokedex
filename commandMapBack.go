package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// commandMapBack displays the previous 20 location areas
func commandMapBack(cfg *config, args []string) error {
	// If there is no previous page, we're at the beginning
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := http.Get(*cfg.Previous)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var locations locationAreaResponse

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locations); err != nil {
		return err
	}

	// Update pagination state
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	// Print location names
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}