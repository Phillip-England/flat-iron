package pages

import (
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func PageSignup(r *gin.Engine, mongoStore *types.MongoStore) {

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "PageSignup.html", gin.H{
			"Path": "/",
			"ErrLoginForm": "",
		})
	})


}