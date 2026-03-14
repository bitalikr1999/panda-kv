package int

import "fmt"

func IntEncoder(val int) []byte {
	return []byte(fmt.Sprintf(":%d\r\n", val))
}
