package xjson

import json "github.com/json-iterator/go"

func Unmarshal(input []byte, data interface{}) error {
	var conf = json.ConfigCompatibleWithStandardLibrary
	return conf.Unmarshal(input, &data)
}

func Marshal(data interface{}) ([]byte, error) {
	var conf = json.ConfigCompatibleWithStandardLibrary
	return conf.Marshal(&data)
}
