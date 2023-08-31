package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Location areas:")
	fmt.Println("---------------")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	fmt.Println("---------------")
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
	fmt.Println("Location areas:")
	fmt.Println("---------------")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	fmt.Println("---------------")
	fmt.Println()

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
