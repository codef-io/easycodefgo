package easycodefgo

import (
	"encoding/json"

	msg "github.com/dc7303/easycodefgo/message"
)

type Response map[string]interface{}

func (self *Response) GetData() interface{} {
	return (*self)[Data]
}

func (self *Response) GetResult() map[string]interface{} {
	return (*self)[Result].(map[string]interface{})
}

func (self *Response) GetMessageInfo() (string, string, string) {
	r := self.GetResult()
	return r[Code].(string), r[Message].(string), r[ExtraMessage].(string)
}

func (self *Response) WriteValueAsString() string {
	b, _ := json.Marshal(self)
	return string(b)
}

// message 정보로 Response 생성 메소드
func newResponseByMessage(message *msg.MessageConstant) *Response {
	result := map[string]interface{}{
		Code:         message.Code,
		Message:      message.Message,
		ExtraMessage: message.ExtraMessage,
	}
	return &Response{
		Result: result,
		Data:   make(map[string]interface{}),
	}
}

// Response 생성 메소드
// 코드에프 결과를 생성하기 위한 메소드
func newResponseByMap(m map[string]interface{}) *Response {
	res := Response(m)
	return &res
}
