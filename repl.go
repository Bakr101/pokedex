package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl (cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		command, exists := getCommands()[commandName]
		if exists{
			err:= command.callback(cfg)
			if err != nil{
				fmt.Println(err)
			}
			
		}else{
			fmt.Println("Invalid command")
			continue
		}	
		
		}
}

func cleanInput(str string) []string{
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct{
	name		string
	description	string
	callback	func(*config) error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "displays the name of 20 map areas",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "displays the previous 20 locations",
			callback: commandMapb,
		},
	}
}