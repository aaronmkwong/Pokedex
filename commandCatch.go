package main

import (
	"fmt"
	"math/rand"
)

// commandCatch attempts to catch a Pokemon and add it to the user's Pokedex
func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: catch <pokemon-name>")
		return nil
	}

	pokemonName := args[0]

	// Print first before API request / catch calculation
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Fetch pokemon data from PokeAPI
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// Catch logic:
	// Higher base experience = harder to catch
	//
	// Example:
	// BaseExperience = 64
	// rand.Intn(64) gives 0–63
	// if result < 40 => caught
	//
	// Lower base experience gives a smaller range,
	// increasing odds of rolling under the threshold.
	const catchThreshold = 40

	roll := rand.Intn(pokemon.BaseExperience)

	if roll < catchThreshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}