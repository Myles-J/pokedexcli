package main

import "fmt"

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Display this help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Get the next 20 locations")
	fmt.Println("mapb: Get the previous 20 locations")
	fmt.Println("explore: Get the pokemon in the area")
	return nil
}
