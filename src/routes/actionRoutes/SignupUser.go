package actionroutes

import (
	"fmt"
	"htmx-scorecard/src/database/userdb"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func SignupUser(r *gin.Engine, mongoStore *types.MongoStore) {
	r.POST("/actions/SignupUser", func(c *gin.Context) {
		user := types.NewUser(c.PostForm("email"), c.PostForm("password"))
		httpErr := userdb.Insert(user, mongoStore.UserCollection)
		if httpErr != nil {
			if httpErr.Code == 500 {
				c.Redirect(303, fmt.Sprintf("/500?ErrServer=%s", httpErr.Message))
			}
			if httpErr.Code == 400 {
				c.Redirect(303, fmt.Sprintf("/signup?ErrSignupForm=%s", httpErr.Message))
				return
			}
		}
		c.Redirect(303, "/")
	})
}