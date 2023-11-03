package entity_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("should return not nil when call", func(t *testing.T) {
		user := "Alice"
		
		newUser := entity.NewUser(user)

		assert.NotNil(t, newUser)
	})
}


func TestUploadPhoto(t *testing.T){
	t.Run("should return added photo when call", func(t *testing.T) {
		user1:= "Alice"
		newUser1 := entity.NewUser(user1)

		newUser1.UploadPhoto()
		eq := newUser1.UserPhoto()

		assert.True(t, eq)
	})
	t.Run("should return ErrConnotUploadMorePhoto when upload photo 2 times", func(t *testing.T) {
		user1:= "Alice"
		newUser1 := entity.NewUser(user1)
		expected := apperror.ErrCannotUploadMorePhoto

		newUser1.UploadPhoto()
		err := newUser1.UploadPhoto()

		assert.ErrorIs(t, err, expected)
	})

	t.Run("should return notify Alice uploaded photo to Alice follower", func(t *testing.T) {
		user1, user2 := "Alice", "Bob"
		newUser1 := entity.NewUser(user1)
		newUser2 := entity.NewUser(user2)
		expected := "\nBob activities:\nAlice uploaded photo\n"

		newUser2.Follow(newUser1)
		newUser1.UploadPhoto()
		notify := newUser2.DisplayActivity()

		assert.Equal(t, expected, notify)
	})
}



// =====================================================

func TestFollow(t *testing.T) {
	t.Run("should return true when Alice follows Bob",func(t *testing.T) {
		user1, user2 := "Alice", "Bob"
		newUser1 := entity.NewUser(user1)
		newUser2 := entity.NewUser(user2)

		newUser1.Follow(newUser2)

		assert.True(t, newUser1.IsFollowed(newUser2))
	})
}
func TestLikedPhoto(t *testing.T) {
	t.Run("should return ErrYouDontHaveAPhoto when like my none photo", func(t *testing.T) {
		user1:= "Alice"
		newUser1 := entity.NewUser(user1)
		expected := apperror.ErrYouDontHaveAPhoto

		err := newUser1.LikedPhoto(newUser1)

		assert.ErrorIs(t, err, expected)
	})

	t.Run("should return ErrUserDoesntHaveAPhoto when like publisher none photo", func(t *testing.T) {
		user1, user2 := "Alice", "Bob"
		newUser1 := entity.NewUser(user1)
		newUser2 := entity.NewUser(user2)
		expected := apperror.ErrUserDoesntHaveAPhoto

		newUser1.Follow(newUser2)
		err := newUser1.LikedPhoto(newUser2)

		assert.ErrorIs(t, err, expected)
	})
	t.Run("should return ErrAlreadyLikedPhoto when like same publisher photo", func(t *testing.T) {
		user1, user2 := "Alice", "Bob"
		newUser1 := entity.NewUser(user1)
		newUser2 := entity.NewUser(user2)
		expected := apperror.ErrAlreadyLikedPhoto

		newUser2.UploadPhoto()
		newUser1.Follow(newUser2)
		newUser1.LikedPhoto(newUser2)
		err := newUser1.LikedPhoto(newUser2)

		assert.ErrorIs(t, err, expected)
	})

	t.Run("should return ErrLikePhotoUserNotFollowedYet when like publisher wasnt follow", func(t *testing.T) {
		user1, user2 := "Alice", "Bob"
		newUser1 := entity.NewUser(user1)
		newUser2 := entity.NewUser(user2)
		expected := apperror.ErrLikePhotoUserNotFollowedYet

		newUser2.UploadPhoto()
		err := newUser1.LikedPhoto(newUser2)

		assert.ErrorIs(t, err, expected)
	})

	t.Run("should notification you liked your photo, when like myself photo", func(t *testing.T) {
		user1:= "Alice"
		newUser1 := entity.NewUser(user1)
		expected := "\nAlice activities:\nYou uploaded photo\nYou liked your photo\n"
		newUser1.UploadPhoto()
		newUser1.LikedPhoto(newUser1)
		activity := newUser1.DisplayActivity()

		assert.Equal(t, expected, activity)
	})

	t.Run("should notification Alice liked your photo, when Alice like Bob photo", func(t *testing.T) {
		user1, user2:= "Alice", "Bob"
		newUser1 := entity.NewUser(user1)
		newUser2 := entity.NewUser(user2)
		expected := "\nBob activities:\nYou uploaded photo\nAlice liked your photo\n"

		newUser2.UploadPhoto()
		newUser1.Follow(newUser2)
		newUser1.LikedPhoto(newUser2)
		activity := newUser2.DisplayActivity()

		assert.Equal(t, expected, activity)
	})
	
	t.Run("should notification Alice liked Bob's photo, when Alice like Bob photo followed by Bill", func(t *testing.T) {
		user1, user2, user3:= "Alice", "Bob", "Bill"
		newUser1 := entity.NewUser(user1)
		newUser2 := entity.NewUser(user2)
		newUser3 := entity.NewUser(user3)
		expected := "\nBill activities:\nAlice liked Bob's photo\n"

		newUser2.UploadPhoto()
		newUser1.Follow(newUser2)
		newUser3.Follow(newUser1)
		newUser1.LikedPhoto(newUser2)
		activity := newUser3.DisplayActivity()

		assert.Equal(t, expected, activity)
	})

	// t.Run("jsisd", func(t *testing.T) {
	// 	alice, bob, bill, john := "Alice", "Bob", "Bill", "John"
	// 	Alice := entity.NewUser(alice)
	// 	Bob := entity.NewUser(bob)
	// 	Bill := entity.NewUser(bill)
	// 	John := entity.NewUser(john)

	// 	Alice.Follow(Bob)
	// 	Alice.Follow(Bill)
	// 	John.Follow(Bob)
	// 	Bob.Follow(Alice)
	// 	Bob.Follow(Bill)
	// 	John.Follow(Alice)

	// 	Alice.UploadPhoto()
	// 	Bob.LikedPhoto(Alice)
	// 	Bill.UploadPhoto()
	// 	Bob.LikedPhoto(Bill)
	// 	Bill.LikedPhoto(Bill)
	// 	Alice.LikedPhoto(Bill)

	// 	assert.Equal(t, Bill.DisplayActivity(), "fddf")

	// })
}