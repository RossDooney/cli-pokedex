package main

import (
	"time"

	"github.com/RossDooney/cli-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	startRepl(cfg)
}
