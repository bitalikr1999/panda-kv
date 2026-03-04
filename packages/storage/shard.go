package storage

import (
	"bitalikr1999/panda-kv/packages/commands"
	"fmt"
)

type Shard struct {
	inputChan StorageInChannel
	data      map[string]string
}

func NewShard() *Shard {

	return &Shard{
		inputChan: make(StorageInChannel),
	}

}

func (s *Shard) Start() {

	for input := range s.inputChan {

		fmt.Println("Got command, yuh", input)

		input.responseChan <- "Ok!"
	}
}

func (s *Shard) Stop() {
	close(s.inputChan)
}

func (s *Shard) Send(command commands.Command) {

	responseChan := make(chan interface{})

	s.inputChan <- ChanMessage{
		responseChan: responseChan,
	}

	response := <-responseChan

	fmt.Println("Response", response)

}
