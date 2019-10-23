package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
)

// Swagger ...
func Swagger(c *gin.Context) {
	swaggerfile, _ := swag.ReadDoc()
	c.String(200, "%s", swaggerfile)
}
