package auth

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"platform-api-demo/util"
	"strings"
)

func GetOpenId() {

	// set the domain url for calling OpenPlatform
	reqUrl := "https://service.xxx.com/platformapi/chainup/open/auth/token"
	method := "POST"

	// set the appKey/appSecret
	appKey := "test"
	appSecret := "abc"

	// set the parameters of the getOpenId interface
	code := "123"

	paramsMap := make(map[string]string)
	paramsMap["appKey"] = appKey
	paramsMap["code"] = code

	// calculate the signature
	sign := util.GetSign(paramsMap, appSecret)

	paramsMap["sign"] = sign

	paramsMapBytes, _ := json.Marshal(paramsMap)
	payload := strings.NewReader(string(paramsMapBytes))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client {Transport:tr}
	req, err := http.NewRequest(method, reqUrl, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("GetOpenId result: %s", string(body)))
}
