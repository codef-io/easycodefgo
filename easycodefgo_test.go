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

// 퍼블릭키 체크 테스트
func TestCheckPublicKey(t *testing.T) {
	ast := assert.New(t)

	b := checkPublicKey()
	ast.False(b)
	PublicKey = "test"
	b = checkPublicKey()
	ast.True(b)

	PublicKey = ""
}

// 2Way 키워드 존재 여부 확인
func TestCheckTwoWayKeyword(t *testing.T) {
	ast := assert.New(t)

	m := map[string]interface{}{
		"is2Way":     true,
		"twoWayInfo": "info",
	}
	b := checkTwoWayKeyword(m)
	ast.True(b)

	delete(m, "is2Way")
	b = checkTwoWayKeyword(m)
	ast.False(b)

	m["is2Way"] = true
	delete(m, "twoWayInfo")
	ast.False(b)
}
