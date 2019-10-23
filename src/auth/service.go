package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"ebaocloud.com/go-js-service/common/conf"
)

// GetToken ...
func GetToken() (string, error) {

	url := "http://" + conf.CasServiceName + "/cas/ebao/v1/json/tickets/"

	userInfo := make(map[string]interface{})
	userInfo["username"] = Username
	userInfo["password"] = Password

	bytesData, err := json.Marshal(userInfo)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(bytesData)

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", "application/json")
	// req.Header.Set("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode > 200 {
		err = errors.New("get token error")
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	bodyData := &Token{}
	// bodyData := &map[string]interface{}{}
	err = json.Unmarshal(body, bodyData)
	if err != nil {
		return "", err
	}
	// fmt.Println(bodyData)
	return bodyData.AccessToken, nil
}

// GetTenant ...
func GetTenant(token string) (string, error) {

	url := "http://" + conf.CasServiceName + "/cas/oauth2.0/profile?access_token=" + token

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	fmt.Println(res)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode > 200 {
		err = errors.New("get tenant error")
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	bodyData := &Profile{}
	err = json.Unmarshal(body, bodyData)
	if err != nil {
		return "", err
	}
	return bodyData.Attributes[0].Value, nil
}
