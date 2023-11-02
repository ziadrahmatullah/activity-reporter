package entity

type Follower interface{
	FollowerNotify()
	Follow(Publisher)
	LikedPhoto(Publisher)
	notifyPublisher()
}