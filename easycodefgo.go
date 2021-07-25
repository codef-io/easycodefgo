package easycodefgo

import (
	"errors"
	"reflect"
)

// 액세스 토큰 관리 구조체
type accessToken struct {
	product string
	demo    string
	sandbox string
}

// CODEF API
type Codef struct {
	accessToken      accessToken // OAUTH2.0 토큰
	demoClientID     string      // 데모 엑세스 토큰 밝브을 위한 클라이언트 아이디
	demoClientSecret string      // 데모 엑세스 토큰 밝브을 위한 클라이언트 시크릿
	clientID         string      // 정식 엑세스 토큰 발급을 위한 클라이언트 아이디
	clientSecret     string      // 정식 엑세스 토큰 발급을 위한 클라이언트 시크릿
	PublicKey        string      // 유저 퍼블릭키
}

// 요청 정보
type requestInfo struct {
	domain       string
	clientID     string
	clientSecret string
}

// 상품 요청
func (self *Codef) RequestProduct(
	productPath string,
	serviceType ServiceType,
	param map[string]interface{},
) (string, error) {

	// 클라이언트 정보 체크
	if !self.checkClientInfo(serviceType) {
		res := newResponseByMessage(messageEmptyClientInfo)
		return res.WriteValueAsString(), nil
	}

	// 퍼블릭키 정보 체크
	if TrimAll(self.PublicKey) == "" {
		res := newResponseByMessage(messageEmptyPublicKey)
		return res.WriteValueAsString(), nil
	}

	// 추가인증 키워드 체크
	if !isEmptyTwoWayKeyword(param) {
		res := newResponseByMessage(messageInvalid2WayKeyword)
		return res.WriteValueAsString(), nil
	}

	reqInfo := self.getReqInfoByServiceType(serviceType)

	res, err := execute(productPath, param, self.getAccessToken(serviceType), reqInfo)
	if err != nil {
		return "", err
	}
	return res.WriteValueAsString(), nil
}

// 상품 추가인증 요청
func (self *Codef) RequestCertification(
	productPath string,
	serviceType ServiceType,
	param map[string]interface{},
) (string, error) {

	// 클라이언트 정보 체크
	if !self.checkClientInfo(serviceType) {
		res := newResponseByMessage(messageEmptyClientInfo)
		return res.WriteValueAsString(), nil
	}

	// 퍼블릭키 정보 체크
	if TrimAll(self.PublicKey) == "" {
		res := newResponseByMessage(messageEmptyPublicKey)
		return res.WriteValueAsString(), nil
	}

	// 추가인증 파라미터 필수 입력 체크
	if !hasTwoWayInfo(param) {
		res := newResponseByMessage(messageInvalid2WayInfo)
		return res.WriteValueAsString(), nil
	}

	reqInfo := self.getReqInfoByServiceType(serviceType)

	// 상품 조회 요청
	res, err := execute(productPath, param, self.getAccessToken(serviceType), reqInfo)
	if err != nil {
		return "", err
	}

	return res.WriteValueAsString(), nil
}

// 클라이언트 정보 검사
func (self *Codef) checkClientInfo(serviceType ServiceType) bool {
	switch serviceType {
	case TypeProduct:
		if TrimAll(self.clientID) == "" || TrimAll(self.clientSecret) == "" {
			return false
		}
	case TypeDemo:
		if TrimAll(self.demoClientID) == "" || TrimAll(self.demoClientSecret) == "" {
			return false
		}
	default:
		if TrimAll(SandboxClientID) == "" || TrimAll(SandboxClientSecret) == "" {
			return false
		}
	}
	return true
}

// connectedID 발급을 위한 계정 등록
func (self *Codef) CreateAccount(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return self.RequestProduct(PathCreateAccount, serviceType, param)
}

// 게정 정보 추가
func (self *Codef) AddAccount(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return self.RequestProduct(PathAddAccount, serviceType, param)
}

// 계정 정보 수정
func (self *Codef) UpdateAccount(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return self.RequestProduct(PathUpdateAccount, serviceType, param)
}

// 계정 정보 삭제
func (self *Codef) DeleteAccount(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return self.RequestProduct(PathDeleteAccount, serviceType, param)
}

// connectedID로 등록된 계정 목록 조회
func (self *Codef) GetAccountList(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return self.RequestProduct(PathGetAccountList, serviceType, param)
}

