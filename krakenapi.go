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
	API_USER_AGENT = "Kraken GO API Agent"
)

// List of valid public methods
var publicMethods = []string{
	"Time",
	"Assets",
	"AssetPairs",
	"Ticker",
	"Depth",
	"Trades",
	"Spread",
}

// List of valid private methods
var privateMethods = []string{
	"Balance",
	"TradeBalance",
	"OpenOrders",
	"ClosedOrders",
	"QueryOrders",
	"TradesHistory",
	"QueryTrades",
	"OpenPositions",
	"Ledgers",
	"QueryLedgers",
	"TradeVolume",
	"AddOrder",
	"CancelOrder",
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

// Create a new Kraken Api struct
func NewKrakenApi(key, secret string) *KrakenApi {
	client := &http.Client{}

	return &KrakenApi{key, secret, client}
}

// Send a query to Kraken api for given method and parameters
func (api *KrakenApi) Query(method string, data map[string]string) (interface{}, error) {
	values := url.Values{}
	for key, value := range data {
		values.Set(key, value)
	}

	// Check if method is public or private
	if isStringInSlice(method, publicMethods) {
		return api.queryPublic(method, values)
	} else if isStringInSlice(method, privateMethods) {
		return api.queryPrivate(method, values)
	}

	return nil, fmt.Errorf("Method '%s' is not valid!", method)
}

// Execute a public method query
func (api *KrakenApi) queryPublic(method string, values url.Values) (interface{}, error) {
	url := fmt.Sprintf("%s/%s/public/%s", API_URL, API_VERISON, method)
	resp, err := api.doRequest(url, values, nil)

	return resp, err
}

// Execute a private method query
func (api *KrakenApi) queryPrivate(method string, values url.Values) (interface{}, error) {
	urlPath := fmt.Sprintf("/%s/private/%s", API_VERISON, method)
	reqUrl := fmt.Sprintf("%s%s", API_URL, urlPath)
	secret, _ := base64.StdEncoding.DecodeString(api.secret)
	values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

	// Create signature
	signature := createSignature(urlPath, values, secret)

	// Add Key and signature to request headers
	headers := map[string]string{
		"API-Key":  api.key,
		"API-Sign": signature,
	}

	resp, err := api.doRequest(reqUrl, values, headers)

	return resp, err
}

// Executes a HTTP Request to the Kraken API and returns the result
func (api *KrakenApi) doRequest(reqUrl string, values url.Values, headers map[string]string) (interface{}, error) {

	// Create request
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}

	req.Header.Add("User-Agent", API_USER_AGENT)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Execute request
	resp, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}
	defer resp.Body.Close()

	// Read request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}

	// Parse request
	var jsonData KrakenResponse
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! (%s)", err.Error())
	}

	// Check for Kraken API error
	if len(jsonData.Error) > 0 {
		return nil, fmt.Errorf("Could not execute request! (%s)", jsonData.Error)
	}

	return jsonData.Result, nil
}

// Helper function to test if given term is in a list of strings
func isStringInSlice(term string, list []string) bool {
	for _, found := range list {
		if term == found {
			return true
		}
	}
	return false
}

// Creates a sha256 hash
func getSha256(input []byte) []byte {
	sha := sha256.New()
	sha.Write(input)
	return sha.Sum(nil)
}

// Create a hmac hash with sha512
func getHMacSha512(message, secret []byte) []byte {
	mac := hmac.New(sha512.New, secret)
	mac.Write(message)
	return mac.Sum(nil)
}

func createSignature(urlPath string, values url.Values, secret []byte) string {
	// See https://www.kraken.com/help/api#general-usage for more information
	shaSum := getSha256([]byte(values.Get("nonce") + values.Encode()))
	macSum := getHMacSha512(append([]byte(urlPath), shaSum...), secret)
	return base64.StdEncoding.EncodeToString(macSum)
}
