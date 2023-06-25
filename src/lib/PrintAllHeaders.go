package lib

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PrintAllHeaders(c *gin.Context) {
	for key, values := range c.Request.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}