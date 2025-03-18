package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if cfg.currentAreaPokemon == nil {
		fmt.Println("you need to explore an area to catch pokemon!")
		return nil
	}

	pokemonList := *cfg.currentAreaPokemon
	found := false
	for i, pokemon := range pokemonList {
		if pokemonList[i] == pokemon {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("%s not found in %v\n", name, cfg.currentLocationName)
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(cfg.currentLocationURL, name)
	if err != nil {
		return err
	}

	exp := pokemon.BaseExperience

	chance := rand.Intn(exp)

	if chance > 50 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	fmt.Printf("%s added to Pokedex.\n", name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}
