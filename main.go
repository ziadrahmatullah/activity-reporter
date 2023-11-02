package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"
)

func promptInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func processSocialGraphValid(app *entity.SocialApp, input string) (err error){
	words := strings.Split(input, " ")
	if len(words) != 3 || words[1] != constant.Follows{
		return apperror.ErrInvalidKeyword
	}
	user1, user2 := words[0], words[2]
	ok1 , userOne := app.IsUserInApp(user1)
	if !ok1{
		userOne = entity.NewUser(user1)
		app.AddUser(userOne)
	}
	ok2 , userTwo := app.IsUserInApp(user2)
	if !ok2{
		userTwo = entity.NewUser(user2)
		app.AddUser(userTwo)
	}
	userOne.Follow(userTwo)
	return
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
	socialApp := entity.NewSocialApp()

	for !exit{
		fmt.Println(menu)
		input := promptInput(scanner, "Enter menu: ")

		switch input{
		case "1":
			socialGraph := promptInput(scanner, "Setup social graph: ")
			err := processSocialGraphValid(socialApp, socialGraph)
			if err != nil{
				fmt.Println(err)
			}
		case "2":
			
		}
	}
}
