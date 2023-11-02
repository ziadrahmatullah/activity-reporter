package notify

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"

type NotifyFollowed struct{}

type NotifyUpload struct{}

type NotyfyLiked struct{}

type NotifyFollow struct{}

func (nf *NotifyFollow) NotifyFollower(user *entity.User){
	
}

type NotifyUploaded struct{}

type NotifyLike struct{}

