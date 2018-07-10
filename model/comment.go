package model

import "time"

type Comment struct {
	ID        string    `json:"id" db:"id"`
	Comment   string    `json:"comment" binding:"required" db:"comment"`
	PostID    string    `json:"postID" db:"post_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// CommentStore Interface
type CommentStore interface {
	Create(string, string) error
	GetComment() []Comment
}

type commentImpl struct{}

// NewCommentStore
func NewCommentStore() CommentStore {
	return &commentImpl{}
}

func (impl *commentImpl) Create(content, postID string) error {
	// marker := Marker{}
	// err := DB.Get(&marker, "SELECT * FROM marker WHERE id = ? AND deleted_at IS NULL", ID)
	// return marker, err
	return nil
}

func (impl *commentImpl) GetComment() []Comment {

	cList := []Comment{}
	// DB.Select(&cList, "SELECT * FROM comment")
	return cList

}
