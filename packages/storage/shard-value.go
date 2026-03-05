package storage

type ShardValue struct {
	Value   string
	Options any
}

type CommandExecuteResponse struct {
	Ok    bool
	Value ShardValue
	Error error
}
