package render

import "github.com/gin-gonic/gin"

func PageLocationSelection(c *gin.Context) {
	c.HTML(200, "PageLocationSelection.html", gin.H{
		"Banner": "Locations",
	})
}

