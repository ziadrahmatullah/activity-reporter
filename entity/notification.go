package entity

import "fmt"

type Notification interface {
	Notify(Publisher, Follower) string
}

type NotifyUpload struct{}

func (nu *NotifyUpload) Notify(publisher Publisher, follower Follower) string {
	if publisher.UserName() == follower.UserName(){
		return "you uploaded photo"
	}
	return fmt.Sprintf("%s uploaded photo", publisher.UserName())
}

type NotifyLike struct{}

func (nl *NotifyLike) Notify(publisher Publisher, follower Follower) string {
	if publisher.UserName() == follower.UserName(){		
		return "you liked your photo"
	}
	return fmt.Sprintf("%s liked your photo", follower.UserName())
}


//==============================================

// type NotifyUploaded struct{}

type NotifyActivityToFollower struct{}

func (naf *NotifyActivityToFollower) Notify(publisher Publisher, follower Follower) string{
	return fmt.Sprintf("%s liked %s's photo", follower.UserName(), publisher.UserName())
}