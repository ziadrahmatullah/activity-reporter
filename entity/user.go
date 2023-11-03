package entity

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"
)

type User struct {
	name          string
	followers     []Follower
	followings    []Publisher
	photoLikedBy  []Follower
	photoLike     []Publisher
	havePhoto     bool
	notifications []string
}

func NewUser(name string) *User {
	newUser := &User{name: name}
	return newUser
}

func (u *User) PublisherNotificationAboutLike(follower Follower, notification Notification) {
	u.photoLikedBy = append(u.photoLikedBy, follower)
	message := notification.Notify(u, follower)
	u.notifications = append(u.notifications, message)
}

func (u *User) FollowedBy(follower Follower) {
	u.followers = append(u.followers, follower)
}

func (u *User) UploadPhoto() (err error) {
	if !u.havePhoto {
		u.havePhoto = true
		u.notifyUploadToFollowers(&NotifyUpload{})
		u.FollowerNotification(u, &NotifyUpload{})
		return
	}
	return apperror.ErrCannotUploadMorePhoto
}

func (u *User) notifyUploadToFollowers(notification Notification) {
	for _, follower := range u.followers {
		follower.FollowerNotification(u, notification)
	}
}

func (u *User) notifyActivityToFollowers(message string, publisher Publisher) {
	for _, follower := range u.followers {
		if !(follower.UserName() == publisher.UserName()) {
			follower.FollowerNotificationAboutActivity(message)
		}
	}
}

func (u *User) UserName() string {
	return u.name
}

func (u *User) UserPhoto() bool {
	return u.havePhoto
}

func (u *User) ShowLikes() int {
	return len(u.photoLikedBy)
}

func (u *User) FollowerNotification(publisher Publisher, notification Notification) {
	message := notification.Notify(publisher, u)
	u.notifications = append(u.notifications, message)
}

func (u *User) FollowerNotificationAboutActivity(message string) {
	u.notifications = append(u.notifications, message)
}

func (u *User) Follow(publisher Publisher) (err error) {
	if u == publisher {
		return apperror.ErrCantFollowThemselves
	}
	if u.IsFollowed(publisher) {
		return apperror.ErrAlreadyFollowUser
	}
	u.followings = append(u.followings, publisher)
	publisher.FollowedBy(u)
	return
}

func (u *User) LikedPhoto(publisher Publisher) (err error) {
	if !publisher.UserPhoto() {
		if u == publisher {
			return apperror.ErrYouDontHaveAPhoto
		}
		return apperror.ErrUserDoesntHaveAPhoto
	}
	if u.isPhotoAlreadyLiked(publisher) {
		return apperror.ErrAlreadyLikedPhoto
	}
	if u == publisher {
		u.photoLike = append(u.photoLike, publisher)
		u.notifyActivityToAll(publisher, &NotifyLike{})
		return
	} else if u.IsFollowed(publisher) {
		u.photoLike = append(u.photoLike, publisher)
		u.notifyActivityToAll(publisher, &NotifyLike{})
		u.FollowerNotification(publisher, &NotifyMySelfAboutLike{})
		return
	}
	return apperror.ErrLikePhotoUserNotFollowedYet
}

func (u *User) notifyActivityToAll(publisher Publisher, notification Notification) {
	publisher.PublisherNotificationAboutLike(u, notification)
	tempNotification := &NotifyActivityToFollower{}
	message := tempNotification.Notify(publisher, u)
	u.notifyActivityToFollowers(message, publisher)
}

func (u *User) IsFollowed(publisher Publisher) bool {
	for _, following := range u.followings {
		if following == publisher {
			return true
		}
	}
	return false
}

func (u *User) isPhotoAlreadyLiked(publisher Publisher) bool {
	for _, user := range u.photoLike {
		if user.UserName() == publisher.UserName() {
			return true
		}
	}
	return false
}

func (u *User) IsHigherLikeThan(user *User) bool {
	return u.ShowLikes() > user.ShowLikes()
}

func (u *User) DisplayActivity() (output string) {
	output += fmt.Sprintf("\n%s activities:\n", u.UserName())
	for _, notification := range u.notifications {
		output += notification + "\n"
	}
	return
}
