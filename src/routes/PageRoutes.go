package routes

import (
	"htmx-scorecard/src/render"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)


func PageLogin(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/", func(c *gin.Context) {
		render.PageLogin(c, "")
	})
}

func PageSignup(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/signup", func(c *gin.Context) {
		render.PageSignup(c, "")
	})
}

func PageServerError(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/signup", func(c *gin.Context) {
		render.PageServerError(c, "")
	})
}

func PageLocationSelection(r *gin.Engine, mongoStore *types.MongoStore) {
	r.GET("/locations", func(c *gin.Context) {
		render.PageLocationSelection(c)
	})
}