package componentroutes

import "github.com/gin-gonic/gin"

func NavGuestClosed(r *gin.Engine) {
	r.POST("/components/NavGuestClosed", func(c *gin.Context) {
		banner := c.Query("Banner")
		c.HTML(200, "NavGuestClosed.html", gin.H{
			"Banner": banner,
		})
	})
}