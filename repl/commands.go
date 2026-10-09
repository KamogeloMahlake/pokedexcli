package repl

type config struct {
	commands map[string]cliCommand
	next *string
	previous *string
}


type cliCommand struct {
	name string 
	description string
	callback func(*config) error
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
