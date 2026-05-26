// run from project root folder to test: go run .   
// terminate the program by pressing ctrl+c
// can run the CLIand tee the output (copies the stdout) to a new file called repl.log
// and .gitignore the log

package main

import (
	"bufio"
	"fmt"
	"os"
)

// cliCommand represents a command the REPL can run.
// Each command has:
// - a name
// - a description (used in help output)
// - a callback function to execute
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	// Scanner reads user input from standard input (keyboard)
	scanner := bufio.NewScanner(os.Stdin)

	// Registry of all supported commands
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
	}

	// Start REPL loop
	// This runs continuously until the user exits the program
	for {
		// Display prompt without a newline
		fmt.Print("Pokedex > ")

		// Wait for user input
		scanner.Scan()
		input := scanner.Text()

		// Normalize and split input into words
		words := CleanInput(input)

		// Ignore empty input
		if len(words) == 0 {
			continue
		}

		// First word is the command name
		commandName := words[0]

		// Look up command in registry
		command, exists := commands[commandName]

		if exists {
			// Execute command callback
			// Print any returned error
			if err := command.callback(); err != nil {
				fmt.Println(err)
			}
		} else {
			// Handle unknown command
			fmt.Println("Unknown command")
		}
	}
}