package main

import (
	"flag"
	"fmt"
	"net"
	"strings"

	"./netConnect"
)

var ServiceIpPort = flag.String("ServerIpPort", "", "deamon connect ipport")
var ClientIpPort = flag.String("ClientIpPort", "", "deamon Listen ipport")
var logPath = flag.String("LogFileDirectory", "log", "Set log path")

func main() {
	flag.Parse()
	var MyConnect netConnect.ConnectInfo
	fmt.Print(*ServiceIpPort, "\n", *logPath)
	if (*ServiceIpPort) != "" {
		arr := strings.SplitN((*ServiceIpPort), ":", 3)
		if len(arr) != 2 {
			panic("ServerIpPort error")
		}
		con, err := net.Dial("tcp", *ServiceIpPort)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
		con.Write([]byte("nihao"))
		con.Close()
	}
}
