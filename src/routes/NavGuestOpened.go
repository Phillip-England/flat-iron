package routes

import "github.com/gin-gonic/gin"

func NavGuestOpened(r *gin.Engine) {
	r.POST("/components/NavGuestOpened", func(c *gin.Context) {
		banner := c.Query("Banner")
		c.HTML(200, "NavGuestOpened.html", gin.H{
			"Banner": banner,
		})
	})
}