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
)

// CODEF API 요청 실행
func execute(
	urlPath string,
	body map[string]interface{},
	accessToken *string,
	reqInfo *requestInfo,
) (*Response, error) {
	setToken(reqInfo.clientID, reqInfo.clientSecret, accessToken)

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	encBodyStr := url.QueryEscape(string(b))

	res, err := requestProduct(reqInfo.domain+urlPath, *accessToken, encBodyStr)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 액세스 토큰 셋팅
func setToken(clientID, clientSecret string, accessToken *string) {
	repeatCount := 3
	i := 0
	if *accessToken == "" {
		for i < repeatCount {
			tokenMap, err := requestToken(clientID, clientSecret)
			if err != nil {
				i++
				continue
			}
			if token, ok := tokenMap["access_token"]; ok {
				*accessToken = token.(string)
			}

			if *accessToken != "" {
				break
			}
			i++
		}
	}
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
	jsonDec := json.NewDecoder(bytes.NewBuffer([]byte(resBody)))
	jsonDec.UseNumber()
	if err := jsonDec.Decode(&m); err != nil {
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

		jsonDec := json.NewDecoder(bytes.NewBuffer([]byte(resultData)))
		jsonDec.UseNumber()
		if err := jsonDec.Decode(&m); err != nil {
			return nil, err
		}

		return newResponseByMap(m), nil
	case http.StatusBadRequest:
		return newResponseByMessage(messageBadRequest), nil
	case http.StatusUnauthorized:
		return newResponseByMessage(messageUnauthorized), nil
	case http.StatusForbidden:
		return newResponseByMessage(messageForbidden), nil
	case http.StatusNotFound:
		return newResponseByMessage(messageNotFound), nil
	default:
		return newResponseByMessage(messageServerError), nil
	}
}
