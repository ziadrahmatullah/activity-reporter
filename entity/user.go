package entity

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"

type User struct {
	name          string
	followers     []Follower
	followings    []Publisher
	whoLikedPhoto []Follower
	whoPhotoLiked []Publisher
	photo         bool
	notification  []Notification
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) UploadFoto() (err error) {
	if u.photo {
		u.photo = true
		return
	}
	return apperror.ErrCannotUploadMorePhoto
}

func (u *User) Folllow(publisher Publisher) (err error) {
	// if u == publisher{
	// return apperror.ErrCantFollowThemselves
	// }
	if !u.isFollowed(publisher) {
		u.followings = append(u.followings, publisher)
		return
	}
	return apperror.ErrAlreadyFollowUser
}

func (u *User) LikedPhoto(publisher Publisher) {
	u.whoPhotoLiked = append(u.whoPhotoLiked, publisher)
}

func (u *User) isFollowed(publisher Publisher) bool {
	for _, following := range u.followings {
		if following == publisher {
			return true
		}
	}
	return false
}
