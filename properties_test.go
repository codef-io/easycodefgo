package easycodefgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// getCodefDomain 테스트
func TestGetCodefDomain(t *testing.T) {
	ast := assert.New(t)
	// default 테스트
	domain := getCodefDomain(StatusSandbox)
	ast.Contains(domain, SandboxDomain)

	// demo 테스트
	domain = getCodefDomain(StatusDemo)
	ast.Contains(domain, DemoDomain)

	// product 테스트
	domain = getCodefDomain(StatusProduct)
	ast.Contains(domain, APIDomain)
}

// getClinetSecret 테스트
func TestGetClientSecret(t *testing.T) {
	ast := assert.New(t)
	// default 테스트
	key := getClientSecret(StatusSandbox)
	ast.Contains(key, SandboxClientSecret)

	// demo 테스트
	key = getClientSecret(StatusDemo)
	ast.Contains(key, DemoClientSecret)

	// product 테스트
	key = getClientSecret(StatusProduct)
	ast.Contains(key, ClientSecret)
}
