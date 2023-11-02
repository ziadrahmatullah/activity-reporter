package apperror

import "errors"

var(
	ErrCannotUploadMorePhoto = errors.New("you cannot upload more than once")
	ErrAlreadyFollowUser = errors.New("you already followed the user")
	ErrAlreadyLikedPhoto = errors.New("you already liked the photo")
	ErrUserDoesntHaveAPhoto = errors.New("inputed user doesnt have a photo")
	ErrYouDontHaveAPhoto = errors.New("you don't have a photo")
	ErrLikePhotoUserNotFollowedYet = errors.New("unable to like user's photo")
	ErrCantFollowThemselves = errors.New("a user cannot follow themselves")
	ErrInvalidKeyword = errors.New("invalid keyword")
)