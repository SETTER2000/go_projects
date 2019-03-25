package router

import (
	"appal/news"
	"github.com/gin-gonic/gin"
	"net/http"
)

func collectHandler(c *gin.Context) {
	category := c.Param("category")
	news.Collect(category)
	c.String(http.StatusOK, "Search is initialized")
}

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Category:\ngeneral\ntechnology\nsports\nbusiness\nentertainment\nscience\n\n\tExample:\n\n\t1. /search/general\n\t2. /result/general")
}

func userHandler(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func resultHandler(c *gin.Context) {
	category := c.Param("category")
	topics := news.Result(category)

	c.JSON(http.StatusOK, topics)
}

func welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
