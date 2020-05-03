package easycodefgo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// checkClientInfo 테스트
func TestCheckClientInfo(t *testing.T) {
	ast := assert.New(t)

	codef := &Codef{}
	// 샌드박스는 클라이언트 정보가 상수로 입력되어 있어 True
	b := codef.checkClientInfo(TypeSandbox)
	ast.True(b)

	// 정식버전에는 클라이언트 정보 입력이 필요하다
	b = codef.checkClientInfo(TypeProduct)
	ast.False(b)

	codef.SetClientInfo("test", "test")
	b = codef.checkClientInfo(TypeProduct)
	ast.True(b)
}

// getClinetSecret 테스트
func TestGetClientSecret(t *testing.T) {
	ast := assert.New(t)

	codef := &Codef{}
	// default 테스트
	key := codef.getClientSecret(TypeSandbox)
	ast.Equal(key, SandboxClientSecret)

	// demo 테스트
	key = codef.getClientSecret(TypeDemo)
	ast.Equal(key, codef.DemoClientSecret)

	// product 테스트
	key = codef.getClientSecret(TypeProduct)
	ast.Equal(key, codef.ClientSecret)
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

	codef := &Codef{}
	// 테스트 데이터 생성
	param, err := createParamForCreateConnectedID()
	ast.NoError(err)

	// CreateAccount 요청
	result, err := codef.RequestProduct(PathCreateAccount, TypeSandbox, param)
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

	// param == nil 로 요청
	result, err = codef.RequestProduct(PathGetCIDList, TypeSandbox, nil)
	ast.NoError(err)

	// 결과 맵으로 변환
	resultMap = make(map[string]interface{})
	err = json.Unmarshal([]byte(result), &resultMap)
	ast.NoError(err)

	// data 존재 여부
	data, ok = resultMap[KeyData]
	ast.True(ok)
	ast.NotEmpty(data)
}

// getReqInfoByServiceType 테스트
func TestGetReqInfoByServiceType(t *testing.T) {
	ast := assert.New(t)

	codef := &Codef{}
	// 샌드박스
	reqInfo := codef.getReqInfoByServiceType(TypeSandbox)
	ast.Equal(SandboxDomain, reqInfo.Domain)
	ast.Equal(SandboxClientID, reqInfo.ClientID)
	ast.Equal(SandboxClientSecret, reqInfo.ClientSecret)

	// 데모
	codef.SetClientInfoForDemo("demoID", "demoSecret")
	reqInfo = codef.getReqInfoByServiceType(TypeDemo)
	ast.Equal(DemoDomain, reqInfo.Domain)
	ast.Equal(codef.DemoClientID, reqInfo.ClientID)
	ast.Equal(codef.DemoClientSecret, reqInfo.ClientSecret)

	// 정식버전
	codef.SetClientInfo("productID", "productSecret")
	reqInfo = codef.getReqInfoByServiceType(TypeProduct)
	ast.Equal(APIDomain, reqInfo.Domain)
	ast.Equal(codef.ClientID, reqInfo.ClientID)
	ast.Equal(codef.ClientSecret, reqInfo.ClientSecret)
}
