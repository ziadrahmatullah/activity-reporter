package entity

type Publisher interface{
	PublisherNotify(Follower, Notification)
	FollowedBy(Follower)
	IsFollowed(Publisher) bool
	UploadPhoto() error
	AddLiker(Follower)
	notifyFollowers(Notification)
	notifyActivityToFollower(string, Publisher)
	UserName() string
	UserPhoto() bool
}