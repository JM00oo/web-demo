package model

import (
	"fmt"
	"time"
)

type Comment struct {
	ID        string    `json:"id" db:"id"`
	Comment   string    `json:"comment" binding:"required" db:"comment"`
	PostID    string    `json:"postID" db:"post_id"`
	OwnerName string    `json:"ownerName" db:"owner_name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// CommentStore Interface
type CommentStore interface {
	Create(content, postID, ownerName string) error
	GetComment(string) []Comment
}

type commentImpl struct{}

// NewCommentStore
func NewCommentStore() CommentStore {
	return &commentImpl{}
}

func (impl *commentImpl) Create(content, postID, ownerName string) error {
	prepString := GetCreateSQLPreString("comment")
	fmt.Println("prepString", prepString)
	tx, err := DB.Beginx()
	if err != nil {
		tx.Rollback()
	}
	stmt, err := tx.Prepare(prepString)
	defer stmt.Close()
	id := GenID()
	timeNow := GetCurrentTimeStamp()
	_, err = stmt.Exec(id, content, postID, ownerName, timeNow)
	tx.Commit()
	return err
}

func (impl *commentImpl) GetComment(postID string) []Comment {

	cList := []Comment{}
	DB.Select(&cList, "SELECT * FROM comment WHERE post_id =?", postID)
	return cList

}
