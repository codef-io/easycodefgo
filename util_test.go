package easycodefgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 모든 공백 제거 테스트
func TestTrimAll(t *testing.T) {
	ast := assert.New(t)
	str := " hello world "
	r := TrimAll(str)
	ast.Contains(r, "helloworld")
}
