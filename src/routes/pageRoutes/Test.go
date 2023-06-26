package pageroutes

import (
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func Test(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/test", func(c *gin.Context) {
		c.HTML(200, "Test.html", nil)
	})
}