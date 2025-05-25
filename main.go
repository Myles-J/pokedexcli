package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Myles-J/pokedexcli/internal/pokeapi"
	"github.com/Myles-J/pokedexcli/internal/pokecache"
	"github.com/Myles-J/pokedexcli/internal/utils"
)

func main() {
	cfg := &config{
		nextURL:       "https://pokeapi.co/api/v2/location-area",
		previousURL:   "",
		cache:         pokecache.NewCache(10 * time.Second),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		cleanedInput, err := utils.CleanInput(input)
		if err != nil {
			fmt.Println("Invalid command")
			continue
		}

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Invalid command")
			continue
		}

		args := cleanedInput[1:]
		commandErr := command.callback(cfg, args)
		if commandErr != nil {
			fmt.Println(commandErr)
		}
	}
}
