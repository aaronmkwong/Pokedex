package main

import "github.com/aaronMkwong/Pokedex/internal/pokeapi"

// config stores shared application state used by commands
type config struct {
	Next          *string
	Previous      *string
	pokeapiClient *pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
}