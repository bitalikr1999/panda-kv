package bulkstring

import (
	"fmt"
	"strings"
)

func EncodeBulkstring(val string) ([]byte, error) {

	parts := strings.Fields(val)
	var sb strings.Builder

	_, err := fmt.Fprintf(&sb, "*%d\r\n", len(parts))
	if err != nil {
		return nil, err
	}

	for _, part := range parts {
		_, err = fmt.Fprintf(&sb, "$%d\r\n%s\r\n", len(part), part)
		if err != nil {
			return nil, err
		}
	}
	return []byte(sb.String()), nil
}
