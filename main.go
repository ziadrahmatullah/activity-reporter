package main

import (
	"bufio"
	"fmt"
	"os"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/application"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/util"
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
	socialApp := application.NewSocialApp()

	for !exit{
		var err error
		fmt.Println(menu)
		input := promptInput(scanner, "Enter menu: ")

		switch input{
		case "1":
			socialGraph := promptInput(scanner, "Setup social graph: ")
			err = util.ProcessSocialGraph(socialApp, socialGraph)
		case "2":
			userAction := promptInput(scanner, "Enter user Actions: ")
			err = util.ProcessUserAction(socialApp, userAction)
		case "3":
			displayUser := promptInput(scanner, "Display activity for: ")
			err = util.ProcessDisplayActivity(socialApp, displayUser)
		case "4":
			fmt.Print(socialApp.Trending())
		case "5":
			fmt.Println("Good bye!")
			exit = true
		default:
			fmt.Println("invalid menu")
		}
		if err!= nil{
			fmt.Println(err)
		}
	}
}
