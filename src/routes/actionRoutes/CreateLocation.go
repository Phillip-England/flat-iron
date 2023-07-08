package actionroutes

import (
	"fmt"
	"htmx-scorecard/src/lib"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func CreateLocation(r *gin.Engine, mongoStore *types.MongoStore) {
	r.POST("/actions/CreateLocation", func(c *gin.Context) {
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
		location := types.NewLocation(user.Id, c.PostForm("name"), c.PostForm("number"))
		fmt.Println(location.User, location.Name, location.Number)
		c.Redirect(303, "/locations")
	})
}