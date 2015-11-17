// xiaoyu chen
// 358860528@qq.com

package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"flag"
)

var (
	sendfile = []byte(`<?xml version="1.0"?>
<cross-domain-policy>
   <allow-access-from domain="*" to-ports="*" secure="false" />
</cross-domain-policy>
`)
)

var listen = flag.String("svrip", ":843", "listen ip:port")

func main() {
	flag.Parse()
	
	ln, err := net.Listen("tcp", *listen)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("start flash sandbox succeed~~!")

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(c)
	}

}

func handleConnection(c net.Conn) {
	bytes := make([]byte, 1024)
	
	fmt.Println("handleConnection:",c.RemoteAddr())
	_, err := c.Read(bytes)
	if err != nil {
		fmt.Println(err)
		fmt.Println("err:",c.RemoteAddr())
		return
	}

	writer := bufio.NewWriter(c)
	writer.Write(sendfile)
	writer.Flush()
	//fmt.Println("recv:" + c.LocalAddr().String())
	//if string(bytes[:l]) == "<policy-file-request/>" {
	time.Sleep(time.Second * 1)
	c.Close()
	fmt.Println("Close:",c.RemoteAddr())
	//}

}
