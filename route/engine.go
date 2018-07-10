package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/web-demo/module"
)

// GetMainEngine api service router engine
func GetMainEngine() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("template/*.html")
	api := r.Group("/api")
	api.POST("/signup", Signup)
	api.POST("/login", Login)

	// apiAuth := r.Group("/api")
	// apiAuth.Use(middleware.LoginRequired())
	// apiAuth.POST("/logout", Logout)
	// apiAuth.POST("/post", CreatePost)
	// apiAuth.POST("/comment", CreateComment)
	// apiAuth.GET("/comment", GetComment)
	// apiAuth.GET("/post", GetPost)

	return r
}
func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// SignupEntry : Signup entry struct
type SignupEntry struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type PostEntry struct {
	Title   string `json:title`
	Content string `json:content`
}

type CommentEntry struct {
	Comment string `json:comment`
}

func Signup(c *gin.Context) {
	var req SignupEntry
	var err error
	err = c.ShouldBindWith(&req, binding.JSON)

	if err != nil {
		fmt.Println("signup err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}
	username := req.Username
	passwdStr := req.Password

	token, err := module.Signup(username, passwdStr)
	if err != nil {
		fmt.Println("signup err", err)
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func Login(c *gin.Context) {
	var req SignupEntry
	var err error
	err = c.ShouldBindWith(&req, binding.JSON)

	if err != nil {
		fmt.Println("login err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}
	emailStr := req.Username
	passwdStr := req.Password

	token, err := module.Login(emailStr, passwdStr)
	if err != nil {
		fmt.Println("login err", err)
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func Logout(c *gin.Context) {
	token := c.Request.Header.Get("token")
	module.Logout(token)
	c.JSON(http.StatusOK, nil)
}

/*
func CreatePost(c *gin.Context) {
	var req PostEntry
	var err error
	err = c.ShouldBindWith(&req, binding.JSON)

	if err != nil {
		fmt.Println("create post err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}
	title := req.Title
	content := req.Content
	userID := c.MustGet("userID")
	err = module.CreatePost(title, content, userID)
	if err != nil {
		fmt.Println("create post err", err)
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, nil)
}
func GetPost(c *gin.Context) {
	posts := module.GetPost()
	c.JSON(http.StatusOK, posts)
}

func CreateComment(c *gin.Context) {
	var req CommentEntry
	var err error
	err = c.ShouldBindWith(&req, binding.JSON)

	if err != nil {
		fmt.Println("create comment err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}
	comment := req.Comment
	postID := c.Query("postID")
	err := module.CreateComment(comment, postID)
	if err != nil {
		fmt.Println("create comment err", err)
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func GetComment(c *gin.Context) {
	comments := module.GetComment()
	c.JSON(http.StatusOK, comments)
}
*/
