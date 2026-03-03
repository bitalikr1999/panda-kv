package data

type RespDataType int

const (
	RespDataString RespDataType = iota
	RespDataNumber
	RespDataBulkString
)

type RespData struct {
	dataType  RespDataType
	arguments []string
}

func New(dataType RespDataType, arguments []string) RespData {

	return RespData{
		dataType:  dataType,
		arguments: arguments,
	}
}

func Empty() RespData {
	return RespData{}
}

func (data RespData) Is(dataType RespDataType) bool {
	return data.dataType == dataType
}

func (data RespData) Arguments() []string {
	return data.arguments
}
