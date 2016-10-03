package netConnect

import "net"

type ConnectInfo struct {
	Con  net.Conn
	Ip   string
	Port string
}

func (con *ConnectInfo) Write(context []byte) (err error) {
	var sendlen int = 0
	alllen := len(context)
	var nowSend int = 0
	for alllen != nowSend {
		sendlen, err = con.Con.Write(context[nowSend:])
		nowSend += sendlen
	}
	return
}

func (con *ConnectInfo) Connect(IP string, port string) (err error) {
	ipport := IP + ":" + port
	con.Con, err = net.Dial("tcp", ipport)
	if err == nil {
		con.Ip = IP
		con.Port = port
	}
	return
}

func (con *ConnectInfo) Close() (err error) {
	err = con.Con.Close()
	con.Ip = ""
	con.Port = ""
	return
}
func NewMyConnect(IP string, port string) (con *ConnectInfo, err error) {
	con = new(ConnectInfo)
	err = con.Connect(IP, port)
	return
}
