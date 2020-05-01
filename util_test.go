package easycodefgo

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 모든 공백 제거 테스트
func TestTrimAll(t *testing.T) {
	ast := assert.New(t)
	str := " hello world "
	r := TrimAll(str)
	ast.Equal(r, "helloworld")
}

// base64 인코딩 테스트
func TestEncodeToFileString(t *testing.T) {
	ast := assert.New(t)

	// 테스트 파일 생성
	mockFile := "mock_for_encode.txt"
	err := createMockFile(mockFile, []byte("hello world"))
	ast.NoError(err)

	// EncodeToFileString 호출
	result, err := EncodeToFileString(mockFile)
	ast.NoError(err)
	ast.Equal("aGVsbG8gd29ybGQ=", result)

	// 테스트 파일 삭제
	err = os.Remove(mockFile)
	ast.NoError(err)
}

// RSA 암호화 테스트
func TestEncryptRSA(t *testing.T) {
	ast := assert.New(t)
	publicKey := "MIIBIjANBgkqhkiG9w0BAQ" +
		"EFAAOCAQ8AMIIBCgKCAQEAuhRrVDeMf" +
		"b2fBaf8WmtGcQ23Cie+qDQqnkKG9eZV" +
		"yJdEvP1rLca+0CUOuAnpE8yGPY3HEbd" +
		"xKTsbIxxV9H8DCEMntXq2VP4loQoYUl" +
		"0h9dTjtBVWvhYev0s7N5B8Qu9LtykE2" +
		"k9KBuSZ+5dXulnHYdYjBaifZL6pzoD1" +
		"ckXoa4TtIuPjZZGXzr3Ivt5LDxPoPfw" +
		"1qMdqWRF9/YQSK1jZYa7PNR1Hbd8KB8" +
		"85VEcXNRU7ADHSgdYRBYB8apsPwaChy" +
		"jgrV98ATLOD7Dl4RlPtXcx/vEKjVMdt" +
		"CqJ2IHKeJoUCzBPY59U/mtIhjPuQmwS" +
		"MLEnLisDWEZMkenO0xJbwOwIDAQAB"
	result, err := EncryptRSA("hello world", publicKey)
	ast.NoError(err)
	ast.NotEmpty(result)
}

// 테스트 파일 생성
func createMockFile(filePath string, data []byte) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
