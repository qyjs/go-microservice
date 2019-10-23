package handler

import (

	//

	"github.com/inconshreveable/log15"
	//
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	"ebaocloud.com/go-js-service/common/conf"
	l "ebaocloud.com/go-js-service/common/log"
)

//Log 日志
func Log() gin.HandlerFunc {
	handler := l.GetLogHandler()
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		newLog := log15.New("Method", context.Request.Method, "Api", context.Request.URL.Path, "IP", context.ClientIP(), "Authorization", token, "X-ebao-user-identity", context.Request.Header.Get("X-ebao-user-identity"), "appname", conf.Cfg.ServiceName)
		if gin.Mode() != gin.DebugMode {
			newLog = newLog.New("log_format", "json")
		}
		newLog.SetHandler(handler)

		context.Set("log", newLog)
		context.Next()
	}
}
