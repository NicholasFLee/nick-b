package routers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// HTMLFileHandler handlers front-end html file
func HTMLFileHandler(c *gin.Context) {
	path := strings.TrimPrefix(c.Request.RequestURI, "/blog/")
	if path == "" {
		path = "index.html"
	}
	uri := fmt.Sprintf("../nick-f/dist/%s", path)
	c.File(uri)
}
