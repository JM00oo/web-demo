package model

import (
	"fmt"
	"time"
)

type Post struct {
	Comment   []Comment `json:"comment,omitempty" db:"-" `
	ID        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" binding:"required" db:"content"`
	OwnerName string    `json:"ownerName" db:"owner_name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// PostStore Interface
type PostStore interface {
	GetPost() []Post
	Create(title, content, userID string) (Post, error)
}

type postImpl struct{}

// NewPostStore
func NewPostStore() PostStore {
	return &postImpl{}
}

func (impl *postImpl) Create(title, content, userName string) (Post, error) {
	prepString := GetCreateSQLPreString("post")
	fmt.Println("prepString", prepString)
	tx, err := DB.Beginx()
	if err != nil {
		tx.Rollback()
	}
	stmt, err := tx.Prepare(prepString)
	defer stmt.Close()
	id := GenID()
	timeNow := GetCurrentTimeStamp()
	_, err = stmt.Exec(id, title, content, userName, timeNow)
	tx.Commit()
	post := impl.GetByID(id)
	return post, err
}
func (impl *postImpl) GetByID(postID string) Post {

	r := Post{}
	DB.Get(&r, "SELECT * FROM post WHERE id=?", postID)
	return r

}

func (impl *postImpl) GetPost() []Post {
	r := []Post{}
	DB.Select(&r, "SELECT * FROM post")
	cStore := NewCommentStore()
	for i := 0; i < len(r); i++ {
		r[i].Comment = cStore.GetComment(r[i].ID)
		fmt.Print("coment", r[i].Comment)
	}
	return r
}
