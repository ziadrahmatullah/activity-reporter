package util

import (
	"fmt"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"
)

func alreadyUser(user1, user2 string, app *entity.SocialApp)(*entity.User, *entity.User){
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
	return userOne, userTwo
}

func ProcessSocialGraphValid(app *entity.SocialApp, input string) (err error){
	words := strings.Split(input, " ")
	if len(words) != 3 || !isWordAction(words[1]){
		return apperror.ErrInvalidKeyword
	}
	user1, user2 := words[0], words[2]
	userOne, userTwo := alreadyUser(user1, user2, app)
	userOne.Follow(userTwo)
	return
}

func isWordAction(input string) bool{
	for _, action := range constant.Actions {
		if action == input{
			return true
		}
	}
	return false
}

func ProcessUserAction(app *entity.SocialApp, input string) (err error) {
	words := strings.Split(input, " ")
	user1 := words[0]
	action := words[1]
	user2 := words [2]
	if len(words) == 3{
		if action != constant.Uploaded{
			return apperror.ErrInvalidKeyword
		}
		ok, userOne :=app.IsUserInApp(user1)
		if !ok{
			return fmt.Errorf("unknown user %s", user1)
		}
		userOne.UploadFoto()
	}else if len(words) == 4{
		if action != constant.Likes || words[3] != "photo"{
			return apperror.ErrInvalidKeyword
		}
		ok1, userOne :=app.IsUserInApp(user1)
		ok2, userTwo :=app.IsUserInApp(user2)
		if !ok1 || !ok2{
			return fmt.Errorf("unknown user %s", user1)
		}
		userOne.LikedPhoto(userTwo)
	}
	return
}

func ProcessDisplayActivity(app *entity.SocialApp, input string)(err error){
	ok, userOne :=app.IsUserInApp(input)
	if !ok {
		return fmt.Errorf("unknown user %s", input)
	}
	fmt.Print(userOne.DisplayActivity())
	return
}