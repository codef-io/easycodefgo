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
	ast.Equal(key, codef.demoClientSecret)

	// product 테스트
	key = codef.getClientSecret(TypeProduct)
	ast.Equal(key, codef.clientSecret)
}

// 2Way 키워드 존재 여부 확인 테스트
func TestIsEmptyTwoWayKeyword(t *testing.T) {
	ast := assert.New(t)

	m := map[string]interface{}{
		"is2Way":     true,
		"twoWayInfo": "info",
	}
	b := isEmptyTwoWayKeyword(m)
	ast.False(b)

	delete(m, "is2Way")
	b = isEmptyTwoWayKeyword(m)
	ast.False(b)

	delete(m, "twoWayInfo")
	b = isEmptyTwoWayKeyword(m)
	ast.True(b)
}

// RequestProduct 샌드박스로 테스트
func TestRequestProductBySandbox(t *testing.T) {
	ast := assert.New(t)

	codef := &Codef{}
	codef.PublicKey = "public_key"

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
	ast.Equal(SandboxDomain, reqInfo.domain)
	ast.Equal(SandboxClientID, reqInfo.clientID)
	ast.Equal(SandboxClientSecret, reqInfo.clientSecret)

	// 데모
	codef.SetClientInfoForDemo("demoID", "demoSecret")
	reqInfo = codef.getReqInfoByServiceType(TypeDemo)
	ast.Equal(DemoDomain, reqInfo.domain)
	ast.Equal(codef.demoClientID, reqInfo.clientID)
	ast.Equal(codef.demoClientSecret, reqInfo.clientSecret)

	// 정식버전
	codef.SetClientInfo("productID", "productSecret")
	reqInfo = codef.getReqInfoByServiceType(TypeProduct)
	ast.Equal(APIDomain, reqInfo.domain)
	ast.Equal(codef.clientID, reqInfo.clientID)
	ast.Equal(codef.clientSecret, reqInfo.clientSecret)
}

// hasTwoWayInfo 테스트
func TestHasTwoWayInfo(t *testing.T) {
	ast := assert.New(t)

	twoWayInfo := map[string]interface{}{
		"jobIndex":        "1",
		"threadIndex":     "1",
		"jti":             "",
		"twoWayTimestamp": "",
	}
	param := map[string]interface{}{
		"is2Way":     "string",
		"twoWayInfo": twoWayInfo,
	}

	// bool 타입이 아닐때
	ok := hasTwoWayInfo(param)
	ast.False(ok)

	// is2Way가 false일 때
	param["is2Way"] = false
	ok = hasTwoWayInfo(param)
	ast.False(ok)

	// 정상 작동
	param["is2Way"] = true
	ok = hasTwoWayInfo(param)
	ast.True(ok)

	// twoWayInfo가 맵으로 캐스팅 되지 않을 때
	param["twoWayInfo"] = "string"
	ok = hasTwoWayInfo(param)
	ast.False(ok)

	// twoWayInfo가 nil일 때
	param["twoWayInfo"] = nil
	ok = hasTwoWayInfo(param)
	ast.False(ok)
}

// checkNeedValueInTwoWayInfo 테스트
func TestCheckNeedValueInTwoWayInfo(t *testing.T) {
	ast := assert.New(t)

	twoWayInfo := map[string]interface{}{
		"jobIndex":        "1",
		"threadIndex":     "1",
		"jti":             "",
		"twoWayTimestamp": "",
	}
	// 정상 작동
	ok := checkNeedValueInTwoWayInfo(twoWayInfo)
	ast.True(ok)

	// jobIndex가 없을 때
	delete(twoWayInfo, "jobIndex")
	ok = checkNeedValueInTwoWayInfo(twoWayInfo)
	ast.False(ok)
}

// ReuqestToken 요청 클라이언트 정보가 입력되지 않았을 때 에러 처리
func TestRequestTokenByEmptyClientInfo(t *testing.T) {
	ast := assert.New(t)

	codef := &Codef{}
	result, err := codef.RequestToken(TypeDemo)
	ast.Error(err)
	ast.Empty(result)
}

// 서비스 타입별 API 요청 관리가 가능한지 확인을 위한 테스트
// 비동기 상황에서 서비스 타입을 혼합해서 사용할 경우 에러 방지를 위함
func TestSetAccessToken(t *testing.T) {
	ast := assert.New(t)

	codef := &Codef{}
	token, _ := codef.RequestToken(TypeSandbox)
	ast.NotEmpty(token)

	codef.SetAccessToken(token, TypeSandbox)
	ast.NotEmpty(codef.getAccessToken(TypeSandbox))
	codef.SetAccessToken(token, TypeDemo)
	ast.NotEmpty(codef.getAccessToken(TypeDemo))
	codef.SetAccessToken(token, TypeProduct)
	ast.NotEmpty(codef.getAccessToken(TypeProduct))
}

// 퍼블릭키 체크 테스트
func TestEmptyPublicKeyError(t *testing.T) {
	ast := assert.New(t)
	codef := &Codef{}
	accountList := []map[string]interface{}{
		map[string]interface{}{
			"countryCode":  "KR",
			"businessType": "BK",
			"clientType":   "P",
			"organization": "0004",
			"loginType":    "1",
			"id":           "user_id",
		},
	}

	res, err := codef.CreateAccount(TypeSandbox, map[string]interface{}{
		"accountList": accountList,
	})
	ast.NoError(err)

	data := map[string]interface{}{}
	json.Unmarshal([]byte(res), &data)
	ast.Equal((data["result"].(map[string]interface{}))["code"], "CF-00015")
}
