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
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}
		command := cleanedInput[0]	

		switch command {
		case "exit":
			fmt.Println("pokedex aborted")
			os.Exit(0)
		default:
			fmt.Printf("invalid command: %v\n", command)
		}

		fmt.Println(cleanedInput)
	}
}

func cleanInput(inputStr string) []string {
	loweredStr := strings.ToLower(inputStr)
	strSlice := strings.Fields(loweredStr)
	return strSlice
}