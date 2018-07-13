package main

import (
	"fmt"

	"github.com/web-demo/model"
	"github.com/web-demo/route"
)

// BuildDate  Build date
var BuildDate string

// GitSummary Git summary for current build
var GitSummary string

// Version Build version
var Version string

func main() {
	fmt.Println("gooooo")
	model.DBInit()
	fmt.Println("gooooo")
	fmt.Printf("Version: %s, Build Date: %s, Git Summary: %s\n", Version, BuildDate, GitSummary)
	r := route.GetMainEngine()
	r.GET("/", route.MainPage)
	r.Run(":8080")
}
