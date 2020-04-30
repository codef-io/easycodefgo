package easycodefgo

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dc7303/easycodefgo/message"
)

const repeatCount = 3

func execute(urlPath string, body map[string]interface{}) *Response {
	_, clientID, clientSecret := getReqInfoByServiceType(ServiceType)

	// TODO: 에러처리 추가 확인 필요
	err := setToken(clientID, clientSecret, &AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(body)
	if err != nil {
		return NewResponse(message.InvalidJson)
	}
	_ = url.QueryEscape(string(b))

	return nil
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

// CODEF POST 요청
func requestProduct(urlPath, token, bodyStr string) (*Response, error) {
	var body *bytes.Buffer = nil
	if bodyStr != "" {
		body = bytes.NewBufferString(bodyStr)
	}

	body = bytes.NewBufferString(bodyStr)
	req, err := http.NewRequest("POST", urlPath, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	client := http.Client{}
	res, err := client.Do(req)
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

		m := make(map[string]interface{})
		err = json.Unmarshal(resBody, &m)
		if err != nil {
			return nil, err
		}

		return NewResponseByMap(m), nil
	case http.StatusBadRequest:
		return NewResponse(message.BadRequest), nil
	case http.StatusUnauthorized:
		return NewResponse(message.Unauthorized), nil
	case http.StatusForbidden:
		return NewResponse(message.Forbidden), nil
	case http.StatusNotFound:
		return NewResponse(message.NotFound), nil
	default:
		return NewResponse(message.ServerError), nil
	}
}
