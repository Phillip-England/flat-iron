package actionroutes

import (
	"fmt"
	"htmx-scorecard/src/database/sessiondb"
	"htmx-scorecard/src/database/userdb"
	"htmx-scorecard/src/types"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(r *gin.Engine, mongoStore *types.MongoStore) {
	r.POST("/actions/LoginUser", func(c *gin.Context) {
		user := types.NewUser(c.PostForm("email"), c.PostForm("password"))
		httpErr := userdb.Exists(user, mongoStore.UserCollection)
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
			session := types.NewSession(user.Id)
			httpErr = sessiondb.ClearUserSessions(session, mongoStore.SessionCollection)
			if httpErr != nil {
				if httpErr.Code == 500 {
					c.Redirect(303, fmt.Sprintf("/500?ErrServer=%s", httpErr.Message))
					return
				}
			}
			httpErr = sessiondb.Insert(session, mongoStore.SessionCollection)
			if httpErr != nil {
				if httpErr.Code == 500 {
					c.Redirect(303, fmt.Sprintf("/500?ErrServer=%s", httpErr.Message))
					return
				}
			}
			sessionToken := session.Id.Hex()
			c.SetCookie("session-token", sessionToken, 86400, "/", os.Getenv("DOMAIN"), true, true)
			c.Redirect(303, "/locations")
	})
}