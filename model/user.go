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
	Create(username, password string) (User, error)
	DeleteByUsername(username string) error
	GetByToken(token string) (User, error)
	GetByPasswordAndUsername(pw, username string) (User, error)
}

type userImpl struct{}

// NewUserStore
func NewUserStore() UserStore {
	return &userImpl{}
}

func (impl *userImpl) Create(username, passowrd string) (User, error) {
	prepString := GetCreateSQLPreString("user")
	fmt.Println("prepString", prepString)
	tx, err := DB.Beginx()
	if err != nil {
		fmt.Println("err in create user", err)
		tx.Rollback()
	}
	stmt, err := tx.Prepare(prepString)
	defer stmt.Close()
	token := RandString(16)
	id := GenID()
	timeNow := GetCurrentTimeStamp()
	_, err = stmt.Exec(id, username, token, passowrd, timeNow)
	tx.Commit()

	user, err := impl.GetByUsername(username)
	return user, err
}

func (impl *userImpl) GetByPasswordAndUsername(pw, username string) (User, error) {
	r := User{}
	err := DB.Get(&r, "SELECT * FROM user WHERE password=? AND username=?", pw, username)
	return r, err

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

func (impl *userImpl) GetByToken(token string) (User, error) {
	r := User{}
	err := DB.Get(&r, "SELECT * FROM user WHERE token=?", token)
	return r, err

}
