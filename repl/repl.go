package repl
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


func StartPokedex() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandsMap()
		
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		
		input:= cleanInput(scanner.Text())
		command, exists := commands[input[0]]
		if exists {
			err := command.callback()
			
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


