package entity

type Publisher interface{
	PublisherNotify(Follower, Notification)
	FollowedBy(Follower)
	UploadFoto() error
	notifyFollowers(Notification)
	userName() string
	userPhoto() bool
}