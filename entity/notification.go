package entity

import "fmt"

type Notification interface {
	Notify(Publisher, Follower) string
}

type NotifyUpload struct{}

func (nu *NotifyUpload) Notify(publisher Publisher, follower Follower) string {
	return fmt.Sprintf("%s uploaded photo", publisher.userName())
}

type NotyfyLiked struct{}

//==============================================

type NotifyUploaded struct{}

type NotifyLike struct{}

func (nl *NotifyLike) Notify(publisher Publisher, follower Follower) string {
	return fmt.Sprintf("%s liked %s's photo", follower.userName(), publisher.userName())
}
