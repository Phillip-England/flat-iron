package actionroutes

import "github.com/gin-gonic/gin"

func Logout(r *gin.Engine) {
	r.GET("/logout", func(c *gin.Context) {
		c.SetCookie("session-token", "", -1, "/", "localhost", true, true)
		c.Redirect(303, "/")
	}) 

}