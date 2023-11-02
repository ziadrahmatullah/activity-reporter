package entity

import (
	"fmt"
	"sort"
)

type SocialApp struct{
	users []*User
}

func NewSocialApp() *SocialApp{
	return &SocialApp{}
}

func (sa *SocialApp) AddUser(user *User){
	sa.users = append(sa.users, user)
}

func (sa *SocialApp) Trending() string{
	sa.sortByTrending()
	title := "Trending photos:\n"
	i := 0
	for _, user := range sa.users{
		if i == 3{
			break
		}
		if !user.UserPhoto(){
			continue
		}
		count := "like"
		if sa.users[i].ShowLikes() > 1{
			count = "likes"
		}
		title += fmt.Sprintf("%d. %s photo got %d %s\n", i+1, user.UserName(), user.ShowLikes(), count )
		i++
	}
	return title
}
func (sa *SocialApp) sortByTrending(){
	sort.Slice(sa.users, func(i, j int) bool {
		return sa.users[i].isHigherLikeThan(sa.users[j])
	})
}

func (sa *SocialApp) IsUserInApp(newUser string) (eq bool, user *User){
	for _, user := range sa.users {
		if user.UserName() == newUser{
			return true, user
		}
	}
	return
}
