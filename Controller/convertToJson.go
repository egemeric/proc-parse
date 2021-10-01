package Controller

import "encoding/json"

type JsonMapString struct {
	Data map[string][]string
}
type JsonMapUint struct {
	Data map[string][]uint
}

func (param *JsonMapString) ConvertToJson() []byte {
	jsonString, err := json.Marshal(param.Data)
	Check(err)
	return jsonString

}
func (param *JsonMapUint) ConvertToJson() []byte {
	jsonString, err := json.Marshal(param.Data)
	Check(err)
	return jsonString

}
