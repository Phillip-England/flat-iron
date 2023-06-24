package routes

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

func NavGuestClosed(r *gin.Engine) {
	r.POST("/components/NavGuestClosed", func(c *gin.Context) {
		html, _ := os.ReadFile("./templates/components/NavGuestClosed.html")
		c.Data(200, "text/html; charset=utf-8", html)
	})
}

func NavUserOpened(r *gin.Engine) {
	r.POST("/components/NavUserOpened/:Banner", func(c *gin.Context) {
		banner := c.Param("Banner")
		c.HTML(200, "NavUserOpened.html", gin.H{
			"Banner": banner,
		})
	})
}

func NavUserClosed(r *gin.Engine) {
	r.POST("/components/NavUserClosed/:Banner", func(c *gin.Context) {
		banner := c.Param("Banner")
		c.HTML(200, "NavUserClosed.html", gin.H{
			"Banner": banner,
		})
	})
}




