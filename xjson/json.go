package xjson

import "encoding/json"

func Unmarshal(input []byte, data interface{}) error {
	return json.Unmarshal(input, &data)
}

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}
