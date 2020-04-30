package easycodefgo

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
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

// RSA 암호화
func EncryptRSA(text, publicKey string) (string, error) {
	pub, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", err
	}

	pubI, err := x509.ParsePKIXPublicKey(pub)
	if err != nil {
		return "", err
	}

	key := pubI.(*rsa.PublicKey)
	data, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(text))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}
