package pageroutes

import (
	"htmx-scorecard/src/lib"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/", func(c *gin.Context) {
		// if we no cookie, go to log in page
		sessionToken, _ := c.Cookie("session-token")
		if sessionToken == "" {
			c.HTML(200, "PageLogin.html", gin.H{
				"ErrLoginForm": c.Query("ErrLoginForm"),
				"Banner": "Chick-fil-A Tools",
			})
			return
		}
		user, _ := lib.Auth(c, mongoStore)
		if user != nil {
			c.Redirect(303, "/locations")
			return
		}
		c.Redirect(303, "/logout")
	})
}

