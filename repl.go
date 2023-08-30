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

		cliCommands := getCommands()

		if command, exists := cliCommands[commandText]; exists {
			err := command.callback(cfg)

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
	callback    func(*config) error
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
	}
}
