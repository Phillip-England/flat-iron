package routes

import "github.com/gin-gonic/gin"

func NavUserOpened(r *gin.Engine) {
	r.POST("/components/NavUserOpened/:Banner", func(c *gin.Context) {
		banner := c.Param("Banner")
		c.HTML(200, "NavUserOpened.html", gin.H{
			"Banner": banner,
		})
	})
}