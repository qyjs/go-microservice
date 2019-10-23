package log

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

//SetLog 设置日志
func SetLog(functionName string) log15.Logger {
	log := log15.New("Function", functionName)
	if gin.Mode() != gin.DebugMode {
		log = log.New("log_format", "json")
	}
	handler := GetLogHandler()

	log.SetHandler(handler)

	return log
}

//GetLogHandler GetLogHandler
func GetLogHandler() log15.Handler {
	var handler log15.Handler
	handler = log15.StreamHandler(os.Stdout, log15.JsonFormat())
	// if gin.Mode() == gin.DebugMode {
	// 	handler = log15.StreamHandler(os.Stdout, log15.TerminalFormat())
	// } else {
	// 	handler = log15.MultiHandler(
	// 		log15.StreamHandler(os.Stdout, log15.TerminalFormat()),
	// 		log15.LvlFilterHandler(log15.LvlInfo, log15.Must.FileHandler("/app/logs/go-msg.log", log15.JsonFormat())),
	// 	)
	// 	handler = log15.BufferedHandler(10000, handler)
	// }
	return log15.CallerFileHandler(handler)
}
