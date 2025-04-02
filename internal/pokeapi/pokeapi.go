package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocations(url string) ([]string, string, error) {
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
