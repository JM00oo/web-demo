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

func beforeTestPlatformAccount() error {
	model.DBInit()
	return nil
}

func afterTestPlatformAccount(t *testing.T) {
	defer model.DB.Close()
}

// func TestSignupURL(t *testing.T) {
// 	defer afterTestPlatformAccount(t)
// 	err := beforeTestPlatformAccount()
//
// 	router := GetMainEngine()
// 	username := "test-username" + model.GetCurrentTimeStampUnixTime()
// 	password := "000000"signupResult
// 	reqURL := "/api/signup"
// 	data := map[string]string{
// 		"email":               email,
// 		"password":            password,
// 		"panoCollectionLimit": "1",
// 	}
// 	body := model.GetFormData(data)
// 	req, err := http.NewRequest("POST", reqURL, body)
// 	req.Header.Add("istaging-api-key", iStagingAPIKey)
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	if err != nil {
// 		fmt.Println("err in testPlatformAccountSignup: ", err)
// 	}
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	fmt.Println(w.Body)
// 	var signupResult gin.H
// 	err = json.Unmarshal([]byte(w.Body.String()), &signupResult)
// 	if err != nil {
// 		fmt.Println("err in testPlatformAccountSignup unmarshall: ", err)
// 	}
// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	assert.Equal(t, err, nil)
// 	assert.NotNil(t, signupResult["id"].(string))
// 	assert.NotNil(t, signupResult["platformAccountToken"].(string))
//
// 	paStore := model.NewPlatformAccountStore()
// 	err = paStore.DeleteByID(signupResult["id"].(string), true)
// 	assert.Equal(t, err, nil)
// }

func TestSignup(t *testing.T) {
	defer afterTestPlatformAccount(t)
	err := beforeTestPlatformAccount()
	router := GetMainEngine()
	testUsername := "test-username"
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
// func TestPASignupNoPassword(t *testing.T) {
// 	defer afterTestPlatformAccount(t)
// 	err := beforeTestPlatformAccount()
// 	router := GetMainEngine()
// 	reqSignup := `{"email": "tenant_test_name"}`
// 	bodyByte := []byte(reqSignup)
//
// 	req, err := http.NewRequest("POST", "/api/v1/platformAccount/signup", bytes.NewBuffer(bodyByte))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("istaging-api-key", iStagingAPIKey)
// 	if err != nil {
// 		fmt.Println("err in testPASignupNoPassword: ", err)
// 	}
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	fmt.Println(w.Body)
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }
//
// func TestPASignupNoEmail(t *testing.T) {
// 	defer afterTestPlatformAccount(t)
// 	err := beforeTestPlatformAccount()
// 	router := GetMainEngine()
// 	reqSignup := `{"password": "123456"}`
// 	bodyByte := []byte(reqSignup)
//
// 	req, err := http.NewRequest("POST", "/api/v1/platformAccount/signup", bytes.NewBuffer(bodyByte))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("istaging-api-key", iStagingAPIKey)
// 	if err != nil {
// 		fmt.Println("err in testPASignupNoEmail: ", err)
// 	}
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	fmt.Println(w.Body.String())
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }
//
// func TestPASignupNoPanoLimit(t *testing.T) {
// 	defer afterTestPlatformAccount(t)
// 	err := beforeTestPlatformAccount()
// 	router := GetMainEngine()
// 	reqSignup := `{"password": "123456", "email": "tenant_test_name"}`
// 	bodyByte := []byte(reqSignup)
//
// 	req, err := http.NewRequest("POST", "/api/v1/platformAccount/signup", bytes.NewBuffer(bodyByte))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("istaging-api-key", iStagingAPIKey)
// 	if err != nil {
// 		fmt.Println("err in testPASignupNoEmail: ", err)
// 	}
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
// 	var resp map[string]interface{}
// 	json.Unmarshal([]byte(w.Body.String()), &resp)
//
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// 	assert.Equal(t, "invalid format. no panoCollectionLimit.", resp["errorMsg"].(string))
//
// }
//
// func TestPASignupNoAPIKey(t *testing.T) {
// 	defer afterTestPlatformAccount(t)
// 	err := beforeTestPlatformAccount()
// 	router := GetMainEngine()
// 	reqSignup := model.PASignupEntry{"t-test-tenant-signup", "000000", 1}
// 	marshalBody, _ := json.Marshal(reqSignup)
// 	bodyByte := []byte(string(marshalBody))
// 	req, err := http.NewRequest("POST", "/api/v1/platformAccount/signup", bytes.NewBuffer(bodyByte))
// 	req.Header.Add("Content-Type", "application/json")
// 	if err != nil {
// 		fmt.Println("err in TestSignupNoAPIKey: ", err)
// 	}
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	fmt.Println(w.Body)
// 	assert.Equal(t, http.StatusUnauthorized, w.Code)
//
// }
//
// func TestPASignupWrongAPIKey(t *testing.T) {
// 	defer afterTestPlatformAccount(t)
// 	err := beforeTestPlatformAccount()
// 	router := GetMainEngine()
// 	reqSignup := model.PASignupEntry{"t-test-tenant-signup", "000000", 1}
// 	marshalBody, _ := json.Marshal(reqSignup)
// 	bodyByte := []byte(string(marshalBody))
// 	req, err := http.NewRequest("POST", "/api/v1/platformAccount/signup", bytes.NewBuffer(bodyByte))
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("istaging-api-key", "NotExistKey")
//
// 	if err != nil {
// 		fmt.Println("err in TestSignupNoAPIKey: ", err)
// 	}
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
//
// 	fmt.Println(w.Body)
// 	assert.Equal(t, http.StatusUnauthorized, w.Code)
//
// }
