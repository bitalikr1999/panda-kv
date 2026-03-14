package main

import (
	"bitalikr1999/panda-kv/packages/resp"
	"bitalikr1999/panda-kv/packages/resp/encoders/bulkstring"
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8090")

	if err != nil {
		panic("Errorr")
	}

	defer conn.Close()

	done := make(chan struct{})

	go listener(conn, done)

	cli(conn)

	close(done)

}

func listener(conn net.Conn, done chan struct{}) {

	reader := bufio.NewReader(conn)

	for {
		select {
		case <-done:
			return
		default:
			{
				data, err := resp.ParseResp(reader)
				if err != nil {
					select {
					case <-done:
						return
					default:
						{
							fmt.Println("Some error", err)
							conn.Close()
							return
						}
					}
				}
				fmt.Println("Data", data)
			}
		}

	}
}

func cli(conn net.Conn) {

	for {

		fmt.Println("Enter a command. GET | SET | DELETE | EXIT")
		var command string

		fmt.Scanln(&command)

		switch strings.ToUpper(command) {
		case "EXIT":
			{
				return
			}
		case "GET":
			{
				command, err := readGetCommand()
				if err != nil {
					continue
				}
				fmt.Println("Command", command)
				conn.Write(command)
			}

		case "SET":
			{
			}
		case "DELETE":
			{
			}
		}

	}

}

func readGetCommand() ([]byte, error) {

	var key string
	fmt.Println("Enter key")
	_, err := fmt.Scan(&key)
	if err != nil {
		return nil, err
	}

	return bulkstring.EncodeBulkstring("GET " + key)
}
