package module

import (
	"fmt"

	"github.com/web-demo/model"
)

func CreatePost(title, content, userName string) (model.Post, error) {
	fmt.Println("title, content, userName", title, content, userName)
	postStore := model.NewPostStore()
	post, err := postStore.Create(title, content, userName)
	return post, err
}

func GetPost() []model.Post {
	postStore := model.NewPostStore()
	posts := postStore.GetPost()
	return posts
}
