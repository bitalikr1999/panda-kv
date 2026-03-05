package storage

import "bitalikr1999/panda-kv/packages/commands"

type ChanMessage struct {
	responseChan chan CommandExecuteResponse
	command      commands.Command
}

type StorageInChannel = chan ChanMessage
