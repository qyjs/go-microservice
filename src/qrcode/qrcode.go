package qrcode

import (
	"encoding/base64"
	"net/http"

	"ebaocloud.com/go-js-service/common/response"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"

	gq "github.com/skip2/go-qrcode"
)

//Args parameter
type Args struct {
	Content string `form:"content" json:"content" binding:"required"`
}

type failResponse struct {
	response.Response
}

//GenerateQrcodeByPath GenerateQrcode
// @Summary GenerateQrcode
// @Description get qrcode by content
// @Accept  json
// @Produce  json
// @Param   content      query string     true  "content"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} response.Response "error response"
// @Router /base64 [get]
func GenerateQrcodeByPath(c *gin.Context) {
	log := c.MustGet("log").(log15.Logger)

	args := &Args{}

	if err := c.Bind(args); err != nil {
		log.Error("bind parameter error", "error", err)
		c.JSON(http.StatusBadRequest, response.NewFail("bind parameter error", 99))
		return
	}
	if args.Content == "" {
		log.Error("parameter is blank")
		c.JSON(http.StatusBadRequest, response.NewFail("parameter is blank", 99))
		return
	}

	png, err := encodeQrcode(args.Content)
	if err != nil {
		log.Error("generate qrcode error", "error", err)
		c.JSON(http.StatusBadRequest, response.NewFail("generate qrcode error:"+err.Error(), 99))
		return
	}

	c.String(http.StatusOK, "%s", base64.StdEncoding.EncodeToString(png))
	return
}

//GenerateQrcodeByBody GenerateQrcodeByBody
func GenerateQrcodeByBody(c *gin.Context) {
	log := c.MustGet("log").(log15.Logger)

	args := &Args{}

	if err := c.Bind(args); err != nil {
		log.Error("bind parameter error", "error", err)
		c.JSON(http.StatusBadRequest, response.NewFail("bind parameter error", 99))
		return
	}
	if args.Content == "" {
		log.Error("parameter is blank")
		c.JSON(http.StatusBadRequest, response.NewFail("parameter is blank", 99))
		return
	}

	png, err := encodeQrcode(args.Content)
	if err != nil {
		log.Error("generate qrcode error", "error", err)
		c.JSON(http.StatusBadRequest, response.NewFail("generate qrcode error:"+err.Error(), 99))
		return
	}

	c.String(http.StatusOK, "%s", base64.StdEncoding.EncodeToString(png))
	return
}

func encodeQrcode(content string) ([]byte, error) {
	return gq.Encode(content, gq.Medium, 256)
}

//GenerateQrcodePng GenerateQrcodePng
// @Summary GenerateQrcodePng
// @Description get qrcode png by content
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   content      query string     true  "content"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} response.Response "error response"
// @Router /png [get]
func GenerateQrcodePng(c *gin.Context) {
	log := c.MustGet("log").(log15.Logger)

	args := &Args{}

	if err := c.Bind(args); err != nil {
		log.Error("bind parameter error", "error", err)
		c.JSON(http.StatusBadRequest, response.NewFail("bind parameter error", 99))
		return
	}
	if args.Content == "" {
		log.Error("parameter is blank")
		c.JSON(http.StatusBadRequest, response.NewFail("parameter is blank", 99))
		return
	}

	png, err := encodeQrcode(args.Content)
	if err != nil {
		log.Error("generate qrcode error", "error", err)
		c.JSON(http.StatusBadRequest, response.NewFail("generate qrcode error:"+err.Error(), 99))
		return
	}

	c.Data(http.StatusOK, "image/png", png)
	return
}
