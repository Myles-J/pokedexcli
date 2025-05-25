package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/Myles-J/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("Pokemon is required")
	}

	pokemonName := args[0]

	pokemonData, err := pokeapi.GetPokemon(pokemonName, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)
	// Use the base experience of the pokemon to determine the catch rate
	random := rand.IntN(pokemonData.BaseExperience)

	if random > 40 {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
		cfg.caughtPokemon[pokemonData.Name] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	return nil
}
