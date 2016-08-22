package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

var ServiceIpPort = flag.String("ServerIpPort", "", "deamon connect ipport")
var ClientIpPort = flag.String("ClientIpPort", "", "deamon Listen ipport")
var logPath = flag.String("LogFileDirectory", "log", "Set log path")

func main() {
	flag.Parse()
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
		//test vscode pull
		//test22222
		//test3333
	}
}
