package routes

import (
	"fmt"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func ActionCreateLocation(r *gin.Engine, mongoStore *types.MongoStore) {
	r.POST("/actions/CreateLocation", func(c *gin.Context) {
		location := types.NewLocation(c.PostForm("name"), c.PostForm("number"))
		fmt.Println(location)
		c.HTML(200, "PageLocationSelection.html", gin.H{
			"Banner": "Locations",
		})
	})
}