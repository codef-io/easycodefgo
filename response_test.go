package easycodefgo

import (
	"fmt"
	"testing"

	"github.com/dc7303/easycodefgo/message"

	"github.com/stretchr/testify/assert"
)

// Response 인스턴스를 JSON string으로 변환 테스트
func TestWriteValueAsString(t *testing.T) {
	ast := assert.New(t)

	res := NewResponse(message.OK)
	str := res.WriteValueAsString()

	// 원하는 결과
	s := fmt.Sprintf(
		`{"Data":{},"Result":{"code":"%s","extraMessage":"%s","message":"%s"}}`,
		message.OK.Code,
		message.OK.ExtraMessage,
		message.OK.Message,
	)
	ast.Contains(str, s)
}
