package routes

import (
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)


func PageLogin(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "PageLogin.html", gin.H{
			"ErrLoginForm": c.Query("ErrLoginForm"),
			"Banner": "Chick-fil-A Tools",
		})
	})
}

func PageSignup(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "PageSignup.html", gin.H{
			"ErrSignupForm": "",
			"Banner": "Chick-fil-A Tools",
		})
	})
}

func PageServerError(r *gin.Engine) {
	r.GET("/500", func(c *gin.Context) {
		c.HTML(200, "500.html", gin.H{
			"ErrServer": c.Query("ErrServer"),
			"Banner": "505 Server Error",
		})
	})
}

func PageLocationSelection(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/locations", func(c *gin.Context) {
		c.HTML(200, "PageLocationSelection.html", gin.H{
			"Banner": "Locations",
		})
	})
}