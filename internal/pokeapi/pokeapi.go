package pokeapi

import (
	"encoding/json"
	"net/http"

	"github.com/Myles-J/pokedexcli/internal/pokecache"
)

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

	var locations struct {
		Results []struct {
			Name string `json:"name"`
		} `json:"results"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	}

	if err := json.NewDecoder(response.Body).Decode(&locations); err != nil {
		return nil, "", err
	}

	locationNames := make([]string, len(locations.Results))
	for i, location := range locations.Results {
		locationNames[i] = location.Name
	}

	return locationNames, locations.Next, nil
}

