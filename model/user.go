package model

import (
	"fmt"
	"time"
)

type User struct {
	ID        string    `json:"id" db:"id"`
	Username  string    `json:"username" binding:"required" db:"username"`
	Token     string    `json:"token" db:"token"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// UserStore Interface
type UserStore interface {
	GetByUsername(username string) (User, error)
	Create(username, password string) (string, error)
	DeleteByUsername(username string) error
}

type userImpl struct{}

// NewUserStore
func NewUserStore() UserStore {
	return &userImpl{}
}

func (impl *userImpl) Create(username, passowrd string) (string, error) {
	prepString := GetCreateSQLPreString("user")
	fmt.Println("prepString", prepString)
	tx, err := DB.Beginx()
	if err != nil {
		tx.Rollback()
	}
	stmt, err := tx.Prepare(prepString)
	defer stmt.Close()
	token := RandString(16)
	id := GenID()
	timeNow := GetCurrentTimeStamp()
	_, err = stmt.Exec(id, username, token, passowrd, timeNow)
	fmt.Println("errrr", err)
	tx.Commit()
	return token, err
}

func (impl *userImpl) GetByUsername(username string) (User, error) {
	r := User{}
	err := DB.Get(&r, "SELECT * FROM user WHERE username=?", username)
	return r, err

}
func (impl *userImpl) DeleteByUsername(username string) error {
	_, err := DB.Exec("DELETE FROM user WHERE username = ?", username)
	return err
}
