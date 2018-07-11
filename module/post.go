package module

import "github.com/web-demo/model"

func CreatePost(title, content, userName string) (model.Post, error) {
	postStore := model.NewPostStore()
	post, err := postStore.Create(title, content, userName)
	return post, err
}

func GetPost() []model.Post {
	postStore := model.NewPostStore()
	posts := postStore.GetPost()
	return posts
}
