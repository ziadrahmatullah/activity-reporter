package entity

type Follower interface{
	FollowerNotification(Publisher, Notification)
	FollowerNotificationAboutActivity(string)
	Follow(Publisher) error
	IsFollowed(Publisher) bool
	LikedPhoto(Publisher) error
	notifyActivityToAll(Publisher, Notification)
	UserName() string
}