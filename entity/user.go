package entity

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/apperror"
)

type User struct {
	name            string
	followers       []Follower
	followings      []Publisher
	whoLikedMyPhoto []Follower
	whoPhotoILiked  []Publisher
	photo           bool
	notifications   []string
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) PublisherNotify(follower Follower, notification Notification) {
	message := notification.Notify(u, follower)
	u.notifications = append(u.notifications, message)
}

func (u *User) FollowedBy(follower Follower) {
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

func (u *User) AddLiker(follower Follower) {
	u.whoLikedMyPhoto = append(u.whoLikedMyPhoto, follower)
}

func (u *User) UserName() string {
	return u.name
}

func (u *User) UserPhoto() bool {
	return u.photo
}

func (u *User) ShowLikes() int {
	return len(u.whoLikedMyPhoto)
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
	if !publisher.UserPhoto() {
		if u == publisher {
			return apperror.ErrYouDontHaveAPhoto
		}
		return apperror.ErrUserDoesntHaveAPhoto
	}
	if u.isPhotoAlreadyLiked(publisher) {
		return apperror.ErrAlreadyLikedPhoto
	}
	if u.isFollowed(publisher) || u == publisher {
		u.whoPhotoILiked = append(u.whoPhotoILiked, publisher)
		publisher.AddLiker(u)
		u.notifyPublisher(publisher, &NotifyLike{})
		return
	}
	return apperror.ErrLikePhotoUserNotFollowedYet
}

func (u *User) notifyPublisher(publisher Publisher, notification Notification) {
	publisher.PublisherNotify(u, notification)
	u.notifyFollowers(&NotfyLikeToFollower{})
}

func (u *User) isFollowed(publisher Publisher) bool {
	for _, following := range u.followings {
		if following == publisher {
			return true
		}
	}
	return false
}

func (u *User) isPhotoAlreadyLiked(publisher Publisher) bool {
	for _, user := range u.whoPhotoILiked {
		if user.UserName() == publisher.UserName() {
			return true
		}
	}
	return false
}

func (u *User) isHigherLikeThan(user *User) bool {
	return u.ShowLikes() > user.ShowLikes()
}
