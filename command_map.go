package main

import (
	"errors"
	"fmt"
	"github.com/martinhiriart/pokedex-go/internal/styling"
)

func callbackMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		return err
	}
	fmt.Println()
	styling.PrintStyledMessage("Location areas:", "MenuHeader")
	styling.PrintStyledMessage("---------------", "Menu")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	styling.PrintStyledMessage("---------------", "Menu")
	fmt.Println()

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func callbackMapB(cfg *config, args ...string) error {

	if cfg.previousLocationAreaURL == nil {
		return errors.New("ERROR: On the first page. No previous pages to navigate to")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)

	if err != nil {
		return err
	}
	fmt.Println()
	styling.PrintStyledMessage("Location areas:", "MenuHeader")
	styling.PrintStyledMessage("---------------", "Menu")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	styling.PrintStyledMessage("---------------", "Menu")
	fmt.Println()

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
