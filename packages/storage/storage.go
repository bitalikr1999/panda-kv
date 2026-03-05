package storage

import (
	"bitalikr1999/panda-kv/packages/commands"
	"fmt"
	"hash/fnv"
)

type Storage struct {
	shards []Shard
}

func New(shardsCount int) *Storage {

	shards := make([]Shard, shardsCount)

	for i := 0; i < shardsCount; i++ {
		shards[i] = *NewShard()

		go shards[i].Start()
	}

	return &Storage{
		shards: shards,
	}

}

func (s *Storage) Send(command commands.Command) CommandExecuteResponse {

	key := command.GetKey()
	shard := s.getShard(key)

	fmt.Println("We found shaard!", shard)

	return shard.Send(command)
}

func (s *Storage) getShard(key string) *Shard {
	h := fnv.New32a()
	h.Write([]byte(key))
	return &s.shards[h.Sum32()%uint32(len(s.shards))]
}

func (s *Storage) Close() {
	for _, shard := range s.shards {
		shard.Stop()
	}
}
