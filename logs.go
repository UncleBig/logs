package logs

import (
	"fmt"

	seelog "github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func SetLogConf(logConfig string) {
	//初始化全局变量Logger为seelog的禁用状态，主要为了防止Logger被多次初始化
	var err error
	Logger = seelog.Disabled
	Logger, err = seelog.LoggerFromConfigAsFile(logConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

}
