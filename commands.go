package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("---Pokedex information---")
	for _, value := range CommandsMap() {
		fmt.Printf("Name: %s\n Description: %s\n", value.name, value.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Exiting program...")
	defer os.Exit(0)
	return nil
}

func CommandsMap() map[string]cliCommand {
	outputMap := make(map[string]cliCommand)

	outputMap["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	outputMap["exit"] = cliCommand{
		name:        "exit",
		description: "Exits the program",
		callback:    commandExit,
	}

	return outputMap
}
