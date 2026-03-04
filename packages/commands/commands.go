package commands

import (
	"bitalikr1999/panda-kv/packages/resp/data"
	"errors"
	"fmt"
	"strings"
)

type Command interface {
	GetKey() string
}

type None struct{}

func (None) GetKey() string {
	return "none"
}

func Create(respData data.RespData) (Command, error) {

	values := respData.Arguments()

	err := validateValues(values)
	if err != nil {
		return None{}, err
	}

	commandName, arguments := values[0], values[1:]

	fmt.Println(commandName)
	switch strings.ToLower(commandName) {
	case "get":
		{
			return createGetCommand(arguments)
		}
	case "set":
		{
			return createSetCommand(arguments)
		}
	case "delete":
		{
			return createDeleteCommand(arguments)
		}
	default:
		{
			return None{}, errors.New("not exist command")
		}
	}
}

func CanCreateCommand(respData data.RespData) bool {
	return respData.Is(data.RespDataBulkString)
}

func validateValues(values []string) error {
	if len(values) < 2 {
		return errors.New("command not valid")
	}
	return nil
}
