package main

import (
	"github.com/Myles-J/pokedexcli/internal/pokeapi"
	"github.com/Myles-J/pokedexcli/internal/pokecache"
)

type config struct {
	nextURL       string
	previousURL   string
	cache         *pokecache.Cache
	caughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Display this help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Get the next 20 locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Get the previous 20 locations",
		callback:    commandMapb,
	},

	"explore": {
		name:        "explore <location_name>",
		description: "Get the names of the pokemon in the area",
		callback:    commandExplore,
	},

	"catch": {
		name:        "catch <pokemon_name>",
		description: "Catch a pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect <pokemon_name>",
		description: "Inspect a pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "View your Pokedex",
		callback:    commandPokedex,
	},
}
