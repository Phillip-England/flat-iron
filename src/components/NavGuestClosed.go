package components

import (
	"os"

	"github.com/gin-gonic/gin"
)

func NavGuestClosed(r *gin.Engine) {

	r.POST("/components/NavGuestClosed", func(c *gin.Context) {
		html, _ := os.ReadFile("./templates/components/NavGuestClosed.html")
		c.Data(200, "text/html; charset=utf-8", html)
	})

}