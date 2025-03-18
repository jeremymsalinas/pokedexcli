package main

import "fmt"

func commandInspect(cfg *config, pokemon_name string) error {
	if pokemon, ok := cfg.pokedex[pokemon_name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, types := range pokemon.Types {
			fmt.Printf("\t- %s\n", types.Type.Name)
		}
		return nil
	}
	fmt.Printf("%s not caught yet", pokemon_name)
	return nil
}
