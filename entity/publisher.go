package entity

type Publisher interface{
	PublisherNotificationAboutLike(Follower, Notification)
	FollowedBy(Follower)
	UploadPhoto() error
	addLiker(Follower)
	notifyUploadToFollowers(Notification)
	notifyActivityToFollowers(string, Publisher)
	UserName() string
	UserPhoto() bool
}