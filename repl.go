package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print(" > ")
		scanner.Scan()
		text := scanner.Text()
		clearned := cleanInput(text)
		if len(clearned) == 0 {
			continue
		}
		commandName := clearned[0]
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Print help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Closes program",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}
