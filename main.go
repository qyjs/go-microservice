package main

import (
	"ebaocloud.com/go-js-service/common/handler"
	_ "ebaocloud.com/go-js-service/docs"
	"ebaocloud.com/go-js-service/src/auth"
	"ebaocloud.com/go-js-service/src/qrcode"
	"ebaocloud.com/go-js-service/src/swagger"
	"ebaocloud.com/go-js-service/src/tenantInfo"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(handler.Log()) //log15日志
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/health", func(c *gin.Context) {
		c.String(200, "%s", "healthy")
	})

	swg := r.Group("api")
	{
		swg.GET("/swagger.json", swagger.Swagger)
	}

	gettenant := r.Group("api")
	{
		gettenant.GET("/tenant", tenantInfo.GetAllTenant)
	}

	rq := r.Group("/qrcode")
	{
		rq.GET("/base64", qrcode.GenerateQrcodeByPath)
		rq.POST("/base64", qrcode.GenerateQrcodeByBody)

		rq.GET("/png", qrcode.GenerateQrcodePng)
		rq.POST("/png", qrcode.GenerateQrcodePng)
	}

	fleet := r.Group("/fleet")
	{
		fleet.GET("/auth", auth.GetTokenAndTenant)
	}

	return r
}

func main() {
	r := setupRouter()

	r.Run()
}
