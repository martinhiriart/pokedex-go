package main

import (
	"errors"
	"fmt"
	"github.com/martinhiriart/pokedex-go/internal/styling"
)

func callbackPokedex(cfg *config, args ...string) error {

	pokemon := cfg.caughtPokemon

	if len(pokemon) == 0 {
		return errors.New("ERROR: no pokemon caught yet")
	}

	styling.PrintStyledMessage("Pokemon in your Pokedex:", "MenuHeader")
	styling.PrintStyledMessage("------------------------", "Menu")
	for _, poke := range pokemon {
		fmt.Printf("%s\n", poke.Name)
	}
	styling.PrintStyledMessage("------------------------", "Menu")
	fmt.Println()
	return nil
}
