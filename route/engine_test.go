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
	model.DBInit()
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
	testUsername := "test-username-2" + model.GetCurrentTimeStampUnixTime()
	reqSignup := `{"username": "` + testUsername + `", "password":"123456"}`
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

func TestPostAndComment(t *testing.T) {
	TEST_DATA := beforeWithExistedUser()
	defer afterWithExistedUser(TEST_DATA)
	router := GetMainEngine()
	user := TEST_DATA["user"].(model.User)
	// Create post
	reqString := `{"title":"this is title", "content":"this is content"}`
	bodyByte := []byte(reqString)

	req, err := http.NewRequest("POST", "/api/post", bytes.NewBuffer(bodyByte))
	req.Header.Set("token", user.Token)
	if err != nil {
		fmt.Println("err in test signup: ", err)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Println("w.Body", w.Body)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, err, nil)
	var resp gin.H
	err = json.Unmarshal([]byte(w.Body.String()), &resp)
	post := resp["result"].(map[string]interface{})
	postID := post["id"].(string)

	reqString = `{"comment":"this is coment"}`
	bodyByte = []byte(reqString)

	req, err = http.NewRequest("POST", "/api/comment?postID="+postID, bytes.NewBuffer(bodyByte))
	req.Header.Set("token", user.Token)
	if err != nil {
		fmt.Println("err in test post post: ", err)
	}
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Println("comment", w.Body)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, err, nil)

	req, err = http.NewRequest("GET", "/api/post", nil)
	req.Header.Set("token", user.Token)
	if err != nil {
		fmt.Println("err in test GET post: ", err)
	}
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, err, nil)
	err = json.Unmarshal([]byte(w.Body.String()), &resp)
	posts := resp["posts"].([]interface{})

	commentList := posts[0].(map[string]interface{})["comment"].([]interface{})
	assert.Equal(t, 1, len(commentList))
	newPostID := posts[0].(map[string]interface{})["id"].(string)
	assert.Equal(t, postID, newPostID)

}
