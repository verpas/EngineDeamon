package netConnect

import "net"

type ConnectInfo struct {
	con  net.Conn
	ip   string
	port string
}
