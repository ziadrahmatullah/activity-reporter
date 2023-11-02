package entity

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"
)

type User struct {
	name          string
	followers     []Follower
	followings    []Publisher
	whoLikedPhoto []Follower
	whoPhotoLiked []Publisher
	photo         bool
	notifications []string
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) PublisherNotify(follower Follower, notification Notification) {
	message := notification.Notify(u, follower)
	u.notifications = append(u.notifications, message)
}

func (u *User) FollowedBy(follower Follower){
	u.followers = append(u.followers, follower)
}

func (u *User) UploadFoto() (err error) {
	if !u.photo {
		u.photo = true
		u.notifyFollowers(&NotifyUpload{})
		return
	}
	return apperror.ErrCannotUploadMorePhoto
}

func (u *User) notifyFollowers(notification Notification) {
	for _, follower := range u.followers {
		follower.FollowerNotify(u, notification)
	}
}

func (u *User) userName() string {
	return u.name
}

func (u *User) userPhoto() bool {
	return u.photo
}

//============================================================
func (u *User) FollowerNotify(publisher Publisher, notification Notification) {
	message := notification.Notify(publisher, u)
	u.notifications = append(u.notifications, message)
}

func (u *User) Follow(publisher Publisher) (err error) {
	if u == publisher {
		return apperror.ErrCantFollowThemselves
	}
	if u.isFollowed(publisher) {
		return apperror.ErrAlreadyFollowUser
	}
	u.followings = append(u.followings, publisher)
	publisher.FollowedBy(u)
	return
}

func (u *User) LikedPhoto(publisher Publisher) (err error) {
	if !publisher.userPhoto() {
		if u == publisher {
			return apperror.ErrYouDontHaveAPhoto
		}
		return apperror.ErrUserDoesntHaveAPhoto
	}
	if u == publisher{
		u.whoLikedPhoto = append(u.whoLikedPhoto, u)
		u.notifyPublisher(publisher, &NotifyLike{})	
		return
	} else if u.isFollowed(publisher){
		u.whoPhotoLiked = append(u.whoPhotoLiked, publisher)
		u.notifyPublisher(publisher, &NotifyLike{})
		return
	}	
	return apperror.ErrLikePhotoUserNotFollowedYet
}

func (u *User) notifyPublisher(publisher Publisher, notification Notification) {
	publisher.PublisherNotify(u, notification)
	u.notifyFollowers(notification)
}

func (u *User) isFollowed(publisher Publisher) bool {
	for _, following := range u.followings {
		if following == publisher {
			return true
		}
	}
	return false
}
