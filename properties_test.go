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
	ast.Equal(domain, SandboxDomain)

	// demo 테스트
	domain = getCodefDomain(StatusDemo)
	ast.Equal(domain, DemoDomain)

	// product 테스트
	domain = getCodefDomain(StatusProduct)
	ast.Equal(domain, APIDomain)
}

// getClinetSecret 테스트
func TestGetClientSecret(t *testing.T) {
	ast := assert.New(t)
	// default 테스트
	key := getClientSecret(StatusSandbox)
	ast.Equal(key, SandboxClientSecret)

	// demo 테스트
	key = getClientSecret(StatusDemo)
	ast.Equal(key, DemoClientSecret)

	// product 테스트
	key = getClientSecret(StatusProduct)
	ast.Equal(key, ClientSecret)
}
