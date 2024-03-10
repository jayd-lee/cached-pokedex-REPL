package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("aborting pokedex")
	os.Exit(0)
	return nil
}
