package pokeapi

import (
	"encoding/json"
	"fmt"
)

// GetPokemon fetches a Pokemon by name from the PokeAPI,
// unmarshals the JSON response into a Pokemon struct,
// and returns it.
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	endpoint := "/pokemon/" + name

	body, err := c.Get(endpoint)
	if err != nil {
		return Pokemon{}, fmt.Errorf("could not fetch pokemon %s: %w", name, err)
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("could not parse pokemon data: %w", err)
	}

	return pokemon, nil
}