package main

import (
	"fmt"
	"net"
	"time"

	"./gotest"
)

type hhh struct {
	int
	string
}

func gogogo(c chan int) {
	fmt.Print("begin sleep")
	time.Sleep(time.Duration(5) * time.Second)
	c <- 1
	fmt.Print("gggggggg")
}
func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Print("listen error")
		return
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			continue
		}
		fmt.Print("fsfdsgfdgfd")
		readstr := make([]byte, 10000)
		for {
			len, err := con.Read(readstr)
			fmt.Print("read:", len, string(readstr))
			str := readstr[0:len]
			fmt.Print("write:   ", +len, "    ", string(str)+"     "+con.LocalAddr().String()+"\n")
			len, err = con.Write(str)
			if err != nil {
				fmt.Print(err.Error())
				con.Close()
				break
			}

		}

	}
	ch := make(chan int)
	go gogogo(ch)
	//var h int
	h := <-ch

	fmt.Print(h)
	fhhh := new(hhh)
	fmt.Print("dfsdfds\n")
	fhhh.int = 0
	fhhh.string = "fdsfdsf"
	fmt.Printf("%v\n%+v\n%T\n%#v\n", fhhh, fhhh, fhhh, fhhh)
	gotest.Print(1)
}
