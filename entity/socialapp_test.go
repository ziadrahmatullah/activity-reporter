package entity_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewSocialApp(t *testing.T) {
	t.Run("should return not Nil when call", func(t *testing.T) {
		newApp := entity.NewSocialApp()

		assert.NotNil(t, newApp)
	})
}

func TestAddUser(t *testing.T) {
	t.Run("should user added when call addUser", func(t *testing.T) {
		user := entity.NewUser("Alice")
		newApp := entity.NewSocialApp()
		newApp.AddUser(user)

		result, _ := newApp.IsUserInApp("Alice")

		assert.True(t, result)
	})
}

func TestIsUserInApp(t *testing.T){
	t.Run("should return user when input username", func(t *testing.T) {
		user := entity.NewUser("Alice")
		newApp := entity.NewSocialApp()
		newApp.AddUser(user)
		expected := user

		_, userInApp := newApp.IsUserInApp("Alice")

		assert.Equal(t, expected, userInApp)
	})

	t.Run("should return false when user not found", func(t *testing.T) {
		newApp := entity.NewSocialApp()
		result, _ := newApp.IsUserInApp("Alice")

		assert.False(t, result)
	})
}

func TestTrending(t *testing.T) {
	// t.Run()
}