package util_test

import (
	"fmt"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/application"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/util"
	"github.com/stretchr/testify/assert"
)

func TestProcessSocialGraph(t *testing.T) {
	t.Run("should add user in app when input Alice follows Bob", func(t *testing.T) {
		input := "Alice follows Bob"
		socialApp := application.NewSocialApp()
		util.ProcessSocialGraph(socialApp, input)

		user1, _ := socialApp.IsUserInApp("Alice")
		user2, _ := socialApp.IsUserInApp("Bob")

		assert.True(t, user1)
		assert.True(t, user2)
	})

	t.Run("should return ErrInvalidKeyword when input ' follows A'", func(t *testing.T) {
		input := " follows A"
		socialApp := application.NewSocialApp()
		expected := apperror.ErrInvalidKeyword

		err := util.ProcessSocialGraph(socialApp, input)

		assert.ErrorIs(t, err, expected)
	})

	t.Run("should return ErrInvalidKeyword when input 'B folllows A'", func(t *testing.T) {
		input := "B folllows A"
		socialApp := application.NewSocialApp()
		expected := apperror.ErrInvalidKeyword

		err := util.ProcessSocialGraph(socialApp, input)

		assert.ErrorIs(t, err, expected)
	})
}

func TestProcessUserAction(t *testing.T){
	t.Run("should return ErrInvalidKeyword when input ' uploaded photo'", func(t *testing.T) {
		input := " uploaded photo"
		socialApp := application.NewSocialApp()
		expected := apperror.ErrInvalidKeyword

		err := util.ProcessUserAction(socialApp, input)

		assert.ErrorIs(t, err, expected)
	})

	t.Run("should return unknown user Bob when Bob not in the app with input 'Bob uploaded photo'", func(t *testing.T) {
		input := "Bob uploaded photo"
		socialApp := application.NewSocialApp()
		expected := fmt.Errorf("unknown user Bob")

		err := util.ProcessUserAction(socialApp, input)

		assert.Equal(t, expected, err)
	})

	t.Run("should uploaded photo when input 'Bob upload photo'", func(t *testing.T) {
		socialApp := application.NewSocialApp()
		util.ProcessSocialGraph(socialApp, "Alice follows Bob")
		input := "Alice uploaded photo"
		util.ProcessUserAction(socialApp, input)
		_, Alice :=  socialApp.IsUserInApp("Alice") 

		eq := Alice.UserPhoto()

		assert.True(t, eq)
	})

	//========================
	t.Run("should return ErrInvalidKeyword when input 'A lkes B photo'", func(t *testing.T) {
		input := "A lkes B photo"
		socialApp := application.NewSocialApp()
		expected := apperror.ErrInvalidKeyword

		err := util.ProcessUserAction(socialApp, input)

		assert.ErrorIs(t, err, expected)
	})

	t.Run("should return unknown user Bob when Bob not in the app with input 'Bob likes Alice photo'", func(t *testing.T) {
		input := "Bob likes Alice photo"
		socialApp := application.NewSocialApp()
		expected := fmt.Errorf("unknown user Bob")

		err := util.ProcessUserAction(socialApp, input)

		assert.Equal(t, expected, err)
	})

	t.Run("should liked photo when input 'Alice likes Bob photo'", func(t *testing.T) {
		socialApp := application.NewSocialApp()
		util.ProcessSocialGraph(socialApp, "Alice follows Bob")
		util.ProcessUserAction(socialApp, "Bob uploaded photo")
		util.ProcessUserAction(socialApp, "Alice likes Bob photo")
		_, Bob :=  socialApp.IsUserInApp("Bob") 
		expected := 1

		count := Bob.ShowLikes()

		assert.Equal(t, expected, count)
	})

	t.Run("should return unable to like Alice's photo when Bob like Alice because bob hasn't followed Alice", func(t *testing.T) {
		socialApp := application.NewSocialApp()
		util.ProcessSocialGraph(socialApp, "Alice follows Bob")
		util.ProcessUserAction(socialApp, "Alice uploaded photo")
		expected := fmt.Errorf("unable to like Alice's photo")

		err := util.ProcessUserAction(socialApp, "Bob likes Alice photo")

		assert.Equal(t, expected, err)
	})

	t.Run("should return Alice doesn't have a photo when Bob like Alice because Alice doesn't uploaded photo", func(t *testing.T) {
		socialApp := application.NewSocialApp()
		util.ProcessSocialGraph(socialApp, "Alice follows Bob")
		expected := fmt.Errorf("Alice doesn't have a photo")

		err := util.ProcessUserAction(socialApp, "Bob likes Alice photo")

		assert.Equal(t, expected, err)
	})
}

func TestProcessDisplayActivity(t *testing.T) {
	t.Run("should return unknown user Bob when input Bob because Bob not in App", func(t *testing.T) {
		socialApp := application.NewSocialApp()
		expected := fmt.Errorf("unknown user Bob")

		err := util.ProcessDisplayActivity(socialApp, "Bob")

		assert.Equal(t, expected, err)
	})

	t.Run("should return Nil when success", func(t *testing.T) {
		socialApp := application.NewSocialApp()
		util.ProcessSocialGraph(socialApp, "Alice follows Bob")

		err := util.ProcessDisplayActivity(socialApp, "Bob")

		assert.Nil(t, err)
	})

	t.Run("should return ErrInvalidKeyword when input 'A lkes B photo'", func(t *testing.T) {
		socialApp := application.NewSocialApp()
		util.ProcessSocialGraph(socialApp, "Alice follows Bob")
		expected := apperror.ErrInvalidKeyword

		err := util.ProcessDisplayActivity(socialApp, "")

		assert.ErrorIs(t, err, expected)
	})
	
}