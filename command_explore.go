package main

import "errors"

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("ERROR: no location area provided")
	}
	locationAreaName = args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

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

func callbackMapB(cfg *config) error {

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

}
