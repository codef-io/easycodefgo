package easycodefgo

import (
	"fmt"
	"testing"

	"github.com/dc7303/easycodefgo/message"

	"github.com/stretchr/testify/assert"
)

// newResponseByMessage 테스트
func TestNewResponseByMessage(t *testing.T) {
	ast := assert.New(t)

	res := newResponseByMessage(message.NotFound)
	data := res.GetData()
	ast.NotNil(data)

	result := res.GetResult()
	ast.NotNil(result)
}

// newResponseByMap 테스트
func TestNewResponseByMap(t *testing.T) {
	ast := assert.New(t)

	m := map[string]interface{}{
		Result: make(map[string]interface{}),
		Data:   make([]map[string]interface{}, 1),
	}
	res := newResponseByMap(m)
	ast.NotNil(res)
}

// Response.GetMessageInfo 테스트
func TestResponse_GetMessageInfo(t *testing.T) {
	ast := assert.New(t)

	res := newResponseByMessage(message.NotFound)
	code, msg, extraMsg := res.GetMessageInfo()
	ast.Equal(message.NotFound.Code, code)
	ast.Equal(message.NotFound.Message, msg)
	ast.Equal(message.NotFound.ExtraMessage, extraMsg)
}

// Response 인스턴스를 JSON string으로 변환 테스트
func TestWriteValueAsString(t *testing.T) {
	ast := assert.New(t)

	res := newResponseByMessage(message.OK)
	str := res.WriteValueAsString()

	// 원하는 결과
	s := fmt.Sprintf(
		`{"data":{},"result":{"code":"%s","extraMessage":"%s","message":"%s"}}`,
		message.OK.Code,
		message.OK.ExtraMessage,
		message.OK.Message,
	)
	ast.Equal(str, s)
}
