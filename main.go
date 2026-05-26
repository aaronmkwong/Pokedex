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

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config{
		Next:     nil,
		Previous: nil,
	}

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
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		words := CleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands[commandName]
		if exists {
			if err := command.callback(cfg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}