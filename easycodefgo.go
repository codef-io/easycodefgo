package easycodefgo

import msg "github.com/dc7303/easycodefgo/message"

// 상품 요청
func RequestProduct(
	productURL string,
	serviceType ServiceType,
	param map[string]interface{},
) (string, error) {
	validFlag := true
	// 클라이언트 정보 체크
	validFlag = checkClientInfo(serviceType)
	if !validFlag {
		res := newResponseByMessage(msg.EmptyClientInfo)
		return res.WriteValueAsString(), nil
	}

	// 추가인증 키워드 체크
	validFlag = checkTwoWayKeyword(param)
	if !validFlag {
		res := newResponseByMessage(msg.Invalid2WayKeyword)
		return res.WriteValueAsString(), nil
	}

	res, err := execute(productURL, param, serviceType)
	if err != nil {
		return "", err
	}
	return res.WriteValueAsString(), nil
}

// 클라이언트 정보 검사
func checkClientInfo(serviceType ServiceType) bool {
	switch serviceType {
	case TypeProduct:
		if TrimAll(ClientID) == "" || TrimAll(ClientSecret) == "" {
			return false
		}
	case TypeDemo:
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

// connectedID 발급을 위한 계정 등록
func CreateAccount(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return RequestProduct(PathCreateAccount, serviceType, param)
}

// 게정 정보 추가
func AddAccount(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return RequestProduct(PathAddAccount, serviceType, param)
}

// 계정 정보 수정
func UpdateAccount(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return RequestProduct(PathUpdateAccount, serviceType, param)
}

// 계정 정보 삭제
func DeleteAccount(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return RequestProduct(PathDeleteAccount, serviceType, param)
}

// connectedID로 등록된 계정 목록 조회
func GetAccountList(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return RequestProduct(PathGetAccountList, serviceType, param)
}

// 클라이언트 정보로 등록된 모든 connectedID 목록 조회
func GetConnectedIDList(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return RequestProduct(PathGetCIDList, serviceType, param)
}

// 토큰 발급
func RequestToken(serviceType ServiceType) (map[string]interface{}, error) {
	switch serviceType {
	case TypeProduct:
		return requestToken(ClientID, ClientSecret)
	case TypeDemo:
		return requestToken(DemoClientID, DemoClientSecret)
	default:
		return requestToken(SandboxClientID, SandboxClientSecret)
	}
}
