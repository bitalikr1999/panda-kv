package commands

import "errors"

type SetCommand struct {
	Key   string
	Value string
}

func createSetCommand(params []string) (SetCommand, error) {

	key := params[0]
	value := params[1]

	if key == "" || value == "" {
		return SetCommand{}, errors.New("not valid set command")
	}

	return SetCommand{
		Key:   key,
		Value: value,
	}, nil
}
