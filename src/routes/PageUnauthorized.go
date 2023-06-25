package routes

import "github.com/gin-gonic/gin"

func PageUnauthorized(r *gin.Engine) {
	r.GET("/401", func(c *gin.Context) {
		c.HTML(200, "401.html", gin.H{
			"Banner": "Unauthorized",
		})
	})
}