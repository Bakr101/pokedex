package main

import "github.com/Bakr101/pokedex/internal/pokeapi"

type config struct{
	pokeApiClient		pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main(){
	cfg := config{
		pokeApiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)

}




