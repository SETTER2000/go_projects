package main

import (
	"appal/news"
	"appal/router"
)

func main() {
	r := router.New()
	a := news.New()
	go a.Serve()
	r.Run()
}
