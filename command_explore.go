package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("ERROR: no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Pokemon found in %s:\n", locationArea.Name)
	fmt.Println("---------------")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	fmt.Println("---------------")
	fmt.Println()
	return nil
}
