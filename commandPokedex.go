package main

import "fmt"

// commandPokedex displays all Pokemon the user has caught.
func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}