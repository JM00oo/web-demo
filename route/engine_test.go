package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/web-demo/model"
)

func beforeTest() error {
	model.DBInit()
	return nil
}

func afterTest() {
	defer model.DB.Close()
}

func beforeWithExistedUser() map[string]interface{} {
	TEST_DATA := make(map[string]interface{})
	userStore := model.NewUserStore()
	user, _ := userStore.Create("postTestUser", "000000")
	TEST_DATA["user"] = user
	return TEST_DATA
}
func afterWithExistedUser(TEST_DATA map[string]interface{}) {
	defer model.DB.Close()

	userStore := model.NewUserStore()
	user := TEST_DATA["user"].(model.User)
	userStore.DeleteByUsername(user.Username)
}

func TestSignup(t *testing.T) {
	defer afterTest()
	err := beforeTest()
	router := GetMainEngine()
	testUsername := "test-username-2"
	reqSignup := `{"username": "` + testUsername + `", "password":"123456"}`
	fmt.Println("reqSignup", reqSignup)
	bodyByte := []byte(reqSignup)

	req, err := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(bodyByte))
	if err != nil {
		fmt.Println("err in test signup: ", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Println(w.Body)
	var signupResult gin.H
	err = json.Unmarshal([]byte(w.Body.String()), &signupResult)
	if err != nil {
		fmt.Println("err in test signup unmarshall: ", err)
	}
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, err, nil)
	fmt.Println("signupResult", signupResult)
	assert.NotNil(t, signupResult["token"].(string))
	store := model.NewUserStore()
	user, err := store.GetByUsername(testUsername)
	assert.Equal(t, testUsername, user.Username)
	assert.Equal(t, err, nil)
	err = store.DeleteByUsername(testUsername)
	assert.Equal(t, err, nil)

}

//
// func TestPostAndComment(t *testing.T) {
// 	defer afterWithExistedUser()
// 	TEST_DATA := beforeWithExistedUser()
// 	router := GetMainEngine()
// 	{
// 	"title":"this is title",
// 	"content":"this is content"
// }
// 	req := `{"title":"this is title", "content":"this is content"}`
// 	bodyByte := []byte(req)
//
// 	req, err := http.NewRequest("POST", "/api/post", bytes.NewBuffer(bodyByte))
// 	if err != nil {
// 		fmt.Println("err in test signup: ", err)
// 	}
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	fmt.Println(w.Body)
// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	assert.Equal(t, err, nil)
//
//
// }
