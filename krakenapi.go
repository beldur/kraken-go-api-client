package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	API_URL        = "https://api.kraken.com"
	API_VERISON    = "0"
	API_USER_AGENT = "Kraken GO Api Agent"
)

var publicMethods = []string{
	"Time",
	"Assets",
	"AssetPairs",
	"Ticker",
	"Depth",
	"Trades",
	"Spread",
}

var privateMethods = []string{
	"Balance",
}

type KrakenResponse struct {
	Error  []string    `json:error`
	Result interface{} `json:result`
}

type KrakenApi struct {
	key    string
	secret string
	client *http.Client
}

func NewKrakenApi(key, secret string) *KrakenApi {
	client := &http.Client{}

	return &KrakenApi{key, secret, client}
}

func (api *KrakenApi) Query(method string, data map[string]string) (interface{}, error) {
	values := url.Values{}
	for key, value := range data {
		values.Set(key, value)
	}

	if isStringInSlice(method, publicMethods) {
		return api.queryPublic(method, values)
	} else if isStringInSlice(method, privateMethods) {
		return api.queryPrivate(method, values)
	}

	return nil, fmt.Errorf("Method '%s' is not valid!", method)
}

func (api *KrakenApi) queryPublic(method string, values url.Values) (interface{}, error) {
	url := fmt.Sprintf("%s/%s/public/%s", API_URL, API_VERISON, method)
	resp, err := api.doRequest(url, values, nil)

	return resp, err
}

func (api *KrakenApi) queryPrivate(method string, values url.Values) (interface{}, error) {
	reqUrl := fmt.Sprintf("%s/%s/private/%s", API_URL, API_VERISON, method)
	nonce := time.Now().UnixNano()
	nonce = 1387755202869365200

	fmt.Println(nonce, reqUrl, api.secret)

	base64Secret, _ := base64.StdEncoding.DecodeString(api.secret)
	values.Set("nonce", fmt.Sprintf("%d", nonce))

	sha := sha256.New()
	sha.Write([]byte(values.Get("nonce") + values.Encode()))

	mac := hmac.New(sha512.New, base64Secret)
	mac.Write([]byte(reqUrl + string(sha.Sum(nil))))

	fmt.Printf("%s\n, % X\n, % X\n, % X\n", values.Encode(), sha.Sum(nil), base64Secret, mac.Sum(nil))

	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	fmt.Printf("%+v", sign)

	headers := map[string]string{
		"API-Key":  api.key,
		"API-Sign": sign,
	}

	resp, err := api.doRequest(reqUrl, values, headers)

	return resp, err
}

func (api *KrakenApi) doRequest(reqUrl string, values url.Values, headers map[string]string) (interface{}, error) {

	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}

	req.Header.Add("User-Agent", API_USER_AGENT)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}

	var jsonData KrakenResponse
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}

	if len(jsonData.Error) > 0 {
		return nil, fmt.Errorf("Could not execute request! (%s)", jsonData.Error)
	}

	return jsonData.Result, nil
}

func isStringInSlice(term string, list []string) bool {
	for _, found := range list {
		if term == found {
			return true
		}
	}
	return false
}
