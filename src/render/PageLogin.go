package render

import "github.com/gin-gonic/gin"

func PageLogin(c *gin.Context, errLoginForm string) {
	c.HTML(200, "PageLogin.html", gin.H{
		"ErrLoginForm": errLoginForm,
	})
}

