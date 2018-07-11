package module

import (
	"errors"

	"github.com/web-demo/model"
)

func Signup(username, password string) (string, error) {
	userStore := model.NewUserStore()
	_, err := userStore.GetByUsername(username)
	if err != nil {
		user, err := userStore.Create(username, password)
		return user.Token, err
	} else {
		return "", errors.New("Duplicted")
	}
}

func Login(username, password string) (string, error) {
	userStore := model.NewUserStore()
	user, err := userStore.GetByUsername(username)
	if err == nil {
		return user.Token, nil
	} else {
		return "", errors.New("Not found")
	}

}

// func Logout(token string) {
// 	userStore := model.NewUserStore()
// 	userStore.Logout(token)
// }
