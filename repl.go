package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl () {
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
			err:= command.callback()
			if err != nil{
				fmt.Println(err)
			}
			continue
		}else{
			fmt.Println("Unkown command")
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
	callback	func() error
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
	}
}