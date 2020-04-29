package easycodefgo

import msg "github.com/dc7303/easycodefgo/message"

type EasyCodef struct {
}

// 상품 요청
func RequestProduct(
	productURL string,
	serviceType ServiceStatus,
	parameter map[string]interface{},
) string {
	ProductURL = productURL
	ServiceType = serviceType

	validFlag := true
	// 클라이언트 정보 체크
	validFlag = checkClientInfo(ServiceType)
	if !validFlag {
		res := NewResponse(msg.EmptyClientInfo)
		return res.WriteValueAsString()
	}

	// 퍼블릭 키 체크
	validFlag = checkPublicKey()
	if !validFlag {
		res := NewResponse(msg.EmptyPublicKey)
		return res.WriteValueAsString()
	}

	return ""
}

// 클라이언트 정보 검사
func checkClientInfo(serviceType ServiceStatus) bool {
	switch serviceType {
	case StatusProduct:
		if TrimAll(ClientID) == "" || TrimAll(ClientSecret) == "" {
			return false
		}
	case StatusDemo:
		if TrimAll(DemoClientID) == "" || TrimAll(DemoClientSecret) == "" {
			return false
		}
	default:
		if TrimAll(SandboxClientID) == "" || TrimAll(SandboxClientSecret) == "" {
			return false
		}
		break
	}
	return true
}

func AddAccount(serviceType ServiceStatus, param map[string]interface{}) string {
	return RequestProduct(PathAddAccount, serviceType, param)
}

// 퍼블릭키 정보 설정 확인
func checkPublicKey() bool {
	if TrimAll(PublicKey) == "" {
		return false
	}
	return true
}
