package main

import (
	"fmt"

	"github.com/trantho123/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiclient    pokeapi.Client
	locationAreaNext *string
	locationAreaPre  *string
}

func main() {
	fmt.Print("start ...\n")
	conf := config{
		pokeapiclient: pokeapi.NewClient(),
	}
	startRepl(&conf)

}
