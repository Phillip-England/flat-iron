package pages

import (
	"htmx-scorecard/src/controllers"
	"htmx-scorecard/src/lib"
	"htmx-scorecard/src/types"

	"github.com/gin-gonic/gin"
)



func PageLogin(r *gin.Engine, mongoStore *types.MongoStore) {

	r.GET("/", func(c *gin.Context) {

		//-------------------------
		// SERVER ACTIONS
		//-------------------------

		action := c.GetHeader("ServerAction")

		//----------------------------------------------------
		//  NO ACTION -> SERVE THE PAGE IN ITS DEFAULT STATE
		//----------------------------------------------------

		if action == "" {
			c.HTML(200, "PageLogin.html", gin.H{
				"Props": types.NewPropsPageLogin(""),
			})
			return
		}

		//-----------------------------
		//  VALIDATING SERVER ACTION
		//-----------------------------


		actions := []string{
			"LoginUser",
		}
		if !lib.IsStringInSlice(action, actions) {
			c.HTML(200, "PageServerError.html", gin.H{
				"Props": types.NewPropsPageServerError("invalid action"),
			})
			return
		}

		//-----------------------------
		// EXECUTING ACTIONS
		//-----------------------------

		if action == "LoginUser" {
			controllers.LoginUser()
		}
		
	})


}