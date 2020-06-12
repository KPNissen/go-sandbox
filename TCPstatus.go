package main

import (
	"bufio"
	"fmt"
	"net"
)

func main () {
	conn. _ := net.Dial ("tcp", "golang.org:80")
	fmt.FPrintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ :=
	fmt.Println(status)
}
