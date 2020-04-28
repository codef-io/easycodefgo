package easycodefgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// getCodefDomain 테스트
func TestGetCodefDomain(t *testing.T) {
	ast := assert.New(t)
	// default 테스트
	domain := getCodefDomain()
	ast.Contains(domain, SandboxDomain)

	// product 테스트
	ServiceType = StatusProduct
	domain = getCodefDomain()
	ast.Contains(domain, APIDomain)

	// 다른 테스트에 영향을 주지 않기 위해 초기화
	ServiceType = StatusSandbox
}

// getClinetSecret 테스트
func TestGetClientSecret(t *testing.T) {
	ast := assert.New(t)
	// default 테스트
	key := getClientSecret()
	ast.Contains(key, SandboxClientSecret)

	// product 테스트
	ServiceType = StatusProduct
	key = getCodefDomain()
	ast.Contains(key, ClientSecret)

	// 다른 테스트에 영향을 주지 않기 위해 초기화
	ServiceType = StatusSandbox
}
