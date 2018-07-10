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
	Logout(token string)
	Create(username, password string) (string, error)
}

type userImpl struct{}

// NewUserStore
func NewUserStore() UserStore {
	return &userImpl{}
}
func (impl *userImpl) Logout(token string) {
	// TODO:
	return
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
	// TODO
	r := User{}
	return r, nil

}
