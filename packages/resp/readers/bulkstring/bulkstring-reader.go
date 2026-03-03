package bulkstring

import (
	"bitalikr1999/panda-kv/packages/resp/data"
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ReadBulkString(r *bufio.Reader) (data.RespData, error) {

	lineCounts, err := readLineCounts(r)
	if err != nil {
		return data.Empty(), err
	}

	values := []string{}

	for i := 0; i < lineCounts; i++ {

		lineValue, err := readLine(r)
		if err != nil {
			return data.Empty(), err
		}

		values = append(values, lineValue)
		err = skipSeparators(r)
		if err != nil {
			return data.Empty(), err
		}
	}

	return data.New(data.RespDataBulkString, values), nil

}

func readLineCounts(r *bufio.Reader) (int, error) {

	countLine, _ := r.ReadString('\n')

	countLine = strings.TrimSuffix(countLine, "\r\n")
	count, err := strconv.Atoi(countLine)

	if err != nil {
		fmt.Println("Failed parse array count", err)
	}

	return count, nil
}

func readLine(r *bufio.Reader) (string, error) {

	lineLength, err := readLineLength(r)
	if err != nil {
		return "", err
	}

	value, err := readValue(r, lineLength)
	if err != nil {
		return "", errors.Join(errors.New("failed to read line length"), err)
	}

	return value, nil

}

func readLineLength(r *bufio.Reader) (int, error) {
	var lengthBuf []byte

	for {
		b, err := r.ReadByte()
		if err != nil {
			fmt.Println("Error while reading bye", err)
			return 0, errors.New("wrong command structure;")
		}

		if b == '$' {
			continue
		}

		if b == '\r' {
			_, err = r.ReadByte()
			if err != nil {
				return 0, err
			}
			break
		}

		lengthBuf = append(lengthBuf, b)
	}

	length, err := strconv.Atoi(string(lengthBuf))
	if err != nil {
		return 0, err
	}
	return length, nil
}

func readValue(r *bufio.Reader, length int) (string, error) {
	var valueBuf []byte

	for i := 0; i < length; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return "", err
		}
		if b == '\r' || b == '\n' {
			return "", errors.New("protocolViolation")
		}
		valueBuf = append(valueBuf, b)
	}

	return string(valueBuf), nil
}

func skipSeparators(r *bufio.Reader) error {

	returnErr := errors.New("failed to read separators")
	rValue, err := r.ReadByte()
	if err != nil {
		return errors.Join(returnErr, err)
	}

	if rValue != '\r' {
		return errors.New("failed contract, missig /r")
	}

	nValue, err := r.ReadByte()
	if err != nil {
		return errors.Join(returnErr, err)
	}

	if nValue != '\n' {
		return errors.New("failed contract, missing /n")
	}

	return nil
}
