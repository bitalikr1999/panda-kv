package commands

import "errors"

type DeleteCommand struct {
	Key string
}

func createDeleteCommand(params []string) (DeleteCommand, error) {

	key := params[0]
	if key == "" {
		return DeleteCommand{}, errors.New("not valid delete command")
	}

	return DeleteCommand{Key: key}, nil
}

func (c DeleteCommand) GetKey() string {
	return c.Key
}
