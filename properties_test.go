package easycodefgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// getCodefDomain 테스트
func TestGetCodefDomain(t *testing.T) {
	ast := assert.New(t)
	// default 테스트
	domain := getCodefDomain(TypeSandbox)
	ast.Equal(domain, SandboxDomain)

	// demo 테스트
	domain = getCodefDomain(TypeDemo)
	ast.Equal(domain, DemoDomain)

	// product 테스트
	domain = getCodefDomain(TypeProduct)
	ast.Equal(domain, APIDomain)
}
