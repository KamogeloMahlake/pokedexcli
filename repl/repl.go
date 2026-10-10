package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/KamogeloMahlake/pokedexcli/internal/pokeapi"
	"github.com/KamogeloMahlake/pokedexcli/internal/pokecache"
)


func StartPokedex() {
	Cfg := &Config{
		commands: getCommandsMap(),
		next: nil,
		previous: nil,
		cache: pokecache.NewCache(5 * time.Minute),
		client: pokeapi.Client{},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		
		input:= cleanInput(scanner.Text())

		command, exists := Cfg.commands[input[0]]
		if exists {
			err := command.callback(Cfg)
			
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}


