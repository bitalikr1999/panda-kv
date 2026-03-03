package bootstrap

import (
	"bitalikr1999/panda-kv/packages/commands"
	"bitalikr1999/panda-kv/packages/resp"
	"bufio"
	"fmt"
	"log"
	"net"
)

func Start() {

	fmt.Println("Starting Panda KV...")
	listener, err := net.Listen("tcp", ":8090")

	if err != nil {
		log.Fatal("Error starting listener", err)
	}

	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error handling connection", err)
			continue
		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	respData, err := resp.ParseResp(reader)
	if err != nil {
		return
	}

	if commands.CanCreateCommand(respData) == false {
		return
	}

	command, err := commands.Create(respData)
	if err != nil {
		return
	}

	fmt.Println("Got command", command)

}
