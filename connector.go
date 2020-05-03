package easycodefgo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dc7303/easycodefgo/message"
)

// CODEF API 요청 실행
func execute(
	urlPath string,
	body map[string]interface{},
	serviceType ServiceStatus,
) (*Response, error) {
	domain, clientID, clientSecret := getReqInfoByServiceType(serviceType)

	err := setToken(clientID, clientSecret, &AccessToken)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	encBodyStr := url.QueryEscape(string(b))

	res, err := requestProduct(domain+urlPath, AccessToken, encBodyStr)
	if err != nil {
		return nil, err
	}

	return res, nil
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

// 액세스 토큰 셋팅
func setToken(clientID, clientSecret string, accessToken *string) error {
	repeatCount := 3
	i := 0
	if *accessToken == "" {
		for i < repeatCount {
			tokenMap, err := requestToken(clientID, clientSecret)
			if err != nil {
				return err
			}
			if token, ok := tokenMap["access_token"]; ok {
				*accessToken = token.(string)
			}

			if *accessToken != "" {
				break
			}

			time.Sleep(time.Millisecond * 20)
			i++
		}
	}

	return nil
}

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

	res, err := http.DefaultClient.Do(req)
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

// CODEF POST 요청
func requestProduct(reqURL, token, bodyStr string) (*Response, error) {
	var body *bytes.Buffer = nil
	if bodyStr != "" {
		body = bytes.NewBufferString(bodyStr)
	}

	body = bytes.NewBufferString(bodyStr)
	req, err := http.NewRequest("POST", reqURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		resultData, err := url.QueryUnescape(string(resBody))
		if err != nil {
			return nil, err
		}
		m := make(map[string]interface{})
		err = json.Unmarshal([]byte(resultData), &m)
		if err != nil {
			return nil, err
		}

		return newResponseByMap(m), nil
	case http.StatusBadRequest:
		return newResponseByMessage(message.BadRequest), nil
	case http.StatusUnauthorized:
		return newResponseByMessage(message.Unauthorized), nil
	case http.StatusForbidden:
		return newResponseByMessage(message.Forbidden), nil
	case http.StatusNotFound:
		return newResponseByMessage(message.NotFound), nil
	default:
		return newResponseByMessage(message.ServerError), nil
	}
}
