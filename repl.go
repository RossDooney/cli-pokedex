package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "list next locations",
			callback:    callbackMap,
		},
		"back": {
			name:        "map",
			description: "list previous locations",
			callback:    callbackMapBack,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}
