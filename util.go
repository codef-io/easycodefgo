package easycodefgo

import (
	"bufio"
	"encoding/base64"
	"os"
	"strings"
)

// 모든 공백 제거
func TrimAll(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	return str
}

// 파일 정보를 base64 문자열로 인코딩
func EncodeToFileString(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	var fSize int64 = fInfo.Size()
	buf := make([]byte, fSize)

	reader := bufio.NewReader(file)
	_, err = reader.Read(buf)
	if err != nil {
		return "", err
	}

	base64str := base64.StdEncoding.EncodeToString(buf)
	return base64str, nil
}
