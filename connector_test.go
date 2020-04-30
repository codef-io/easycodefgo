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
}
