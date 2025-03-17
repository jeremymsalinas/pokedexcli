package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(locations *Locations) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if locations.Next != "" {
		url = locations.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error occured getting locations")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading request body")
	}
	defer res.Body.Close()

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return fmt.Errorf("Error decoding response")
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil

}

func commandMapB(locations *Locations) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if locations.Previous != "" {
		url = locations.Previous
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error occured getting locations")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading request body")
	}
	defer res.Body.Close()

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return fmt.Errorf("Error decoding response")
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil

}
