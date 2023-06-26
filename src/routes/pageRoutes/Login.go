package pageRoutes

import (
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "PageLogin.html", gin.H{
			"ErrLoginForm": c.Query("ErrLoginForm"),
			"Banner": "Chick-fil-A Tools",
		})
	})
}

