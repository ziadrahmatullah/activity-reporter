package entity

type Follower interface{
	FollowerNotify(Publisher, Notification)
	Follow(Publisher) error
	LikedPhoto(Publisher) error
	notifyPublisher(Publisher, Notification)
	userName() string
}