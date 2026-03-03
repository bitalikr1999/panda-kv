package bulkstring

import (
	"bitalikr1999/panda-kv/packages/resp/data"
	"bufio"
)

type Bulkstring struct{}

func (Bulkstring) Read(r *bufio.Reader) (data.RespData, error) {
	return ReadBulkString(r)
}
