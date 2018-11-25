package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/", indexHandler)
	r.GET("/user/:name/*action", userHandler)
	r.GET("/search/:category", collectHandler)
	r.GET("/result/:category", resultHandler)
	return r
}
