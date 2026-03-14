package string_encoder

func EncodeString(val string) []byte {
	return []byte("+" + val + "\r\n")
}
