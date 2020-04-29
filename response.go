package easycodefgo

import (
	"encoding/json"

	msg "github.com/dc7303/easycodefgo/message"
)

type Response struct {
	Result map[string]interface{}
	Data   interface{}
}

func (self *Response) WriteValueAsString() string {
	m := map[string]interface{}{
		"Result": self.Result,
		"Data":   self.Data,
	}
	b, _ := json.Marshal(m)
	return string(b)
}

func NewResponse(message *msg.MessageConstant) *Response {
	result := map[string]interface{}{
		Code:         message.Code,
		Message:      message.Message,
		ExtraMessage: message.ExtraMessage,
	}
	return &Response{
		result,
		make(map[string]interface{}),
	}
}
