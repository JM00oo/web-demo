package module

import "github.com/web-demo/model"

func CreateComment(comment, postID, ownerName string) error {
	commentStore := model.NewCommentStore()
	err := commentStore.Create(comment, postID, ownerName)
	return err
}
