package main

import (
	"fmt"
)

func commandMap(cfg *config) error{
	

	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Locations areas:")
	for _, location := range resp.Results{
		fmt.Printf(" - %s\n", location.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}