package main

import "fmt"

func commandHelp() error {
	fmt.Println("Pokedex - a cli tool to look up any pokemon")
	fmt.Println("Commands:")
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}