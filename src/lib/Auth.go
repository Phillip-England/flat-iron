package lib

import (
	"htmx-scorecard/src/database/userdb"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context, mongoStore *types.MongoStore) (*types.User, *types.HttpErr) {
	cookie, err := c.Cookie("session-token")
	if err != nil {
		return nil, types.NewHttpErr(0, 401, "unauthorized")
	}
	user, httpErr := userdb.FindUserBySession(cookie, mongoStore)
	if httpErr != nil {
		return nil, httpErr
	}
	return user, nil
}