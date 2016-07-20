// test
package main

import (
	"charles/logs"
	"charles/utils"
	"flag"
	"fmt"
	"os"

	"github.com/dlintw/goconf"
)

func main() {

	conf_file := flag.String("config", "./server.ini", "set run config file.")
	flag.Parse()

	l_conf, err := goconf.ReadConfigFile(*conf_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadConfiguration: Error: Could not open %q for reading: %s\n", conf_file, err)
		os.Exit(1)
	}
	var log_conf, _ = l_conf.GetString("log", "log_conf")
	logs.SetLogConf(log_conf)
	logs.Logger.Trace("Trance")
	logs.Logger.Debug("Debug")
	logs.Logger.Info("Info")
	logs.Logger.Error("Error")
	logs.Logger.Critical("Critical")
	//do not forget flush
	defer logs.Logger.Flush()

	return
}
