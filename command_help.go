package main

import (
	"fmt"

	"github.com/martinhiriart/pokedex-go/internal/styling"
)

func callbackHelp(cfg *config, args ...string) error {

	commands := getCommands()

	fmt.Println()
	styling.PrintStyledMessage(fmt.Sprint("Welcome to the pokedex-go cli!"), "MenuHeader")
	fmt.Println()
	styling.PrintStyledMessage(fmt.Sprint("Here are the available commands:"), "Menu")
	styling.PrintStyledMessage(fmt.Sprint("--------------------------------"), "Menu")
	for _, cmd := range commands {
		fmt.Printf("%v - %v\n", cmd.name, cmd.description)
	}
	styling.PrintStyledMessage(fmt.Sprint("--------------------------------"), "Menu")
	fmt.Println()

	return nil
}
