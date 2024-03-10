package main

import (
	"fmt"
	"log"
	"time"

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
		pokeapiclient: pokeapi.NewClient(time.Hour),
	}
	err := startRepl(&conf)
	if err != nil {
		log.Fatalln(err)
	}

}