// 클라이언트 정보로 등록된 모든 connectedID 목록 조회
func (self *Codef) GetConnectedIDList(serviceType ServiceType, param map[string]interface{}) (string, error) {
	return self.RequestProduct(PathGetCIDList, serviceType, param)
}

// 토큰 발급
func (self *Codef) RequestToken(serviceType ServiceType) (string, error) {
	existClientInfo := self.checkClientInfo(serviceType)
	if !existClientInfo {
		return "", errors.New("The ClientID and ClientSecret values ​​are empty. Please set the value according to the service type.")
	}
	switch serviceType {
	case TypeProduct:
		token, err := requestToken(self.clientID, self.clientSecret)
		return convertToString(token), err
	case TypeDemo:
		token, err := requestToken(self.demoClientID, self.demoClientSecret)
		return convertToString(token), err
	default:
		token, err := requestToken(SandboxClientID, SandboxClientSecret)
		return convertToString(token), err
	}
}


// 클라이언트 시크릿 반환
func (self *Codef) getClientSecret(serviceType ServiceType) string {
	switch serviceType {
	case TypeProduct:
		return self.clientSecret
	case TypeDemo:
		return self.demoClientSecret
	default:
		return SandboxClientSecret
	}
}

// 정식서버 사용을 위한 클라이언트 정보 설정
func (self *Codef) SetClientInfo(clientID, clientSecret string) {
	self.clientID = clientID
	self.clientSecret = clientSecret
}

// 데모 서버 사용을 위한 클라이언트 정보 설정
func (self *Codef) SetClientInfoForDemo(clientId, clientSecret string) {
	self.demoClientID = clientId
	self.demoClientSecret = clientSecret
}

// 액세스 토큰 정보 셋팅
func (self *Codef) SetAccessToken(accessToken string, serviceType ServiceType) {
	switch serviceType {
	case TypeProduct:
		self.accessToken.product = accessToken
	case TypeDemo:
		self.accessToken.demo = accessToken
	default:
		self.accessToken.sandbox = accessToken
	}
}

// 액세스 토큰 정보 셋팅
func (self *Codef) getAccessToken(serviceType ServiceType) *string {
	switch serviceType {
	case TypeProduct:
		return &self.accessToken.product
	case TypeDemo:
		return &self.accessToken.demo
	default:
		return &self.accessToken.sandbox
	}
}

// 서비스 상태에 해당하는 요청 정보를 가져온다
// return (domain, clientID, clientSecret)
func (self *Codef) getReqInfoByServiceType(serviceType ServiceType) *requestInfo {
	switch serviceType {
	case TypeProduct:
		return &requestInfo{APIDomain, self.clientID, self.clientSecret}
	case TypeDemo:
		return &requestInfo{DemoDomain, self.demoClientID, self.demoClientSecret}
	default:
		return &requestInfo{SandboxDomain, SandboxClientID, SandboxClientSecret}
	}
}

// 2Way 키워드가 없는지 확인
func isEmptyTwoWayKeyword(param map[string]interface{}) bool {
	if _, ok := param["is2Way"]; ok {
		return false
	}
	if _, ok := param["twoWayInfo"]; ok {
		return false
	}
	return true
}

// 2way 상품 요청 시 필수 데이터 존재하는지 확인
func hasTwoWayInfo(param map[string]interface{}) bool {
	is2Way, ok := param["is2Way"]
	if !ok || reflect.TypeOf(is2Way).Kind() != reflect.Bool || !is2Way.(bool) {
		return false
	}

	twoWayInfo, ok := param["twoWayInfo"].(map[string]interface{})
	if !ok || twoWayInfo == nil {
		return false
	}

	return checkNeedValueInTwoWayInfo(twoWayInfo)
}

// twoWayInfo 정보 내부에 필요한 데이터가 존재하는지 체크
func checkNeedValueInTwoWayInfo(twoWayInfo map[string]interface{}) bool {
	if _, ok := twoWayInfo["jobIndex"]; !ok {
		return false
	}
	if _, ok := twoWayInfo["threadIndex"]; !ok {
		return false
	}
	if _, ok := twoWayInfo["jti"]; !ok {
		return false
	}
	if _, ok := twoWayInfo["twoWayTimestamp"]; !ok {
		return false
	}
	return true
}

func convertToString(tokenMap map[string]interface{}) (string){
	return tokenMap["access_token"].(string)
}