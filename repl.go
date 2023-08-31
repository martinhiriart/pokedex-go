package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startCli(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	replName := "pokedex-go"
	for {
		fmt.Printf("%v > ", replName)
		scanner.Scan()
		text := scanner.Text()

		cleanText := sanitizeInput(text)

		if len(cleanText) == 0 {
			continue
		}

		commandText := cleanText[0]

		args := []string{}

		if len(cleanText) > 1 {
			args = cleanText[1:]
		}

		cliCommands := getCommands()

		if command, exists := cliCommands[commandText]; exists {
			err := command.callback(cfg, args...)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("ERROR: invalid command\n")
			continue
		}
	}
}

func sanitizeInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help message",
			callback:    callbackHelp,
		},
		"clear": {
			name:        "clear",
			description: "Clears the current terminal screen",
			callback:    callbackClear,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex cli",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 location areas in the world",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 location areas in the world",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Shows all of the Pokemon in a given location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempts to catch the specified Pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {caught_pokemon_name}",
			description: "Views information of a Pokemon that you have caught",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all of the pokemon in your Pokedex",
			callback:    callbackPokedex,
		},
	}
}
