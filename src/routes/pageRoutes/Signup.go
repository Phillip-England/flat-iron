package pageroutes

import (
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func Signup(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "PageSignup.html", gin.H{
			"ErrSignupForm": "",
			"Banner": "Chick-fil-A Tools",
		})
	})
}