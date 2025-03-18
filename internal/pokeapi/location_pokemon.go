package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocationPokemon(pageURL *string, name string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + name
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationDetails{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationDetails{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	locationResp := LocationDetails{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationDetails{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
