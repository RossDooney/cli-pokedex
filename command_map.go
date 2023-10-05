package main

import (
	"fmt"
	"log"

	"github.com/RossDooney/cli-pokedex/internal/pokeapi"
)

func callbackMap(cfg *config) error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s \n", area.Name)
	}
	return nil

}
