package easycodefgo

type ServiceStatus int

const (
	OAuthDomain         = "https://oauth.codef.io"               // OAUTH 서버 도메인
	PathGetToken        = "/oauth/token"                         // OAUTH 엑세스 토큰 발급 URL PATH
	SandboxDomain       = "https://sandbox.codef.io"             // 샌드박스 서버 도메인
	SandboxClientID     = "ef27cfaa-10c1-4470-adac-60ba476273f9" // 샌드박스 엑세스 토큰 발급을 위한 클라이언트 아이디
	SandboxClientSecret = "83160c33-9045-4915-86d8-809473cdf5c3" // 샌드박스 액세스 토큰 발급을 위한 클라이언트 시크릿
	DemoDomain          = "https://development.codef.io"         // 데모 서버 도메인
	APIDomain           = "https://api.codef.io"                 // 정식 서버 도메인

	PathCreateAccount  = "/v1/account/create"           // 계정 등록 URL
	PathAddAccount     = "/v1/account/add"              // 계정 추가 URL
	PathUpdateAccount  = "/v1/account/update"           // 계정 수정 URL
	PathDeleteAccount  = "/v1/account/delete"           // 계정 삭제 URL
	PathGetAccountList = "/v1/account/list"             // 계정 목록 조회 URL
	PathGetCIDList     = "/v1/account/connectedId-list" // 커넥티드 아이디 목록 조회 URL

	Result       = "result"       // 응답부 수행 결과 키워드
	Code         = "code"         // 응답부 수행 결과 메시지 코드 키워드
	Message      = "message"      // 응답부 수행 결과 메시지 키워드
	ExtraMessage = "extraMessage" // 응답부 수행 결과 추가 메시지 키워드
	Data         = "data"
	AccountList  = "accountList" // 계정 목록 키워드
	ConnectedID  = "connectedId"

	InvalidToken = "invalidToken" // 엑세스 토큰 거절 사유1
	AccessDenied = "accessDenied" // 엑세스 토큰 거절 사유2
)

const (
	StatusProduct ServiceStatus = iota // 정식버전
	StatusDemo                         // 데모 버전
	StatusSandbox                      // 샌드박스
)

var (
	AccessToken      string          // OAUTH2.0 토큰
	ProductURL       string          // 요청 상품 주소
	DemoClientID     string          // 데모 엑세스 토큰 밝브을 위한 클라이언트 아이디
	DemoClientSecret string          // 데모 엑세스 토큰 밝브을 위한 클라이언트 시크릿
	ClientID         string          // 정식 엑세스 토큰 발급을 위한 클라이언트 아이디
	ClientSecret     string          // 정식 엑세스 토큰 발급을 위한 클라이언트 시크릿
	PublicKey        string          // RSA암호화를 위한 퍼블릭키
	ServiceType      = StatusSandbox // 코드에프 서비스 타입(0: 정식, 1:데모 2: 샌드박스)
)

// 클라이언트 아이디
func getCodefDomain(serviceType ServiceStatus) string {
	switch serviceType {
	case StatusProduct:
		return APIDomain
	case StatusDemo:
		return DemoDomain
	default:
		return SandboxDomain
	}
}

// 클라이언트 시크릿 반환
func getClientSecret(serviceType ServiceStatus) string {
	switch serviceType {
	case StatusProduct:
		return ClientSecret
	case StatusDemo:
		return DemoClientSecret
	default:
		return SandboxClientSecret
	}
}

// 정식서버 사용을 위한 클라이언트 정보 설정
func SetClientInfo(clientID, clientSecret string) {
	ClientID = clientID
	ClientSecret = clientSecret
}

// 데모 서버 사용을 위한 클라이언트 정보 설정
func SetClientInfoForDemo(clientId, clientSecret string) {
	DemoClientID = clientId
	DemoClientSecret = clientSecret
}

// RSA 암호화를 위한 퍼블릭키 생성
func SetPublicKey(publicKey string) {
	PublicKey = publicKey
}
