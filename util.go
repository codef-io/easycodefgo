package easycodefgo

import "strings"

// 모든 공백 제거
func TrimAll(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	return str
}
