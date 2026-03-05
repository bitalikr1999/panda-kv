package storage

import (
	"bitalikr1999/panda-kv/packages/commands"
	"errors"
)

type Shard struct {
	inputChan StorageInChannel
	data      map[string]ShardValue
}

func NewShard() *Shard {

	return &Shard{
		inputChan: make(StorageInChannel),
		data:      make(map[string]ShardValue),
	}

}

func (s *Shard) Start() {

	for input := range s.inputChan {

		input.responseChan <- s.execute(input.command)
	}
}

func (s *Shard) Stop() {
	close(s.inputChan)
}

func (s *Shard) Send(command commands.Command) CommandExecuteResponse {

	responseChan := make(chan CommandExecuteResponse)

	s.inputChan <- ChanMessage{
		responseChan: responseChan,
		command:      command,
	}

	response := <-responseChan

	return response

}

func (s *Shard) execute(command commands.Command) CommandExecuteResponse {

	switch c := command.(type) {

	case commands.SetCommand:
		{
			return s.save(c.Key, c.Value)
		}

	case commands.GetCommand:
		{
			return s.get(c.Key)
		}
	}

	return CommandExecuteResponse{
		Ok:    false,
		Error: errors.New("cant indentify comand"),
	}
}

func (s *Shard) save(key string, value string) CommandExecuteResponse {
	v := ShardValue{
		Value: value,
	}

	s.data[key] = v

	return CommandExecuteResponse{
		Ok:    true,
		Value: v,
	}
}

func (s *Shard) get(key string) CommandExecuteResponse {
	v, ok := s.data[key]
	if !ok {
		return CommandExecuteResponse{
			Ok:    false,
			Error: errors.New("not found"),
		}
	}

	return CommandExecuteResponse{
		Ok:    true,
		Value: v,
	}
}
