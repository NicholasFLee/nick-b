package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// HTMLFileHandler handlers front-end html file
func HTMLFileHandler(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		path = "index.html"
	}
	uri := fmt.Sprintf("../nick-f/dist/%s", path)
	c.File(uri)
}
