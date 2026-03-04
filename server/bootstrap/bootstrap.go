package bootstrap

import (
	"bitalikr1999/panda-kv/packages/commands"
	"bitalikr1999/panda-kv/packages/resp"
	"bitalikr1999/panda-kv/packages/storage"
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

	storage := storage.New(3)

	fmt.Println("Storage created", storage)

	storage.Send(commands.SetCommand{Key: "some ke5"})

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

	if !commands.CanCreateCommand(respData) {
		return
	}

	command, err := commands.Create(respData)
	if err != nil {
		return
	}

	fmt.Println("Got command", command)

}
