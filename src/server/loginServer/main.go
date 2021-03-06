package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"server/loginServer/account"
)

var end = make(chan int)

var config *account.Config
var deal_4g *Deal4G
var deal_4c *Deal4C

func init() {
	config = new(account.Config)
	config.Init()

	deal_4g = new(Deal4G)
	deal_4g.Init()

	deal_4c = new(Deal4C)
	deal_4c.Init()

}

func CheckError(err error) bool {
	if err != nil {
		fmt.Println("err3:", err.Error())
		return false
	}
	return true
}

func SendPackage(conn net.Conn, pid int, body []byte) {
	var pid_32 int32 = int32(pid)
	len := 8 + len(body)
	var len_32 = int32(len)

	len_buf := bytes.NewBuffer([]byte{})
	binary.Write(len_buf, binary.BigEndian, len_32)

	pid_buf := bytes.NewBuffer([]byte{})
	binary.Write(pid_buf, binary.BigEndian, pid_32)

	msg := append(len_buf.Bytes(), pid_buf.Bytes()...)
	msg2 := append(msg, body...)
	conn.Write(msg2)
}

func main() {

	//Deal4Client
	listener, err := net.Listen("tcp", config.Listen4CAddress)

	//Deal4Game
	listener2, err2 := net.Listen("tcp", config.Listen4GameAddress)

	if !CheckError(err) || !CheckError(err2) {
		return
	}

	go deal_4c.Deal4Client(listener)
	go deal_4g.Deal4GameServer(listener2)

	<-end
}
