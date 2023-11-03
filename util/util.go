package util

import (
	"errors"
	"fmt"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/application"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"
)

func alreadyUser(user1, user2 string, app *application.SocialApp)(*entity.User, *entity.User){
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

func ProcessSocialGraph(app *application.SocialApp, input string) (err error){
	words := strings.Split(input, " ")
	action := words[1]
	user1, user2 := words[0], words[2]
	if len(words) != 3 || !isWordAction(action) || user1 == "" || user2 == ""{
		return apperror.ErrInvalidKeyword
	}
	
	userOne, userTwo := alreadyUser(user1, user2, app)
	err = userOne.Follow(userTwo)
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

func ProcessUserAction(app *application.SocialApp, input string) (err error) {
	words := strings.Split(input, " ")
	user1 := words[0]
	action := words[1]
	user2 := words [2]
	if len(words) == 3{
		if action != constant.Uploaded || user1 == "" || user2 == ""{
			return apperror.ErrInvalidKeyword
		}
		ok, userOne :=app.IsUserInApp(user1)
		if !ok{
			return fmt.Errorf("unknown user %s", user1)
		}
		err = userOne.UploadPhoto()
	}else if len(words) == 4{
		if action != constant.Likes || words[3] != "photo"{
			return apperror.ErrInvalidKeyword
		}
		ok1, userOne :=app.IsUserInApp(user1)
		ok2, userTwo :=app.IsUserInApp(user2)
		if !ok1 || !ok2{
			return fmt.Errorf("unknown user %s", user1)
		}
		err = userOne.LikedPhoto(userTwo)
		if err != nil{
			switch{
			case errors.Is(err, apperror.ErrLikePhotoUserNotFollowedYet):
				err = fmt.Errorf("unable to like %s's photo", userTwo.UserName())
			case errors.Is(err, apperror.ErrUserDoesntHaveAPhoto):
				err = fmt.Errorf("%s doesn't have a photo", userTwo.UserName())	
			}
		}
	}
	return 
}

func ProcessDisplayActivity(app *application.SocialApp, input string)(err error){
	ok, userOne :=app.IsUserInApp(input)
	if !ok {
		return fmt.Errorf("unknown user %s", input)
	}
	fmt.Print(userOne.DisplayActivity())
	return
}