package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}
