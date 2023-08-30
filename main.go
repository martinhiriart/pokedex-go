package main

import (
	"time"

	"github.com/martinhiriart/pokedex-go/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cacheInterval := time.Hour
	cfg := config{
		pokeapiClient: pokeapi.NewClient(cacheInterval),
	}
	startCli(&cfg)
}
