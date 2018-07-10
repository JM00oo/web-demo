package module

import "github.com/web-demo/model"

func CreatePost(title, content, userID string) error {
	postStore := model.NewPostStore()
	err := postStore.Create(title, content, userID)
	return err
}

func GetPost() []model.Post {
	postStore := model.NewPostStore()
	posts := postStore.GetPost()
	return posts
}
