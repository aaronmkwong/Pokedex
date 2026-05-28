// from root of project on the command line
// go run .

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/aaronMkwong/Pokedex/internal/pokeapi"
	"github.com/aaronMkwong/Pokedex/internal/pokecache"
)


// cliCommand describes a REPL command
type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func main() {
	// Initialize cache
	cache := pokecache.NewCache(5 * time.Second)

	// Initialize PokeAPI client
	pokeClient := pokeapi.NewClient(cache)

	// Shared app config passed to every command
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	// Registry of supported commands
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
	}

	// Create scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// REPL loop
	for {
		// Prompt user
		fmt.Print("Pokedex > ")

		// Wait for input
		scanner.Scan()

		// Clean input into lowercase words
		words := CleanInput(scanner.Text())

		// Skip empty input
		if len(words) == 0 {
			continue
		}

		// First word is command name
		commandName := words[0]

		// Remaining words are arguments
		args := words[1:]

		// Look up command
		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		// Execute command
		err := command.callback(cfg, args)
		if err != nil {
			fmt.Println(err)
		}
	}
}