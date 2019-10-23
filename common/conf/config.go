/*
 * @Author: dengyue.chen
 * @Date: 2017-03-31 11:36:54
 * @Last Modified by:   dengyue.chen
 * @Last Modified time: 2017-03-31 11:36:54
 */

package conf

import (
	l "ebaocloud.com/go-js-service/common/log"
)

var (
	log = l.SetLog("LoadConfig")
)

const (
	//DefaultServiceName 服务名
	DefaultServiceName = "go-js-service"

	//ConfServiceName service name
	ConfServiceName = "SERVICE_NAME"

	// CasServiceName = "go-cas-service"
	CasServiceName = "172.25.16.153:30016"
	// AdminServiceName ...
	AdminServiceName = "172.25.16.153:30015"
)

//Conf conf
type Conf struct {
	ServiceName string
}

//Cfg config
var Cfg Conf

func init() {
	log.Info("----Read Conf Start----")
	Cfg.ServiceName = DefaultServiceName
	log.Info("++++Read Conf End++++")
}
