package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Myles-J/pokedexcli/internal/pokeapi"
)

type config struct {
	nextURL     string
	previousURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
}

func main() {
	cfg := &config{
		nextURL:     "https://pokeapi.co/api/v2/location-area",
		previousURL: "",
	}
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Display this help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Get the next 20 locations")
	fmt.Println("mapb: Get the previous 20 locations")
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
	locations, nextURL, err := pokeapi.GetLocations(cfg.nextURL)
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

func commandMapb(cfg *config) error {
	if cfg.previousURL == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locations, nextURL, err := pokeapi.GetLocations(cfg.previousURL)
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

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	trimmed := strings.Fields(lower)
	return trimmed
}
