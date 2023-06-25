package routes

import (
	"htmx-scorecard/src/actions"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func ActionLoginUser(r *gin.Engine, mongoStore *types.MongoStore) {
	r.POST("/actions/LoginUser", func(c *gin.Context) {
		actions.LoginUser(c, mongoStore)
	})
}

func ActionCreateLocation(r *gin.Engine, mongoStore *types.MongoStore) {
	r.POST("/actions/CreateLocation", func(c *gin.Context) {
		actions.CreateLocation(c, mongoStore)
	})
}