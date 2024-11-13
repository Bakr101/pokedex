package main

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *config)error{
	if cfg.prevLocationAreaURL == nil{
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.prevLocationAreaURL)

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