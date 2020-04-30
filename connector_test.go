package easycodefgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// getReqInfoByServiceType 테스트
func TestGetReqInfoByServiceType(t *testing.T) {
	ast := assert.New(t)

	// 샌드박스
	domain, id, secret := getReqInfoByServiceType(StatusSandbox)
	ast.Contains(SandboxDomain, domain)
	ast.Contains(SandboxClientID, id)
	ast.Contains(SandboxClientSecret, secret)

	// 데모
	SetClientInfoForDemo("demoID", "demoSecret")
	domain, id, secret = getReqInfoByServiceType(StatusDemo)
	ast.Contains(DemoDomain, domain)
	ast.Contains(DemoClientID, id)
	ast.Contains(DemoClientSecret, secret)

	// 정식버전
	SetClientInfo("productID", "productSecret")
	domain, id, secret = getReqInfoByServiceType(StatusProduct)
	ast.Contains(APIDomain, domain)
	ast.Contains(ClientID, id)
	ast.Contains(ClientSecret, secret)

	// 초기화
	SetClientInfoForDemo("", "")
	SetClientInfo("", "")
}

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
	err := setToken(SandboxClientID, SandboxClientSecret, &AccessToken)
	ast.NoError(err)
	ast.NotContains("", AccessToken)

	// 초기화
	AccessToken = ""
}

// requestProduct 테스트
func __TestRequestProduct(t *testing.T) {
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

	password, err := EncryptRSA("password", publicKey)
	ast.NoError(err)
	accountList := make([]map[string]interface{}, 1)
	m := map[string]interface{}{
		"countryCode":  "KR",
		"businessType": "BK",
		"clientType":   "P",
		"organization": "0004",
		"loginType":    "0",
		"password":     password,
		"keyFile":      "",
		"derFile":      "",
	}

	accountList = append(accountList, m)
	// TODO: connectedID 발급 구현 후 테스트 추가 작성
	//requestProduct(SandboxDomain + PathAddAccount, AccessToken,

}
