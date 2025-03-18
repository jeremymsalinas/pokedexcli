package main

import (
	"fmt"
)

func commandExplore(cfg *config, area_name string) error {
	exploreResp, err := cfg.pokeapiClient.ListLocationPokemon(cfg.exploredLocations, area_name)
	if err != nil {
		return err
	}
	cfg.prevLocationsURL = exploreResp.Location.URL

	for _, loc := range exploreResp.PokemonEncounters {
		fmt.Println("- "+loc.Pokemon.Name)
	}
	return nil
}