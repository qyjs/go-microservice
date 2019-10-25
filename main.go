package main

import (
	"io"
	"os"

	"ebaocloud.com/go-js-service/common/handler"
	_ "ebaocloud.com/go-js-service/docs"
	"ebaocloud.com/go-js-service/src/auth"
	"ebaocloud.com/go-js-service/src/qrcode"
	"ebaocloud.com/go-js-service/src/swagger"
	"ebaocloud.com/go-js-service/src/tenantInfo"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

	// database, err := sqlx.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	// if err != nil {
	// 	fmt.Println("open mysql failed,", err)
	// 	return
	// }
	// Db = database
	// fmt.Println("aaa")
}

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

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := setupRouter()

	r.Run(":1080")
}
