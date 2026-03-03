package main

import "net"

func main() {

	conn, err := net.Dial("tcp", "localhost:8090")

	if err != nil {
		panic("Errorr")
	}
	defer conn.Close()

	// cmd := "*2\r\n$3\r\nGET\r\n$4\r\nkey1\r\n"
	cmd := "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nhello\r\n"
	conn.Write([]byte(cmd))
}
