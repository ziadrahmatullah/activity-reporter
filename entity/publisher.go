package entity

type Publisher interface{
	PublisherNotify(Follower, Notification)
	FollowedBy(Follower)
	UploadFoto() error
	AddLiker(Follower)
	notifyFollowers(Notification)
	UserName() string
	UserPhoto() bool
}