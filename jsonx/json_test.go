package jsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Resp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

var jsonByte = []byte(`{"code":"unauthenticated","msg":"unauthenticated"}`)
var jsonStruct = Resp{Code: "unauthenticated", Msg: "unauthenticated"}

func TestUnmarshal(t *testing.T) {
	var resp Resp
	err := Unmarshal(jsonByte, &resp)
	assert.Nil(t, err)
	assert.Equal(t, resp, jsonStruct)
}

func TestMarshal(t *testing.T) {
	marshal, err := Marshal(jsonStruct)
	assert.Nil(t, err)
	assert.Equal(t, string(marshal), string(jsonByte))
}
