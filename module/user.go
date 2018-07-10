package module

import (
	"errors"

	"github.com/web-demo/model"
)

func getPAToken(paID string) string {
	// FIXME use JWT
	return model.RandString(32)
}
func Signup(username, password string) (string, error) {
	userStore := model.NewUserStore()
	_, err := userStore.GetByUsername(username)
	if err != nil {
		token, err := userStore.Create(username, password)
		return token, err
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
