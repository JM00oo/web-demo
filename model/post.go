package model

import (
	"fmt"
	"time"
)

type Post struct {
	ID        string    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" binding:"required" db:"content"`
	OwnerID   string    `json:"ownerID" db:"owner_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// func CreatePost(title, content, userID string) error {
// 	postStore := model.NewPostStore()
// 	err := postStore.Create(title, content, userID)
// 	return err
// }
// func GetPost() []model.Post {
// 	postStore := model.NewPostStore()
// 	posts := postStore.GetPost()
// 	return posts
// }

// PostStore Interface
type PostStore interface {
	GetPost() []Post
	Create(title, content, userID string) error
}

type postImpl struct{}

// NewPostStore
func NewPostStore() PostStore {
	return &postImpl{}
}

func (impl *postImpl) Create(title, content, userID string) error {
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
	_, err = stmt.Exec(id, title, content, userID, timeNow)
	fmt.Println("errrr", err)
	tx.Commit()
	return err
}

func (impl *postImpl) GetPost() []Post {
	r := []Post{}
	DB.Select(&r, "SELECT * FROM post")
	return r
}

//
// func (impl *userImpl) GetByUsername(username string) (User, error) {
// 	r := User{}
// 	err := DB.Get(&r, "SELECT * FROM user WHERE username=?", username)
// 	return r, err
//
// }
// func (impl *userImpl) DeleteByUsername(username string) error {
// 	_, err := DB.Exec("DELETE FROM user WHERE username = ?", username)
// 	return err
// }
