package easycodefgo

import (
	"encoding/json"
)

type Response map[string]interface{}

func (self *Response) GetData() interface{} {
	return (*self)[KeyData]
}

func (self *Response) GetResult() map[string]interface{} {
	return (*self)[KeyResult].(map[string]interface{})
}

func (self *Response) GetMessageInfo() (string, string, string) {
	r := self.GetResult()
	return r[KeyCode].(string), r[KeyMessage].(string), r[KeyExtraMessage].(string)
}

func (self *Response) WriteValueAsString() string {
	b, _ := json.Marshal(self)
	return string(b)
}

// message 정보로 Response 생성 메소드
func newResponseByMessage(message *messageConstant) *Response {
	result := map[string]interface{}{
		KeyCode:         message.Code,
		KeyMessage:      message.Message,
		KeyExtraMessage: message.ExtraMessage,
	}
	return &Response{
		KeyResult: result,
		KeyData:   make(map[string]interface{}),
	}
}

// Response 생성 메소드
// 코드에프 결과를 생성하기 위한 메소드
func newResponseByMap(m map[string]interface{}) *Response {
	res := Response(m)
	return &res
}
