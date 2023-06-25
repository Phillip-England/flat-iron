package routes

import (
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func PageLocationSelection(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/locations", func(c *gin.Context) {
		c.HTML(200, "PageLocationSelection.html", gin.H{
			"Banner": "Locations",
		})
	})
}