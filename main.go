package main

import (
	postApp "github.com/madhukar32/post-webapp/pkg/posts"
)

func main() {
	postApp.CreateRouter(8000)
}
