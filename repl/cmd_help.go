package repl

import "fmt"

func commandHelp(conf *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, c := range getCommandsMap() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
