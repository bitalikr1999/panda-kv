package storage

type ChanMessage struct {
	responseChan chan interface{}
}

type StorageInChannel = chan ChanMessage
