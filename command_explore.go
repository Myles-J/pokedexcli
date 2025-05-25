package main

import (
	"errors"
	"fmt"

	"github.com/Myles-J/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("area is required")
	}

	area := args[0]

	pokemon, err := pokeapi.GetPokemonForLocation(area, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemon {
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}
