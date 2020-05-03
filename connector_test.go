package easycodefgo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 액세스 토큰 요청 테스트
func TestRequestToken(t *testing.T) {
	ast := assert.New(t)
	m, err := requestToken(SandboxClientID, SandboxClientSecret)
	ast.NoError(err)
	ast.NotNil(m)

	token, ok := m["access_token"]
	ast.True(ok)
	ast.NotNil(token)
	ast.IsType("", token)
}

// 토큰 셋팅 테스트
func TestSetToken(t *testing.T) {
	ast := assert.New(t)
	codef := &Codef{}

	err := setToken(SandboxClientID, SandboxClientSecret, &codef.AccessToken)
	ast.NoError(err)
	ast.NotEmpty(codef.AccessToken)
}

// requestProduct 테스트
func TestRequestProduct(t *testing.T) {
	ast := assert.New(t)
	param, err := createParamForCreateConnectedID()
	ast.NoError(err)

	data, err := json.Marshal(param)
	ast.NoError(err)

	accessToken := ""
	// 404 에러 발생 테스트
	res, err := requestProduct(SandboxDomain+"/failPath", accessToken, string(data))
	ast.NoError(err)
	code, _, _ := res.GetMessageInfo()
	ast.Equal("CF-00404", code)
	ast.Empty(res.GetData())

	// Connected ID 정상 발급 테스트
	err = setToken(SandboxClientID, SandboxClientSecret, &accessToken)
	ast.NoError(err)

	res, err = requestProduct(SandboxDomain+PathCreateAccount, accessToken, string(data))
	testExistConnectedID(ast, res)
}

// execute 테스트
func TestExecute(t *testing.T) {
	ast := assert.New(t)
	codef := &Codef{}

	param, err := createParamForCreateConnectedID()
	ast.NoError(err)

	reqInfo := codef.getReqInfoByServiceType(TypeSandbox)
	res, err := execute(PathCreateAccount, param, &codef.AccessToken, reqInfo)
	ast.NoError(err)

	testExistConnectedID(ast, res)
}

// connectedID 발급 성공 여부 테스트
func testExistConnectedID(ast *assert.Assertions, res *Response) {
	code, _, _ := res.GetMessageInfo()
	ast.Equal("CF-00000", code)
	result, ok := res.GetData().(map[string]interface{})
	ast.True(ok)
	_, ok = result["connectedId"]
	ast.True(ok)
}

// ConnectedID 발급을 위한 Body 테스트 데이터 생성
func createParamForCreateConnectedID() (map[string]interface{}, error) {
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

	password, err := EncryptRSA("password", publicKey)
	if err != nil {
		return nil, err
	}
	accountList := []map[string]interface{}{}
	m := map[string]interface{}{
		"countryCode":  "KR",
		"businessType": "BK",
		"clientType":   "P",
		"organization": "0004",
		"loginType":    "1",
		"id":           "testID",
		"password":     password,
	}

	accountList = append(accountList, m)
	param := map[string]interface{}{
		"accountList": accountList,
	}

	return param, nil
}
