package websever

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestWebGin(t *testing.T) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")
}
