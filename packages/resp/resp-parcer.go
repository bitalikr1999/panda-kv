package resp

import (
	"bitalikr1999/panda-kv/packages/resp/data"
	"bitalikr1999/panda-kv/packages/resp/readers"
	"bitalikr1999/panda-kv/packages/resp/readers/bulkstring"
	"bufio"
)

var parsers = map[string]readers.RespReader{
	"*": bulkstring.Bulkstring{},
}

func ParseResp(r *bufio.Reader) (data.RespData, error) {

	b, err := r.ReadByte()

	if err != nil {
		return data.Empty(), err
	}

	parser := parsers[string(b)]

	if parser == nil {
		return data.Empty(), nil
	}

	return parser.Read(r)
}
