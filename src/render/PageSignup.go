package render

import "github.com/gin-gonic/gin"

func PageSignup(c *gin.Context, errSignupForm string) {
	c.HTML(200, "PageLogin.html", gin.H{
		"ErrSignupForm": errSignupForm,
	})
}