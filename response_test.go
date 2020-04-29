package easycodefgo

import (
	"fmt"
	"testing"

	msg "github.com/dc7303/easycodefgo/message"
	"github.com/stretchr/testify/assert"
)

// Response 인스턴스를 JSON string으로 변환 테스트
func TestWriteValueAsString(t *testing.T) {
	ast := assert.New(t)

	res := NewResponse(msg.OK)
	str := res.WriteValueAsString()

	// 원하는 결과
	s := fmt.Sprintf(
		`{"Data":{},"Result":{"code":"%s","extraMessage":"%s","message":"%s"}}`,
		msg.OK.Code,
		msg.OK.ExtraMessage,
		msg.OK.Message,
	)
	ast.Contains(str, s)
}
