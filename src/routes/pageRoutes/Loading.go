package pageroutes

import (
	"htmx-scorecard/src/lib"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func Loading(r *gin.Engine, mongoStore *types.MongoStore) {
	r.POST("/loading", func(c *gin.Context) {
		lib.SleepyRoute()
		c.Data(200, "text/html; charset=utf-8", nil)
	})
}
