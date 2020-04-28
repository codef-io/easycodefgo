package easycodefgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// checkClientInfo 테스트
func TestCheckClientInfo(t *testing.T) {
	ast := assert.New(t)
	// 샌드박스는 클라이언트 정보가 상수로 입력되어 있어 True
	b := checkClientInfo(StatusSandbox)
	ast.True(b)

	// 정식버전에는 클라이언트 정보 입력이 필요하다
	b = checkClientInfo(StatusProduct)
	ast.False(b)

	SetClientInfo("test", "test")
	b = checkClientInfo(StatusProduct)
	ast.True(b)

	// 초기화
	SetClientInfo("", "")
}
