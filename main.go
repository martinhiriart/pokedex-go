package main

import (
	"time"

	"github.com/martinhiriart/pokedex-go/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	cacheInterval := time.Hour
	cfg := config{
		pokeapiClient: pokeapi.NewClient(cacheInterval),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startCli(&cfg)
}
