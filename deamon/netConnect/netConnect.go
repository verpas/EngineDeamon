package netConnect

import "net"

type ConnectInfo struct {
	con  net.Conn
	IP   string
	port string
}
