package commands

type GetCommand struct {
	Key string
}

func createGetCommand(params []string) (GetCommand, error) {

	key := params[0]

	return GetCommand{
		Key: key,
	}, nil
}
