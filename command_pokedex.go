package main

import "fmt"

func commandPokedex(cfg *config, pokemon_name string) error {
	fmt.Println("Pokedex:")
	for _, name := range cfg.pokedex {
		fmt.Printf(" - %s\n",name.Name)
	}
	return nil
}
