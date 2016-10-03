package main

import (
	"flag"
	"fmt"
	"strings"

	"./netConnect"
)

var ServiceIpPort = flag.String("ServerIpPort", "", "deamon connect ipport")
var ClientIpPort = flag.String("ClientIpPort", "", "deamon Listen ipport")
var logPath = flag.String("LogFileDirectory", "log", "Set log path")

func main() {
	flag.Parse()
	var MyConnect *netConnect.ConnectInfo
	fmt.Print(*ServiceIpPort, "\n", *logPath)
	if (*ServiceIpPort) != "" {
		arr := strings.SplitN((*ServiceIpPort), ":", 3)
		if len(arr) != 2 {
			panic("ServerIpPort error")
		}
		var err error
		MyConnect, err = netConnect.NewMyConnect(arr[0], arr[1])
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		MyConnect.Write([]byte("nihao"))
		MyConnect.Close()
	}
}
