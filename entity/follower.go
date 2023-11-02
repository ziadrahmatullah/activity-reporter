package entity

type Follower interface{
	FollowerNotify(Publisher, Notification)
	FollowerActivityNotify(string)
	Follow(Publisher) error
	LikedPhoto(Publisher) error
	notifyPublisher(Publisher, Notification)
	UserName() string
}