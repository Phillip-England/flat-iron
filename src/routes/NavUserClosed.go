package routes

import "github.com/gin-gonic/gin"

func NavUserClosed(r *gin.Engine) {
	r.POST("/components/NavUserClosed/:Banner", func(c *gin.Context) {
		banner := c.Param("Banner")
		c.HTML(200, "NavUserClosed.html", gin.H{
			"Banner": banner,
		})
	})
}