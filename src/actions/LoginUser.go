package actions

import (
	"htmx-scorecard/src/render"
	"htmx-scorecard/src/types"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context, mongoStore *types.MongoStore) *types.HttpErr {
		user := types.NewUser(c.PostForm("email"), c.PostForm("password"))
		httpErr := user.Find(mongoStore.UserCollection)
		if httpErr != nil {
			if httpErr.Code == 500 {
				render.PageServerError(c, httpErr.Message)
				return nil
			}
			if httpErr.Code == 400 {
				render.PageLogin(c, httpErr.Message)
				return nil
			}
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password"))); err != nil {
			render.PageLogin(c, httpErr.Message)
			return nil
		}
		sessionModel := types.NewSession(user.Id)
		httpErr = sessionModel.ClearUserSessions(mongoStore.SessionCollection)
		if httpErr != nil {
			if httpErr.Code == 500 {
				render.PageServerError(c, httpErr.Message)
				return nil
			}
		}
		httpErr = sessionModel.Insert(mongoStore.SessionCollection)
		if httpErr != nil {
			if httpErr.Code == 500 {
				render.PageServerError(c, httpErr.Message)
				return nil
			}
		}
		sessionToken := sessionModel.Id.Hex()
		c.SetCookie("session-token", sessionToken, 86400, "/", os.Getenv("DOMAIN"), true, true)
		render.PageLocationSelection(c)
		return nil
}