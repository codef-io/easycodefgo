package easycodefgo

// CODEF API
type Codef struct {
	AccessToken      string // OAUTH2.0 토큰
	DemoClientID     string // 데모 엑세스 토큰 밝브을 위한 클라이언트 아이디
	DemoClientSecret string // 데모 엑세스 토큰 밝브을 위한 클라이언트 시크릿
	ClientID         string // 정식 엑세스 토큰 발급을 위한 클라이언트 아이디
	ClientSecret     string // 정식 엑세스 토큰 발급을 위한 클라이언트 시크릿
}

// 요청 정보
type requestInfo struct {
	Domain       string
	ClientID     string
	ClientSecret string
}

// 상품 요청
func (self *Codef) RequestProduct(
	productPath string,
	serviceType ServiceType,
	param map[string]interface{},
) (string, error) {
	validFlag := true

	// 클라이언트 정보 체크
	validFlag = self.checkClientInfo(serviceType)
	if !validFlag {
		res := newResponseByMessage(messageEmptyClientInfo)
		return res.WriteValueAsString(), nil
	}

	// 추가인증 키워드 체크
	validFlag = checkTwoWayKeyword(param)
	if !validFlag {
		res := newResponseByMessage(messageInvalid2WayKeyword)
		return res.WriteValueAsString(), nil
	}

	reqInfo := self.getReqInfoByServiceType(serviceType)

	res, err := execute(productPath, param, &self.AccessToken, reqInfo)
	if err != nil {
		return "", err
	}
	return res.WriteValueAsString(), nil
}

// 클라이언트 정보 검사
func (self *Codef) checkClientInfo(serviceType ServiceType) bool {
	switch serviceType {
	case TypeProduct:
		if TrimAll(self.ClientID) == "" || TrimAll(self.ClientSecret) == "" {
			return false
		}
	case TypeDemo:
		if TrimAll(self.DemoClientID) == "" || TrimAll(self.DemoClientSecret) == "" {
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
func (self *Codef) GetConnectedIDList(serviceType ServiceType) (string, error) {
	return self.RequestProduct(PathGetCIDList, serviceType, nil)
}

// 토큰 발급
func (self *Codef) RequestToken(serviceType ServiceType) (map[string]interface{}, error) {
	switch serviceType {
	case TypeProduct:
		return requestToken(self.ClientID, self.ClientSecret)
	case TypeDemo:
		return requestToken(self.DemoClientID, self.DemoClientSecret)
	default:
		return requestToken(SandboxClientID, SandboxClientSecret)
	}
}

// 클라이언트 시크릿 반환
func (self *Codef) getClientSecret(serviceType ServiceType) string {
	switch serviceType {
	case TypeProduct:
		return self.ClientSecret
	case TypeDemo:
		return self.DemoClientSecret
	default:
		return SandboxClientSecret
	}
}

// 정식서버 사용을 위한 클라이언트 정보 설정
func (self *Codef) SetClientInfo(clientID, clientSecret string) {
	self.ClientID = clientID
	self.ClientSecret = clientSecret
}

// 데모 서버 사용을 위한 클라이언트 정보 설정
func (self *Codef) SetClientInfoForDemo(clientId, clientSecret string) {
	self.DemoClientID = clientId
	self.DemoClientSecret = clientSecret
}

// 서비스 상태에 해당하는 요청 정보를 가져온다
// return (domain, clientID, clientSecret)
func (self *Codef) getReqInfoByServiceType(serviceType ServiceType) *requestInfo {
	switch serviceType {
	case TypeProduct:
		return &requestInfo{APIDomain, self.ClientID, self.ClientSecret}
	case TypeDemo:
		return &requestInfo{DemoDomain, self.DemoClientID, self.DemoClientSecret}
	default:
		return &requestInfo{SandboxDomain, SandboxClientID, SandboxClientSecret}
	}
}

// 2Way 키워드 존재 여부 확인
func checkTwoWayKeyword(param map[string]interface{}) bool {
	if _, ok := param["is2Way"]; ok {
		return false
	}
	if _, ok := param["twoWayInfo"]; ok {
		return false
	}
	return true
}
