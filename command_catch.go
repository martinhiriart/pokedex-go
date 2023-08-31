package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("ERROR: no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}
	const threshold = 50
	difficulty := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing Pokeball at %s...\n", pokemon.Name)
	if difficulty > threshold {
		return fmt.Errorf("%s escaped", pokemon.Name)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}
