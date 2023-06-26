package actionRoutes

import (
	"fmt"
	"htmx-scorecard/src/types"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(r *gin.Engine, mongoStore *types.MongoStore) {
	r.POST("/actions/LoginUser", func(c *gin.Context) {
		user := types.NewUser(c.PostForm("email"), c.PostForm("password"))
		httpErr := user.Exists(mongoStore.UserCollection)
		if httpErr != nil {
			if httpErr.Code == 500 {
				c.Redirect(303, fmt.Sprintf("/500?ErrServer=%s", httpErr.Message))
				return
			}
			if httpErr.Code == 400 {
				c.Redirect(303, fmt.Sprintf("/?ErrLoginForm=%s", httpErr.Message))
				return
			}
		}
		
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password"))); err != nil {
			c.Redirect(303, fmt.Sprintf("/?ErrLoginForm=%s", "invalid credentials"))
			return
		}
			sessionModel := types.NewSession(user.Id)
			httpErr = sessionModel.ClearUserSessions(mongoStore.SessionCollection)
			if httpErr != nil {
				if httpErr.Code == 500 {
					c.Redirect(303, fmt.Sprintf("/500?ErrServer=%s", httpErr.Message))
					return
				}
			}
			httpErr = sessionModel.Insert(mongoStore.SessionCollection)
			if httpErr != nil {
				if httpErr.Code == 500 {
					c.Redirect(303, fmt.Sprintf("/500?ErrServer=%s", httpErr.Message))
					return
				}
			}
			sessionToken := sessionModel.Id.Hex()
			c.SetCookie("session-token", sessionToken, 86400, "/", os.Getenv("DOMAIN"), true, true)
			c.Redirect(303, "/locations")
	})
}