package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
}

func main() {
	// wait for input
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

		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Display this help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	trimmed := strings.Fields(lower)
	return trimmed
}
