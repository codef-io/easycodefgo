package easycodefgo

type EasyCodef struct {
}

// 상품 요청
func RequestProduct(
	productURL string,
	serviceType ServiceCode,
	parameter map[string]string,
) string {
	ProductURL = productURL
	ServiceType = serviceType

	validFlag := true
	validFlag = checkClientInfo(ServiceType)
	if !validFlag {
		return "false"
	}

	return "he"
}

// 클라이언트 정보 검사
func checkClientInfo(serviceType ServiceCode) bool {
	switch serviceType {
	case StatusProduct:
		if TrimAll(ClientID) == "" || TrimAll(ClientSecret) == "" {
			return false
		}
	case StatusDemo:
		if TrimAll(DemoClientSecret) == "" || TrimAll(DemoClientSecret) == "" {
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
