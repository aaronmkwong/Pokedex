package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"

	// If we already have a next page stored, use it
	if cfg.Next != nil {
		url = *cfg.Next
	}

	resp, err := http.Get(url)
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