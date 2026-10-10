package repl

import (
	"github.com/KamogeloMahlake/pokedexcli/internal/pokeapi"
	"github.com/KamogeloMahlake/pokedexcli/internal/pokecache"
)

type Config struct {
	commands map[string]cliCommand
	next *string
	previous *string
	cache pokecache.Cache
	client pokeapi.Client
}


type cliCommand struct {
	name string 
	description string
	callback func(*Config) error
}

func getCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit, 
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Each subsequent call to map should display the next 20 locations, and so on.",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "displays the previous 20 locations.",
			callback: mapB,
		},
	}	
}
