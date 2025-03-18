package main

import (
	"fmt"
)

func commandExplore(cfg *config, area_name string) error {
	exploreResp, err := cfg.pokeapiClient.ListLocationPokemon(cfg.exploredLocations, area_name)
	if err != nil {
		return err
	}
	cfg.currentLocationURL = exploreResp.Location.URL
	cfg.currentLocationName = exploreResp.Location.Name
	catchablePokemon := []string{}
	fmt.Printf("Exploring %s...\n",area_name)
	fmt.Println("Found Pokemon:")
	for _, loc := range exploreResp.PokemonEncounters {
		fmt.Println("- "+loc.Pokemon.Name)
		catchablePokemon = append(catchablePokemon, loc.Pokemon.Name)
	}
	cfg.currentAreaPokemon = &catchablePokemon
	return nil
}