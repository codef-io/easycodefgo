package easycodefgo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// checkClientInfo 테스트
func TestCheckClientInfo(t *testing.T) {
	ast := assert.New(t)
	// 샌드박스는 클라이언트 정보가 상수로 입력되어 있어 True
	b := checkClientInfo(TypeSandbox)
	ast.True(b)

	// 정식버전에는 클라이언트 정보 입력이 필요하다
	b = checkClientInfo(TypeProduct)
	ast.False(b)

	SetClientInfo("test", "test")
	b = checkClientInfo(TypeProduct)
	ast.True(b)

	// 초기화
	SetClientInfo("", "")
}

// 2Way 키워드 존재 여부 확인
func TestCheckTwoWayKeyword(t *testing.T) {
	ast := assert.New(t)

	m := map[string]interface{}{
		"is2Way":     true,
		"twoWayInfo": "info",
	}
	b := checkTwoWayKeyword(m)
	ast.False(b)

	delete(m, "is2Way")
	b = checkTwoWayKeyword(m)
	ast.False(b)

	delete(m, "twoWayInfo")
	b = checkTwoWayKeyword(m)
	ast.True(b)
}

// RequestProduct 샌드박스로 테스트
func TestRequestProductBySandbox(t *testing.T) {
	ast := assert.New(t)

	// 테스트 데이터 생성
	param, err := createParamForCreateConnectedID()
	ast.NoError(err)

	// CreateAccount 요청
	result, err := RequestProduct(PathCreateAccount, TypeSandbox, param)
	ast.NoError(err)

	// 결과 맵으로 변환
	resultMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(result), &resultMap)
	ast.NoError(err)

	// data 존재 여부
	data, ok := resultMap[KeyData]
	ast.True(ok)
	ast.NotEmpty(data)

	// ConnectedID 존재 여부
	castData := data.(map[string]interface{})
	cid, ok := castData[KeyConnectedID]
	ast.True(ok)
	ast.NotEmpty(cid)
}
