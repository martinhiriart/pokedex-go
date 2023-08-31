package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	pokemon := cfg.caughtPokemon

	if len(pokemon) == 0 {
		return errors.New("ERROR: no pokemon caught yet")
	}

	fmt.Println("Pokemon in your Pokedex:")
	fmt.Println("------------------------")
	for _, poke := range pokemon {
		fmt.Printf("%s\n", poke.Name)
	}
	fmt.Println("------------------------")
	fmt.Println()
	return nil
}
