package Controller

import "encoding/json"

func ConvertToJson(data map[string][]string) []byte {
	jsonString, err := json.Marshal(data)
	Check(err)
	return jsonString

}
