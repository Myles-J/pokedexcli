package main

import (
	"fmt"

	"github.com/Myles-J/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, args []string) error {
	locations, nextURL, err := pokeapi.GetLocations(cfg.nextURL, cfg.cache)
	if err != nil {
		return err
	}

	cfg.previousURL = cfg.nextURL
	cfg.nextURL = nextURL

	fmt.Println("You can choose from the following locations:")
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.previousURL == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locations, nextURL, err := pokeapi.GetLocations(cfg.previousURL, cfg.cache)
	if err != nil {
		return err
	}

	cfg.nextURL = cfg.previousURL
	cfg.previousURL = nextURL

	fmt.Println("You can choose from the following locations:")
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
