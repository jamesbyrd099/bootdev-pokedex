package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := CommandsMap()
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		textRaw := scanner.Text()
		if len(textRaw) == 0 {
			continue
		}
		parsed := strings.Fields(strings.ToLower(textRaw))
		text := parsed[0]
		command, exists := commands[text]
		if !exists {
			fmt.Println("Command doesn't exist...")
			continue
		}
		fmt.Println("")
		if err := command.callback(); err != nil {
			fmt.Printf("\nerror executing %s: %v\n", command.name, err)
		}
	}
}
