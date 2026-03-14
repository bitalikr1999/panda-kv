package bootstrap

import (
	"bitalikr1999/panda-kv/packages/commands"
	"bitalikr1999/panda-kv/packages/resp"
	string_encoder "bitalikr1999/panda-kv/packages/resp/encoders/string"
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

	defer listener.Close()
	defer storage.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error handling connection", err)
			continue
		}

		go handleConnection(conn, storage)

	}

}

func handleConnection(conn net.Conn, storage *storage.Storage) {

	defer conn.Close()

	for {
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

		response := storage.Send(command)

		fmt.Println(response.Value.Value)

		if response.Ok && response.Value.Value == "" {
			fmt.Println("asdasd")
			_, err := conn.Write(string_encoder.EncodeString("OK"))
			if err != nil {
				fmt.Println("Some error")
			}
		}
	}

}
