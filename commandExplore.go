package main

import (
	"encoding/json"
	"fmt"
)

// ExploreResponse matches the JSON returned from:
// /location-area/{name}
type ExploreResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// commandExplore explores a location area and lists the Pokemon found there
func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: explore <location-area-name>")
		return nil
	}

	areaName := args[0]

	// Use endpoint path. Your pokeapi client should handle:
	// - base URL
	// - HTTP request
	// - cache lookup
	// - caching response
	endpoint := "/location-area/" + areaName

	body, err := cfg.pokeapiClient.Get(endpoint)
	if err != nil {
		return err
	}

	var resp ExploreResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")

	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}