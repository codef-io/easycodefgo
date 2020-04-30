package easycodefgo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const repeatCount = 3

func execute(url string, body map[string]interface{}) {
	//domain, clientID, clientSecret := getReqInfoByServiceType(ServiceType)

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

//func setToken(clientID, clientSecret string) {
//	i := 0
//	if AccessToken == "" {
//		for i < repeatCount {
//
//		}
//	}
//}

// 액세스 토큰 요청
func requestToken(clientID, clientSecret string) (map[string]interface{}, error) {
	body := bytes.NewBufferString("grant_type=client_credentials&scope=read")
	url := OAuthDomain + PathGetToken
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	auth := clientID + ":" + clientSecret
	authEnc := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+authEnc)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Response status code: %d", res.StatusCode))
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(resBody, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
