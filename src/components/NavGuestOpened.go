package components

import (
	"os"

	"github.com/gin-gonic/gin"
)

func NavGuestOpened(r *gin.Engine) {

	r.POST("/components/NavGuestOpened", func(c *gin.Context) {
		html, _ := os.ReadFile("./templates/components/NavGuestOpened.html")
		c.Data(200, "text/html; charset=utf-8", html)
	})

}