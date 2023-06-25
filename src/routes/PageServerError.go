package routes

import "github.com/gin-gonic/gin"

func PageServerError(r *gin.Engine) {
	r.GET("/500", func(c *gin.Context) {
		c.HTML(200, "500.html", gin.H{
			"ErrServer": c.Query("ErrServer"),
			"Banner": "505 Server Error",
		})
	})
}