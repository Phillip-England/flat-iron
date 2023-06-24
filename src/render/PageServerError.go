package render

import "github.com/gin-gonic/gin"

func PageServerError(c *gin.Context, errSignupForm string) {
	c.HTML(200, "PageServerError.html", gin.H{
		"ErrSignupForm": errSignupForm,
	})
}