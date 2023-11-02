package entity

type Publisher interface{
	PublisherNotify()
	UploadFoto()
	notifyFollower()
}