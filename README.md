# pokedexcli

A command-line Pokédex application written in Go. This tool allows you to explore Pokémon data, locations, and more directly from your terminal, using the [PokeAPI](https://pokeapi.co/).

## Features
- Browse Pokémon locations and areas
- Explore which Pokémon appear in each area
- Catch and inspect Pokémon
- Maintain your own Pokédex of caught Pokémon
- Caching for efficient repeated queries

## Requirements
- Go 1.24.1 or later

## Installation
Clone the repository and build the binary:

```sh
git clone https://github.com/Myles-J/pokedexcli.git
cd pokedexcli
go build -o pokedexcli
```

## Usage
Run the CLI:

```sh
./pokedexcli
```

You will be greeted with a prompt:

```
Pokedex >
```

Type commands to interact with the Pokédex. Available commands:

| Command                | Description                   |
| ---------------------- | ----------------------------- |
| help                   | Display this help message     |
| exit                   | Exit the Pokedex              |
| map                    | Get the next 20 locations     |
| mapb                   | Get the previous 20 locations |
| explore <location>     | Get the Pokémon in the area   |
| catch <pokemon_name>   | Catch a Pokémon               |
| inspect <pokemon_name> | Inspect a caught Pokémon      |
| pokedex                | View your Pokédex             |

Example session:

```
Pokedex > map
Pokedex > explore pallet-town-area
Pokedex > catch pikachu
Pokedex > pokedex
Pokedex > inspect pikachu
```

## Project Structure
- `main.go` - Entry point and REPL loop
- `command_*.go` - Individual command implementations
- `internal/pokeapi/` - PokeAPI client and types
- `internal/pokecache/` - Simple in-memory cache
- `internal/utils/` - Utility functions

## API Reference
This project uses [PokeAPI v2](https://pokeapi.co/docs/v2).

## License
This project is for educational purposes and is not affiliated with Nintendo, Game Freak, or The Pokémon Company. Pokémon and Pokémon character names are trademarks of Nintendo.