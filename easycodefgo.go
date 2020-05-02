package easycodefgo

import msg "github.com/dc7303/easycodefgo/message"

type EasyCodef struct {
}

// 상품 요청
func RequestProduct(
	productURL string,
	serviceType ServiceStatus,
	param map[string]interface{},
) (string, error) {
	validFlag := true
	// 클라이언트 정보 체크
	validFlag = checkClientInfo(serviceType)
	if !validFlag {
		res := newResponseByMessage(msg.EmptyClientInfo)
		return res.WriteValueAsString(), nil
	}

	// 퍼블릭 키 체크
	validFlag = checkPublicKey()
	if !validFlag {
		res := newResponseByMessage(msg.EmptyPublicKey)
		return res.WriteValueAsString(), nil
	}

	// 추가인증 키워드 체크
	validFlag = checkTwoWayKeyword(param)
	if !validFlag {
		res := newResponseByMessage(msg.Invalid2WayKeyword)
		return res.WriteValueAsString(), nil
	}

	res, err := execute(productURL, param)
	if err != nil {
		return "", err
	}
	return res.WriteValueAsString(), nil
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
	}
	return true
}

// 퍼블릭키 정보 설정 확인
func checkPublicKey() bool {
	if TrimAll(PublicKey) == "" {
		return false
	}
	return true
}

// 2Way 키워드 존재 여부 확인
func checkTwoWayKeyword(param map[string]interface{}) bool {
	if _, ok := param["is2Way"]; !ok {
		return false
	}
	if _, ok := param["twoWayInfo"]; !ok {
		return false
	}
	return true
}

func CreateAccount(serviceType ServiceStatus, param map[string]interface{}) (string, error) {
	return RequestProduct(PathCreateAccount, serviceType, param)
}

func AddAccount(serviceType ServiceStatus, param map[string]interface{}) (string, error) {
	return RequestProduct(PathAddAccount, serviceType, param)
}
