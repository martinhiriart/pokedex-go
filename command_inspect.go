package main

import (
	"errors"
	"fmt"
	"github.com/martinhiriart/pokedex-go/internal/styling"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("ERROR: no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("ERROR: this Pokemon has not been caught")
	}

	styling.PrintStyledMessage("Pokemon Stats:", "MenuHeader")
	styling.PrintStyledMessage("------------------------", "Menu")
	styling.PrintStyledMessage("Name: ", "StatDisplay")
	fmt.Println(pokemon.Name)
	styling.PrintStyledMessage("Height: ", "StatDisplay")
	fmt.Println(pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  %v\n", typ.Type.Name)
	}
	styling.PrintStyledMessage("------------------------", "Menu")
	return nil
}
