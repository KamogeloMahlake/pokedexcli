package repl
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


func StartPokedex() {
	cfg := &config{
		commands: getCommandsMap(),
		next: nil,
		previous: nil,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		
		input:= cleanInput(scanner.Text())

		command, exists := cfg.commands[input[0]]
		if exists {
			err := command.callback(cfg)
			
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


