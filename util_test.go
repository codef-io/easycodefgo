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
	ast.Contains(r, "helloworld")
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
	ast.Contains("aGVsbG8gd29ybGQ=", result)

	// 테스트 파일 삭제
	err = os.Remove(mockFile)
	ast.NoError(err)
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
