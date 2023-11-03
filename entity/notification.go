package entity

import "fmt"

type Notification interface {
	Notify(Publisher, Follower) string
}

type NotifyUpload struct{}

func (nu *NotifyUpload) Notify(publisher Publisher, follower Follower) string {
	if publisher.UserName() == follower.UserName(){
		return "You uploaded photo"
	}
	return fmt.Sprintf("%s uploaded photo", publisher.UserName())
}

type NotifyLike struct{}

func (nl *NotifyLike) Notify(publisher Publisher, follower Follower) string {
	if publisher.UserName() == follower.UserName(){		
		return "You liked your photo"
	}
	return fmt.Sprintf("%s liked your photo", follower.UserName())
}

type NotifyActivityToFollower struct{}

func (naf *NotifyActivityToFollower) Notify(publisher Publisher, follower Follower) string{
	return fmt.Sprintf("%s liked %s's photo", follower.UserName(), publisher.UserName())
}

type NotifyMySelfAboutLike struct{}

func (nms *NotifyMySelfAboutLike) Notify(publisher Publisher, follower Follower) string{
	return fmt.Sprintf("You liked %s's photo", publisher.UserName())
}