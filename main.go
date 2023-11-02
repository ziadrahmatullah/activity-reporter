package main

import (
	"bufio"
	"fmt"
	"os"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	exit := false
	menu :="Activity Reporter\n" +
		"1. Setup\n" +
		"2. Action\n" +
		"3. Display\n" +
		"4. Trending\n" +
		"5. Exit\n"

	for !exit{
		fmt.Println(menu)
		input := promptInput(scanner, "Enter menu: ")

		switch input{
		case "1":
			// socialGraph := promptInput(scanner, "Setup social graph: ")

		}
	}
}
