package application_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/application"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewSocialApp(t *testing.T) {
	t.Run("should return not Nil when call", func(t *testing.T) {
		newApp := application.NewSocialApp()

		assert.NotNil(t, newApp)
	})
}

func TestAddUser(t *testing.T) {
	t.Run("should user added when call addUser", func(t *testing.T) {
		user := entity.NewUser("Alice")
		newApp := application.NewSocialApp()
		newApp.AddUser(user)

		result, _ := newApp.IsUserInApp("Alice")

		assert.True(t, result)
	})
}

func TestIsUserInApp(t *testing.T){
	t.Run("should return user when input username", func(t *testing.T) {
		user := entity.NewUser("Alice")
		newApp := application.NewSocialApp()
		newApp.AddUser(user)
		expected := user

		_, userInApp := newApp.IsUserInApp("Alice")

		assert.Equal(t, expected, userInApp)
	})

	t.Run("should return false when user not found", func(t *testing.T) {
		newApp := application.NewSocialApp()
		result, _ := newApp.IsUserInApp("Alice")

		assert.False(t, result)
	})
}

func TestTrending(t *testing.T) {
	t.Run("should return rank 1 Bob and rank 2 Alice", func(t *testing.T) {
		alice, bob, bill, john, cecil := "Alice", "Bob", "Bill", "John", "Ceci"
		Alice := entity.NewUser(alice)
		Bob := entity.NewUser(bob)
		Bill := entity.NewUser(bill)
		John := entity.NewUser(john)
		Cecil := entity.NewUser(cecil)
		app := application.NewSocialApp()
		app.AddUser(Alice)
		app.AddUser(Bob)
		app.AddUser(Bill)
		app.AddUser(John)
		Alice.Follow(Bob)
		Cecil.Follow(Alice)
		Alice.Follow(Bill)
		John.Follow(Bob)
		Bob.Follow(Alice)
		Bob.Follow(Bill)
		John.Follow(Alice)
		Alice.UploadPhoto()
		Bob.LikedPhoto(Alice)
		Bill.UploadPhoto()
		Bob.LikedPhoto(Bill)
		John.UploadPhoto()
		Bill.LikedPhoto(Bill)
		Alice.LikedPhoto(Bill)
		Cecil.UploadPhoto()
		expected := "Trending photos:\n1. Bill photo got 3 likes\n2. Alice photo got 1 like\n3. John photo got 0 like\n"

		trending := app.Trending()

		assert.Equal(t, expected, trending)
	})
}