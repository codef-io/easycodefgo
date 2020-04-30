package easycodefgo

const repeatCount = 3

func execute(url string, body map[string]interface{}) {
	//getReqInfoByServiceType(ServiceType)

}

// 서비스 상태에 해당하는 요청 정보를 가져온다
// return (domain, clientID, clientSecret)
func getReqInfoByServiceType(serviceType ServiceStatus) (string, string, string) {
	switch serviceType {
	case StatusProduct:
		return APIDomain, ClientID, ClientSecret
	case StatusDemo:
		return DemoDomain, DemoClientID, DemoClientSecret
	default:
		return SandboxDomain, SandboxClientID, SandboxClientSecret
	}
}
