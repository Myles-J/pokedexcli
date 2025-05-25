package pokeapi

import (
	"encoding/json"
	"net/http"

	"github.com/Myles-J/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

func GetLocations(url string, cache *pokecache.Cache) ([]string, string, error) {
	if val, ok := cache.Get(url); ok {
		locations := []string{}
		if err := json.Unmarshal(val, &locations); err != nil {
			return nil, "", err
		}
		return locations, url, nil
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	var locations LocationResponse

	if err := json.NewDecoder(response.Body).Decode(&locations); err != nil {
		return nil, "", err
	}

	locationNames := make([]string, len(locations.Results))
	for i, location := range locations.Results {
		locationNames[i] = location.Name
	}

	return locationNames, locations.Next, nil
}

func GetPokemonForLocation(area string, cache *pokecache.Cache) ([]string, error) {
	url := baseURL + "/location-area/" + area
	if val, ok := cache.Get(url); ok {
		pokemon := []string{}
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return nil, err
		}
		return pokemon, nil
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var areaResponse AreaResponse

	if err := json.NewDecoder(response.Body).Decode(&areaResponse); err != nil {
		return nil, err
	}

	pokemonNames := make([]string, len(areaResponse.PokemonEncounters))
	for i, pokemonEncounter := range areaResponse.PokemonEncounters {
		pokemonNames[i] = pokemonEncounter.Pokemon.Name
	}

	return pokemonNames, nil
}

func GetPokemon(name string, cache *pokecache.Cache) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	if val, ok := cache.Get(url); ok {
		pokemon := Pokemon{}
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	response, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	var pokemonResponse Pokemon

	if err := json.NewDecoder(response.Body).Decode(&pokemonResponse); err != nil {
		return Pokemon{}, err
	}

	return pokemonResponse, nil
}
