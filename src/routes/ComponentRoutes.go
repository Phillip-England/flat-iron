package routes

import (
	"github.com/gin-gonic/gin"
)

func NavGuestOpened(r *gin.Engine) {
	r.POST("/components/NavGuestOpened", func(c *gin.Context) {
		banner := c.Query("Banner")
		c.HTML(200, "NavGuestOpened.html", gin.H{
			"Banner": banner,
		})
	})
}

func NavGuestClosed(r *gin.Engine) {
	r.POST("/components/NavGuestClosed", func(c *gin.Context) {
		banner := c.Query("Banner")
		c.HTML(200, "NavGuestClosed.html", gin.H{
			"Banner": banner,
		})
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




