package main

import "fmt"

func callbackHelp(cfg *config) error {

	commands := getCommands()

	fmt.Println()
	fmt.Println("Welcome to the pokedex-go cli!")
	fmt.Println()
	fmt.Println("Here are the available commands:")
	fmt.Println("--------------------------------")
	for _, cmd := range commands {
		fmt.Printf("%v - %v\n", cmd.name, cmd.description)
	}
	fmt.Println("--------------------------------")
	fmt.Println()

	return nil
}
