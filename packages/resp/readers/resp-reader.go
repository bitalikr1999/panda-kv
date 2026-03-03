package readers

import (
	"bitalikr1999/panda-kv/packages/resp/data"
	"bufio"
)

type RespReader interface {
	Read(r *bufio.Reader) (data.RespData, error)
}
