package routes

import (
	"fmt"
	"htmx-scorecard/src/lib"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func PageLocationSelection(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/locations", func(c *gin.Context) {
		user, httpErr := lib.Auth(c, mongoStore)
		if httpErr != nil {
			if httpErr.Code == 500 {
				c.Redirect(303, "/500")
				return
			}
			if httpErr.Code == 401 {
				c.Redirect(303, "/")
				return
			}
		}
		fmt.Println(user)
		c.HTML(200, "PageLocationSelection.html", gin.H{
			"Banner": "Locations",
		})
	})
}