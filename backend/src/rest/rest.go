package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default()
	r.GET("/relativepath/to/url", func(c *gin.Context) {

	})
}
