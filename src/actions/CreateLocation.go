package actions

import (
	"fmt"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func CreateLocation(c *gin.Context, mongoStore *types.MongoStore) {
	location := types.NewLocation(c.PostForm("name"), c.PostForm("number"))
	fmt.Println(location)
	c.HTML(200, "PageLocationSelection.html", gin.H{
		"Banner": "Locations",
	})
}